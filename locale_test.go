package i18n

import (
	"testing"
)

func TestBCP47_LanguageOnly(t *testing.T) {
	tag := BCP47("", "fr")
	if got := tag.String(); got != "fr-FR" {
		t.Fatalf("expected fr-FR, got %s", got)
	}
}

func TestBCP47_CountryOnly_UsesDefault(t *testing.T) {
	tag := BCP47("US", "")
	if got := tag.String(); got != "en-US" {
		t.Fatalf("expected en-US, got %s", got)
	}
}

func TestBCP47_BothLanguageAndCountry(t *testing.T) {
	tag := BCP47("CA", "fr")
	if got := tag.String(); got != "fr-CA" {
		t.Fatalf("expected fr-CA, got %s", got)
	}
}

func TestBCP47_UnknownCountry_WithLanguage_FallsBackToLanguageOnly(t *testing.T) {
	tag := BCP47("ZZ", "es")
	if got := tag.String(); got != "es-ES" {
		t.Fatalf("expected es-ES, got %s", got)
	}
}

func TestBCP47_UnknownCountry_NoLanguage_DefaultsToEnglish(t *testing.T) {
	tag := BCP47("ZZ", "")
	if got := tag.String(); got != "en-US" {
		t.Fatalf("expected en-US, got %s", got)
	}
}

func TestBCP47_InvalidLanguage_ReturnsError(t *testing.T) {
	tag := BCP47("US", "xx")
	if tag.String() != "en-US" {
		t.Fatalf("expected en-US, got %s", tag.String())
	}
}
