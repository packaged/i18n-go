package translate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// File is a Translator that loads its translations from a file.
// Supported formats: JSON (.json) and YAML (.yaml, .yml).
type File struct {
	m *Map
}

// NewFileFrom loads translations from the given source file and returns a File translator.
// The file format is detected by extension: .json, .yaml, .yml
func NewFileFrom(path string) (*File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(path))
	translations := make(map[string]string)

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &translations); err != nil {
			return nil, fmt.Errorf("parse json: %w", err)
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &translations); err != nil {
			return nil, fmt.Errorf("parse yaml: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	m := NewMap()
	for k, v := range translations {
		m.Add(k, v)
	}
	return &File{m: m}, nil
}

func (f *File) Translate(key string) string {
	return f.m.Translate(key)
}

func (f *File) TranslateWith(key string, args map[string]interface{}) string {
	return f.m.TranslateWith(key, args)
}
