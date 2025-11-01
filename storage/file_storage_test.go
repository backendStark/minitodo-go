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

func TestFileStorage_Load_FileNotExist(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "not_existing_file.json")
	fs := NewFileStorage(filename)

	if fs == nil {
		t.Fatal("Expected FileStorage to be non-nil, got nil")
	}

	tasks, err := fs.Load()
	if err != nil {
		t.Error("Error loading tasks")
	}

	if len(tasks) != 0 {
		t.Error("Expected not exisiting file creating with empty tasks list")
	}

	if _, err := os.Stat(filename); err != nil {
		t.Error("Expected Stat method return FileInfo")
	}

	data, err := os.ReadFile(filename)

	if err != nil {
		t.Error("Error reading file:", filename)
	}

	if string(data) != "[]" {
		t.Error("Expected data equal '[]', but got:", string(data))
	}
}
