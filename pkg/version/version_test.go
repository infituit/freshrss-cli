package version

import "testing"

func TestFormatVersion(t *testing.T) {
	originalVersion, originalBuildTime, originalGitHash := VERSION, BUILDTIME, GITHASH
	t.Cleanup(func() {
		VERSION, BUILDTIME, GITHASH = originalVersion, originalBuildTime, originalGitHash
	})

	VERSION = "v0.1.0-alpha.1"
	BUILDTIME = "2026-10-14T12:00:00+0000"
	GITHASH = "abc1234"

	const want = "Version: v0.1.0-alpha.1\nBuildTime: 2026-10-14T12:00:00+0000\nGitHash: abc1234\n"
	if got := FormatVersion(); got != want {
		t.Fatalf("FormatVersion() = %q, want %q", got, want)
	}
}
