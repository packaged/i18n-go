package i18n

// entry represents a translation record that can contain singular and plural forms.
type entry struct {
	Singular string
	Plural   string
}

type Map struct {
	translations map[string]entry
}

func NewMap() *Map {
	return &Map{translations: make(map[string]entry)}
}

func NewBasicMap(from map[string]string) *Map {
	m := NewMap()
	for k, v := range from {
		m.Add(k, v)
	}
	return m
}

// Add registers a singular-only translation value for backward compatibility.
func (m *Map) Add(key, value string) {
	m.translations[key] = entry{Singular: value}
}

// AddPlural registers both singular and plural versions for a translation key.
func (m *Map) AddPlural(key, singular, plural string) {
	m.translations[key] = entry{Singular: singular, Plural: plural}
}

func (m *Map) Translate(key string) string {
	return m.translations[key].Singular
}

func (m *Map) TranslateWith(key string, args map[string]interface{}) string {
	return Replacements(m.translations[key].Singular, args)
}

// TranslatePlural returns the singular form when number==1, otherwise the plural form.
// If a plural form is not set, it falls back to ENPlural transformation of the singular.
func (m *Map) TranslatePlural(key string, number int64, args map[string]interface{}) string {
	e := m.translations[key]
	original := ENPlural(e.Singular, number)
	if number != 1 && e.Plural != "" {
		original = ENPlural(e.Plural, number)
	}
	if args == nil {
		args = map[string]interface{}{}
	}
	if _, ok := args["qty"]; !ok {
		args["qty"] = number
	}
	return Replacements(original, args)
}
