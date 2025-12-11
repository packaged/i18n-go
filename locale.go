package i18n

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

// BCP47 builds and parses a canonical BCP-47 tag using the provided optional country (region) and language.
// - country should be ISO 3166-1 alpha-2 (e.g., "US").
// - language should be ISO 639-1 (e.g., "en").
// If country is given but language is empty, the country's default language from CountryLanguages is used.
// If both are empty, English ("en") is used.
func BCP47(country, lang string) (language.Tag, error) {
	c := strings.ToUpper(strings.TrimSpace(country))
	l := strings.ToLower(strings.TrimSpace(lang))

	// If only country is provided, use default language for that country when available
	if l == "" && c != "" {
		if def, ok := CountryLanguages[c]; ok {
			l = def
		}
	}
	// Default language when still empty
	if l == "" {
		l = "en"
	}

	// Validate base language
	if _, err := language.ParseBase(l); err != nil {
		return language.Und, err
	}

	// Validate region; if invalid or not a known country in our defaults, drop it
	if c != "" {
		if _, err := language.ParseRegion(c); err != nil {
			c = ""
		}
		if _, ok := CountryLanguages[c]; !ok {
			// Only keep regions we know as real countries for our use-case
			c = ""
		}
	}

	var tagStr string
	if c != "" {
		tagStr = fmt.Sprintf("%s-%s", l, c)
	} else {
		tagStr = l
	}
	tag, err := language.Parse(tagStr)
	if err != nil {
		// This should rarely happen after validations; fallback to language-only
		tag2, err2 := language.Parse(l)
		if err2 != nil {
			return language.Und, err
		}
		return tag2, nil
	}
	return tag, nil
}
