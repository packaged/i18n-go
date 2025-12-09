package i18n

import "golang.org/x/text/language"

type Direction string

const (
	RTL Direction = "rtl"
	LTR Direction = "ltr"
)

func GetDirection(lang language.Tag) Direction {
	// Normalize the tag to its base language for simpler checking
	base, _ := lang.Base()

	switch base.String() {
	case "ar", "he", "fa", "ur", "az", "ps", "dv", "yi":
		// Arabic, Hebrew, Persian, Urdu, Azerbaijani (in Iran), Pashto, Dhivehi, Yiddish
		return RTL
	}

	// Check specific script/dialect tags if the base is ambiguous
	switch lang.String() {
	case "ckb": // Kurdish (Sorani) - uses Arabic script
		return RTL
	}

	return LTR
}
