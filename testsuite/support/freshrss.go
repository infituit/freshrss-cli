//go:build integration

package support

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	defaultFreshRSSPort = 18080
	apiPassword         = "freshrss-api-password"
	webPassword         = "freshrss-web-password"
)

// FreshRSS describes the disposable FreshRSS instance used by an integration test.
type FreshRSS struct {
	BaseURL     string
	Username    string
	APIPassword string
}

// RequireFreshRSS starts a disposable FreshRSS Compose project and removes it at test end.
func RequireFreshRSS(t testing.TB) FreshRSS {
	t.Helper()
	root := repositoryRoot(t)
	port := configuredPort(t)
	project := fmt.Sprintf("rss-cli-test-%d", time.Now().UnixNano())
	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)
	envFile := filepath.Join(t.TempDir(), "compose.env")
	writeComposeEnv(t, envFile, port, baseURL)

	compose := func(args ...string) {
		t.Helper()
		if err, output := runCompose(root, project, envFile, args...); err != nil {
			t.Fatalf("docker compose failed: %v\n%s", err, output)
		}
	}
	cleanup := func() {
		t.Helper()
		if err, output := runCompose(root, project, envFile, "down", "--volumes", "--remove-orphans"); err != nil {
			t.Logf("FreshRSS cleanup failed: %v\n%s", err, output)
		}
	}

	t.Cleanup(cleanup)
	compose("up", "-d", "--wait")
	waitForFreshRSS(t, baseURL)
	return FreshRSS{BaseURL: baseURL, Username: "freshrss", APIPassword: apiPassword}
}

func runCompose(root, project, envFile string, args ...string) (error, []byte) {
	allArgs := []string{"compose", "--project-name", project, "--project-directory", root, "--env-file", envFile, "-f", filepath.Join(root, "testsuite/docker-compose.yml")}
	allArgs = append(allArgs, args...)
	cmd := exec.Command("docker", allArgs...)
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	return err, output
}

func repositoryRoot(t testing.TB) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("go.mod not found")
		}
		dir = parent
	}
}

func configuredPort(t testing.TB) int {
	t.Helper()
	value := strings.TrimSpace(os.Getenv("FRESHRSS_TEST_PORT"))
	if value == "" {
		return defaultFreshRSSPort
	}
	port, err := strconv.Atoi(value)
	if err != nil || port < 1 || port > 65535 {
		t.Fatalf("FRESHRSS_TEST_PORT must be a TCP port, got %q", value)
	}
	return port
}

func writeComposeEnv(t testing.TB, path string, port int, baseURL string) {
	t.Helper()
	content := fmt.Sprintf("FRESHRSS_TEST_PORT=%d\nFRESHRSS_TEST_BASE_URL=%s\nFRESHRSS_TEST_API_PASSWORD=%s\nFRESHRSS_TEST_WEB_PASSWORD=%s\n", port, baseURL, apiPassword, webPassword)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write Compose environment: %v", err)
	}
}

func waitForFreshRSS(t testing.TB, baseURL string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	client := &http.Client{Timeout: 5 * time.Second}
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
		if err == nil {
			resp, requestErr := client.Do(req)
			if requestErr == nil {
				resp.Body.Close()
				if resp.StatusCode < http.StatusInternalServerError {
					return
				}
			}
		}
		select {
		case <-ctx.Done():
			t.Fatalf("FreshRSS did not become ready: %v", ctx.Err())
		case <-time.After(time.Second):
		}
	}
}
