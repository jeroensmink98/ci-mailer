package main

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Setup a temporary file
	content := "Hello, world!"
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Test readFile
	got := readFile(tmpfile.Name())
	if got != content {
		t.Errorf("readFile(%q) = %q, want %q", tmpfile.Name(), got, content)
	}
}
