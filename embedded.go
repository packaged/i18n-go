package translate

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// Embedded is a Translator that loads its translations from a file within an fs.FS
// (for example, an embed.FS). The data is compiled into the binary at build time
// when using //go:embed in the caller and passed via the fs parameter.
//
// Supported formats: JSON (.json) and YAML (.yaml, .yml).
type Embedded struct {
	m *Map
}

// NewEmbeddedFile loads translations from the given fs and path, returning an Embedded translator.
// The file format is detected by extension: .json, .yaml, .yml
func NewEmbeddedFile(fsys fs.FS, path string) (*Embedded, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("read embedded file: %w", err)
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
	return &Embedded{m: m}, nil
}

func (e *Embedded) Translate(key string) string {
	return e.m.Translate(key)
}

func (e *Embedded) TranslateWith(key string, args map[string]interface{}) string {
	return e.m.TranslateWith(key, args)
}
