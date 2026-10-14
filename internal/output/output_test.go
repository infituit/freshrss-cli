package output

import (
	"strings"
	"testing"
)

func TestWriteTableRendersMapAsKeyValue(t *testing.T) {
	var out strings.Builder
	if err := Write(&out, Table, map[string]any{"name": "home", "current": true}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "KEY") || !strings.Contains(out.String(), "name") || !strings.Contains(out.String(), "home") {
		t.Fatalf("table output = %q", out.String())
	}
}
