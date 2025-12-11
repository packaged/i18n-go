package i18n

import (
	"testing"
)

func TestBCP47_LanguageOnly(t *testing.T) {
	tag, err := BCP47("", "fr")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := tag.String(); got != "fr" {
		t.Fatalf("expected fr, got %s", got)
	}
}

func TestBCP47_CountryOnly_UsesDefault(t *testing.T) {
	tag, err := BCP47("US", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := tag.String(); got != "en-US" {
		t.Fatalf("expected en-US, got %s", got)
	}
}

func TestBCP47_BothLanguageAndCountry(t *testing.T) {
	tag, err := BCP47("CA", "fr")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := tag.String(); got != "fr-CA" {
		t.Fatalf("expected fr-CA, got %s", got)
	}
}

func TestBCP47_UnknownCountry_WithLanguage_FallsBackToLanguageOnly(t *testing.T) {
	tag, err := BCP47("ZZ", "es")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := tag.String(); got != "es" {
		t.Fatalf("expected es, got %s", got)
	}
}

func TestBCP47_UnknownCountry_NoLanguage_DefaultsToEnglish(t *testing.T) {
	tag, err := BCP47("ZZ", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := tag.String(); got != "en" {
		t.Fatalf("expected en, got %s", got)
	}
}

func TestBCP47_InvalidLanguage_ReturnsError(t *testing.T) {
	if _, err := BCP47("US", "xx"); err == nil {
		t.Fatalf("expected error for invalid language code, got nil")
	}
}
