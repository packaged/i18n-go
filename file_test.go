package translate

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, pattern string, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, pattern)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("failed writing temp file: %v", err)
	}
	return path
}

func TestFileJSON(t *testing.T) {
	jsonContent := `{
        "hello_world_b10a": "Hello World",
        "hello_name_a696": "Hello {name}"
    }`

	path := writeTempFile(t, "en.json", jsonContent)
	tr, err := NewFileFrom(path)
	if err != nil {
		t.Fatalf("failed to load json file: %v", err)
	}

	if val := tr.Translate("hello_world_b10a"); val != "Hello World" {
		t.Error("Basic JSON translation failed, got " + val)
	}

	if val := tr.TranslateWith("hello_name_a696", Args("name", "John")); val != "Hello John" {
		t.Error("Arg JSON translation failed, got " + val)
	}
}

func TestFileYAML(t *testing.T) {
	// Simple flat YAML map
	yamlContent := "" +
		"hello_world_b10a: Hello World\n" +
		"hello_name_a696: 'Hello {name}'\n"

	path := writeTempFile(t, "en.yaml", yamlContent)
	tr, err := NewFileFrom(path)
	if err != nil {
		t.Fatalf("failed to load yaml file: %v", err)
	}

	if val := tr.Translate("hello_world_b10a"); val != "Hello World" {
		t.Error("Basic YAML translation failed, got " + val)
	}

	if val := tr.TranslateWith("hello_name_a696", Args("name", "Jane")); val != "Hello Jane" {
		t.Error("Arg YAML translation failed, got " + val)
	}
}
