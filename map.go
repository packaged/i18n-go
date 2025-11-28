package translate

import (
	"fmt"
	"strings"
)

type Map struct {
	translations map[string]string
}

func NewMap() *Map {
	return &Map{translations: make(map[string]string)}
}

func (m *Map) Add(key, value string) {
	m.translations[key] = value
}

func (m *Map) Translate(key string) string {
	return m.translations[key]
}

func (m *Map) TranslateWith(key string, args map[string]interface{}) string {
	original := m.translations[key]
	for k, v := range args {
		original = strings.ReplaceAll(original, "{"+k+"}", fmt.Sprintf("%v", v))
	}
	return original
}
