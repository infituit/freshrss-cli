package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/infituit/freshrss-cli/internal/config"
	"github.com/infituit/freshrss-cli/pkg/fever"
	"github.com/infituit/freshrss-cli/pkg/version"
)

func Run(args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		printUsage(stdout)
		return nil
	}

	switch args[0] {
	case "help", "-h", "--help":
		printUsage(stdout)
		return nil
	case "version", "--version":
		fmt.Fprint(stdout, version.FormatVersion())
		return nil
	case "auth-check":
		return runAuthCheck(args[1:], stdout, stderr)
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runAuthCheck(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("auth-check", flag.ContinueOnError)
	flags.SetOutput(stderr)

	var cfg config.Config
	flags.StringVar(&cfg.BaseURL, "url", "", "FreshRSS base URL or Fever endpoint")
	flags.StringVar(&cfg.Username, "user", "", "FreshRSS username")
	flags.StringVar(&cfg.APIPassword, "api-password", "", "FreshRSS API password")
	flags.StringVar(&cfg.APIKey, "api-key", "", "precomputed Fever API key")

	if err := flags.Parse(args); err != nil {
		return err
	}

	cfg = config.MergeEnv(cfg)
	client, err := fever.NewClient(cfg, http.DefaultClient)
	if err != nil {
		return err
	}

	resp, err := client.AuthCheck(context.Background())
	if err != nil {
		return err
	}
	if resp.Auth != 1 {
		return fmt.Errorf("authentication failed")
	}

	fmt.Fprintf(stdout, "authenticated: api_version=%d last_refreshed_on_time=%s\n", resp.APIVersion, resp.LastRefreshedOnTime)
	return nil
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "freshrss-cli - FreshRSS Fever API command line client")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  freshrss-cli help")
	fmt.Fprintln(w, "  freshrss-cli version")
	fmt.Fprintln(w, "  freshrss-cli auth-check [-url URL] [-user USER -api-password PASSWORD | -api-key KEY]")
}
