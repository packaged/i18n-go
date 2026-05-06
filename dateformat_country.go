package i18n

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/language"
)

// countryDateLocale maps an ISO 3166-1 alpha-2 country code to the BCP-47
// locale that drives date and time formatting. Unknown countries fall back
// to "en-<COUNTRY>", matching PHP's IntlDateFormatter behaviour.
var countryDateLocale = map[string]string{
	"US": "en-US", "GB": "en-GB", "FR": "fr-FR", "DE": "de-DE",
	"ES": "es-ES", "IT": "it-IT", "JP": "ja-JP", "CN": "zh-CN",
	"KR": "ko-KR", "RU": "ru-RU", "BR": "pt-BR", "MX": "es-MX",
	"CA": "en-CA", "AU": "en-AU", "NZ": "en-NZ", "IN": "hi-IN",
	"ZA": "en-ZA", "NL": "nl-NL", "BE": "nl-BE", "CH": "de-CH",
	"AT": "de-AT", "PL": "pl-PL", "SE": "sv-SE", "NO": "nb-NO",
	"DK": "da-DK", "FI": "fi-FI", "PT": "pt-PT", "GR": "el-GR",
	"TR": "tr-TR", "CZ": "cs-CZ", "HU": "hu-HU", "RO": "ro-RO",
	"IL": "he-IL", "SA": "ar-SA", "AE": "ar-AE", "TH": "th-TH",
	"VN": "vi-VN", "ID": "id-ID", "MY": "ms-MY", "SG": "en-SG",
	"PH": "en-PH", "HK": "zh-HK", "TW": "zh-TW", "AR": "es-AR",
	"CL": "es-CL", "CO": "es-CO", "IE": "en-IE",
}

// FormatDateTimeForCountry renders t using the date/time conventions of the
// given ISO 3166-1 alpha-2 country code. style must be one of the
// DateFormat*, TimeFormat*, or DateTimeFormat* constants.
//
// Returns an error if the country code is not recognised. The supported set
// is taken from CountryCurrencies, which covers every assigned ISO code.
func FormatDateTimeForCountry(t time.Time, countryCode string, style int) (string, error) {
	cc := strings.ToUpper(strings.TrimSpace(countryCode))
	if _, ok := CountryCurrencies[cc]; !ok {
		return "", fmt.Errorf("%s is not a supported country", countryCode)
	}

	locale, ok := countryDateLocale[cc]
	if !ok {
		locale = "en-" + cc
	}
	return FormatDateTime(language.Make(locale), style, t), nil
}
