// Package version provides build version metadata.
package version

import "fmt"

var (
	// VERSION is the application version.
	VERSION = "debug"
	// BUILDTIME is the build timestamp.
	BUILDTIME = "unknown"
	// GITHASH is the short Git commit used for the build.
	GITHASH = "unknown"
)

// FormatVersion returns the application version metadata in a human-readable form.
func FormatVersion() string {
	return fmt.Sprintf("Version: %s\nBuildTime: %s\nGitHash: %s\n", VERSION, BUILDTIME, GITHASH)
}
