package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

// GetProjectRoot returns the absolute path to the project root directory
func GetProjectRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	
	// Walk up until we find go.mod
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			t.Fatal("could not find project root")
		}
		wd = parent
	}
}

// TestdataPath returns the path to a file in the testdata directory
func TestdataPath(t *testing.T, segments ...string) string {
	t.Helper()
	root := GetProjectRoot(t)
	parts := append([]string{root, "testdata"}, segments...)
	return filepath.Join(parts...)
}

// TempDir creates a temporary directory for tests
func TempDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	return dir
}