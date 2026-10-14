package main

import (
	"os"

	"github.com/infituit/rss-cli/internal/cli"
	"github.com/infituit/rss-cli/internal/output"
)

func main() {
	if err := cli.Run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		_ = output.WriteError(os.Stderr, err)
		os.Exit(1)
	}
}
