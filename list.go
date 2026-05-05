package i18n

import (
	"strings"

	"golang.org/x/text/language"
)

type listPattern struct {
	separator string
	pairJoin  string
	finalJoin string
}

var defaultListPattern = listPattern{
	separator: ", ",
	pairJoin:  " and ",
	finalJoin: " and ",
}

var listPatterns = map[string]listPattern{
	"af": {separator: ", ", pairJoin: " en ", finalJoin: " en "},
	"ar": {separator: "، ", pairJoin: " و", finalJoin: " و"},
	"bg": {separator: ", ", pairJoin: " и ", finalJoin: " и "},
	"bn": {separator: ", ", pairJoin: " এবং ", finalJoin: " এবং "},
	"ca": {separator: ", ", pairJoin: " i ", finalJoin: " i "},
	"cs": {separator: ", ", pairJoin: " a ", finalJoin: " a "},
	"da": {separator: ", ", pairJoin: " og ", finalJoin: " og "},
	"de": {separator: ", ", pairJoin: " und ", finalJoin: " und "},
	"el": {separator: ", ", pairJoin: " και ", finalJoin: " και "},
	"en": {separator: ", ", pairJoin: " and ", finalJoin: " and "},
	"es": {separator: ", ", pairJoin: " y ", finalJoin: " y "},
	"et": {separator: ", ", pairJoin: " ja ", finalJoin: " ja "},
	"fa": {separator: "، ", pairJoin: " و ", finalJoin: " و "},
	"fi": {separator: ", ", pairJoin: " ja ", finalJoin: " ja "},
	"fr": {separator: ", ", pairJoin: " et ", finalJoin: " et "},
	"he": {separator: ", ", pairJoin: " ו", finalJoin: " ו"},
	"hi": {separator: ", ", pairJoin: " और ", finalJoin: " और "},
	"hr": {separator: ", ", pairJoin: " i ", finalJoin: " i "},
	"hu": {separator: ", ", pairJoin: " és ", finalJoin: " és "},
	"id": {separator: ", ", pairJoin: " dan ", finalJoin: " dan "},
	"it": {separator: ", ", pairJoin: " e ", finalJoin: " e "},
	"ja": {separator: "、", pairJoin: "と", finalJoin: "と"},
	"ko": {separator: ", ", pairJoin: " 및 ", finalJoin: " 및 "},
	"lt": {separator: ", ", pairJoin: " ir ", finalJoin: " ir "},
	"lv": {separator: ", ", pairJoin: " un ", finalJoin: " un "},
	"ms": {separator: ", ", pairJoin: " dan ", finalJoin: " dan "},
	"nb": {separator: ", ", pairJoin: " og ", finalJoin: " og "},
	"nl": {separator: ", ", pairJoin: " en ", finalJoin: " en "},
	"nn": {separator: ", ", pairJoin: " og ", finalJoin: " og "},
	"pl": {separator: ", ", pairJoin: " i ", finalJoin: " i "},
	"pt": {separator: ", ", pairJoin: " e ", finalJoin: " e "},
	"ro": {separator: ", ", pairJoin: " și ", finalJoin: " și "},
	"ru": {separator: ", ", pairJoin: " и ", finalJoin: " и "},
	"sk": {separator: ", ", pairJoin: " a ", finalJoin: " a "},
	"sl": {separator: ", ", pairJoin: " in ", finalJoin: " in "},
	"sr": {separator: ", ", pairJoin: " i ", finalJoin: " i "},
	"sv": {separator: ", ", pairJoin: " och ", finalJoin: " och "},
	"sw": {separator: ", ", pairJoin: " na ", finalJoin: " na "},
	"ta": {separator: ", ", pairJoin: " மற்றும் ", finalJoin: " மற்றும் "},
	"th": {separator: ", ", pairJoin: " และ", finalJoin: " และ"},
	"tr": {separator: ", ", pairJoin: " ve ", finalJoin: " ve "},
	"uk": {separator: ", ", pairJoin: " і ", finalJoin: " і "},
	"ur": {separator: "، ", pairJoin: " اور ", finalJoin: " اور "},
	"vi": {separator: ", ", pairJoin: " và ", finalJoin: " và "},
	"zh": {separator: "、", pairJoin: "和", finalJoin: "和"},
}

// LocalizedList joins items using list punctuation and conjunctions for the given language.
func LocalizedList(tag language.Tag, items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return items[0]
	}

	pattern := listPatternFor(tag)
	if len(items) == 2 {
		return items[0] + pattern.pairJoin + items[1]
	}

	return strings.Join(items[:len(items)-1], pattern.separator) + pattern.finalJoin + items[len(items)-1]
}

// TranslatedList translates each key and then joins the results using LocalizedList.
// If translator is nil it falls back to joining the provided keys directly.
func TranslatedList(tag language.Tag, translator Translator, keys []string) string {
	if translator == nil {
		return LocalizedList(tag, keys)
	}

	items := make([]string, len(keys))
	for i, key := range keys {
		items[i] = translator.Translate(key)
	}
	return LocalizedList(tag, items)
}

func listPatternFor(tag language.Tag) listPattern {
	base, _ := tag.Base()
	if pattern, ok := listPatterns[base.String()]; ok {
		return pattern
	}
	return defaultListPattern
}
