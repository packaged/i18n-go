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
	// Use a generic map to support either string values or plural objects
	translations := make(map[string]interface{})

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
		switch tv := v.(type) {
		case string:
			m.Add(k, tv)
		case map[string]interface{}:
			singular := ""
			plural := ""
			if s, ok := tv["singular"].(string); ok {
				singular = s
			}
			if s, ok := tv["s"].(string); ok {
				singular = s
			}
			if p, ok := tv["plural"].(string); ok {
				plural = p
			}
			if p, ok := tv["p"].(string); ok {
				plural = p
			}
			if singular != "" || plural != "" {
				if plural == "" {
					// allow singular-only objects
					m.Add(k, singular)
				} else {
					m.AddPlural(k, singular, plural)
				}
			}
		default:
			// Unsupported type, ignore silently
		}
	}
	return &File{m: m}, nil
}

func (f *File) Translate(key string) string {
	return f.m.Translate(key)
}

func (f *File) TranslateWith(key string, args map[string]interface{}) string {
	return f.m.TranslateWith(key, args)
}

func (f *File) TranslatePlural(key string, number int64, args map[string]interface{}) string {
	return f.m.TranslatePlural(key, number, args)
}
