package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func createTempFile(t *testing.T, content string) string {
	t.Helper()

	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.json")
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return filename
}

func TestNewFileStorage_Success(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "test.json")
	fs := NewFileStorage(filename)
	if fs == nil {
		t.Fatal("Expected FileStorage to be non-nil, got nil")
	}

	if fs.filename != filename {
		t.Errorf("Expect FileStorage filename equal %s, but got: %s", filename, fs.filename)
	}
}
