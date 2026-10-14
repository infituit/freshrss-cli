package credential

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "credentials.enc")
	want := Values{Username: "alice", APIPassword: "api-password"}
	if err := Save(path, want, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) == `{"username":"alice","api_password":"api-password"}` {
		t.Fatal("credentials must not be stored as plaintext")
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("credentials = %+v, want %+v", got, want)
	}
}

func TestDeleteIsIdempotent(t *testing.T) {
	if err := Delete(filepath.Join(t.TempDir(), "missing")); err != nil {
		t.Fatal(err)
	}
}
