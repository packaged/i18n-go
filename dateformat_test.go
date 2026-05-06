package i18n

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestFormatDateTime(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	d := time.Date(2026, time.January, 30, 15, 4, 5, 0, loc)

	tests := []struct {
		name  string
		tag   language.Tag
		style int
		want  string
	}{
		{"en-US short date", language.Make("en-US"), DateFormatShort, "1/30/26"},
		{"en-US medium date", language.Make("en-US"), DateFormatMedium, "Jan 30, 2026"},
		{"en-US long date", language.Make("en-US"), DateFormatLong, "January 30, 2026"},
		{"en-US full date", language.Make("en-US"), DateFormatFull, "Friday, January 30, 2026"},
		{"en-US short time", language.Make("en-US"), TimeFormatShort, "3:04 PM"},
		{"en-US medium time", language.Make("en-US"), TimeFormatMedium, "3:04:05 PM"},

		{"en-GB short date", language.Make("en-GB"), DateFormatShort, "30/01/2026"},
		{"en-GB long date", language.Make("en-GB"), DateFormatLong, "30 January 2026"},
		{"en-GB short time", language.Make("en-GB"), TimeFormatShort, "15:04"},

		{"en-CA short date", language.Make("en-CA"), DateFormatShort, "2026-01-30"},

		{"fr-FR short date", language.Make("fr-FR"), DateFormatShort, "30/01/2026"},
		{"fr-FR medium date", language.Make("fr-FR"), DateFormatMedium, "30 janv. 2026"},
		{"fr-FR long date", language.Make("fr-FR"), DateFormatLong, "30 janvier 2026"},
		{"fr-FR full date", language.Make("fr-FR"), DateFormatFull, "vendredi 30 janvier 2026"},

		{"de-DE short date", language.Make("de-DE"), DateFormatShort, "30.01.26"},
		{"de-DE long date", language.Make("de-DE"), DateFormatLong, "30. Januar 2026"},
		{"de-DE full date", language.Make("de-DE"), DateFormatFull, "Freitag, 30. Januar 2026"},

		{"es-ES long date", language.Make("es-ES"), DateFormatLong, "30 de enero de 2026"},
		{"es-ES medium date", language.Make("es-ES"), DateFormatMedium, "30 ene 2026"},

		{"it-IT long date", language.Make("it-IT"), DateFormatLong, "30 gennaio 2026"},

		{"pt-BR long date", language.Make("pt-BR"), DateFormatLong, "30 de janeiro de 2026"},

		{"nl-NL long date", language.Make("nl-NL"), DateFormatLong, "30 januari 2026"},

		{"ru-RU long date", language.Make("ru-RU"), DateFormatLong, "30 января 2026 г."},

		{"pl-PL long date", language.Make("pl-PL"), DateFormatLong, "30 stycznia 2026"},

		{"ja-JP long date", language.Make("ja-JP"), DateFormatLong, "2026年1月30日"},
		{"ja-JP full date", language.Make("ja-JP"), DateFormatFull, "2026年1月30日 金曜日"},

		{"zh-CN long date", language.Make("zh-CN"), DateFormatLong, "2026年1月30日"},
		{"zh-CN full date", language.Make("zh-CN"), DateFormatFull, "2026年1月30日 星期五"},

		{"ko-KR long date", language.Make("ko-KR"), DateFormatLong, "2026년 1월 30일"},
		{"ko-KR full date", language.Make("ko-KR"), DateFormatFull, "2026년 1월 30일 금요일"},

		// Region-only fallback: en-AU is registered, but "en-NZ" routes through aliases.
		{"en-NZ falls back to en-GB", language.Make("en-NZ"), DateFormatShort, "30/01/2026"},
		// Language-only fallback: "fr" → fr-FR.
		{"fr falls back to fr-FR", language.Make("fr"), DateFormatLong, "30 janvier 2026"},
		// Unknown locale → en-US.
		{"unknown falls back to en-US", language.Make("xx-XX"), DateFormatShort, "1/30/26"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDateTime(tc.tag, tc.style, d)
			if got != tc.want {
				t.Errorf("FormatDateTime(%s, %d) = %q, want %q", tc.tag, tc.style, got, tc.want)
			}
		})
	}
}

func TestFormatDateTimeForCountry(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	d := time.Date(2026, time.January, 30, 15, 4, 5, 0, loc)

	tests := []struct {
		name    string
		country string
		style   int
		want    string
	}{
		{"US long", "US", DateFormatLong, "January 30, 2026"},
		{"GB long", "GB", DateFormatLong, "30 January 2026"},
		{"FR long", "FR", DateFormatLong, "30 janvier 2026"},
		{"DE long", "DE", DateFormatLong, "30. Januar 2026"},
		{"JP long", "JP", DateFormatLong, "2026年1月30日"},
		{"lowercase normalised", "us", DateFormatShort, "1/30/26"},
		{"GB time short", "GB", TimeFormatShort, "15:04"},
		{"unknown country falls through to en-XX fallback", "PE", DateFormatLong, "January 30, 2026"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FormatDateTimeForCountry(d, tc.country, tc.style)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("FormatDateTimeForCountry(%q, %d) = %q, want %q", tc.country, tc.style, got, tc.want)
			}
		})
	}
}

func TestFormatDateTimeForCountry_Unknown(t *testing.T) {
	if _, err := FormatDateTimeForCountry(time.Now(), "ZZ", DateFormatLong); err == nil {
		t.Fatal("expected error for unsupported country code, got nil")
	}
}

func TestFormatDateTime_DateTime(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	d := time.Date(2026, time.January, 30, 15, 4, 5, 0, loc)

	tests := []struct {
		name  string
		tag   language.Tag
		style int
		want  string
	}{
		{"en-US datetime medium", language.Make("en-US"), DateTimeFormatMedium, "Jan 30, 2026, 3:04:05 PM"},
		{"en-GB datetime medium", language.Make("en-GB"), DateTimeFormatMedium, "30 Jan 2026, 15:04:05"},
		{"fr-FR datetime medium", language.Make("fr-FR"), DateTimeFormatMedium, "30 janv. 2026, 15:04:05"},
		{"de-DE datetime medium", language.Make("de-DE"), DateTimeFormatMedium, "30.01.2026, 15:04:05"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDateTime(tc.tag, tc.style, d)
			if got != tc.want {
				t.Errorf("FormatDateTime(%s) = %q, want %q", tc.tag, got, tc.want)
			}
		})
	}
}
