package i18n

import (
	"strings"
	"time"

	"golang.org/x/text/language"
)

// Date / time / datetime style constants.
// Numeric values are not part of the API; do not persist them.
const (
	DateFormatFull = iota
	DateFormatLong
	DateFormatMedium
	DateFormatShort
	TimeFormatFull
	TimeFormatLong
	TimeFormatMedium
	TimeFormatShort
	DateTimeFormatFull
	DateTimeFormatLong
	DateTimeFormatMedium
	DateTimeFormatShort
)

// LocaleDateData carries CLDR-derived patterns and localized names for a locale.
//
// Patterns are Go time layouts (using the reference time Mon Jan 2 15:04:05 MST 2006).
// Months/Weekdays are post-format substitutions for the English values produced by
// time.Format. Leave any element empty to keep the English value.
type LocaleDateData struct {
	DateFull   string
	DateLong   string
	DateMedium string
	DateShort  string

	TimeFull   string
	TimeLong   string
	TimeMedium string
	TimeShort  string

	DateTimeFull   string
	DateTimeLong   string
	DateTimeMedium string
	DateTimeShort  string

	Months        [12]string // January..December
	MonthsShort   [12]string // Jan..Dec
	Weekdays      [7]string  // Sunday..Saturday
	WeekdaysShort [7]string  // Sun..Sat
}

var englishMonths = [12]string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

var englishMonthsShort = [12]string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun",
	"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}

var englishWeekdays = [7]string{
	"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
}

var englishWeekdaysShort = [7]string{
	"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
}

// FormatDateTime formats t for tag using the given style constant.
// Unknown locales fall back to en-US patterns and English names.
func FormatDateTime(tag language.Tag, style int, t time.Time) string {
	data := LookupDateLocale(tag)
	pattern := stylePattern(data, style)
	if pattern == "" {
		pattern = data.DateMedium
	}
	return localizedFormat(t, pattern, data)
}

// LookupDateLocale resolves the LocaleDateData for tag, walking down to the
// language-only entry and finally en-US.
func LookupDateLocale(tag language.Tag) LocaleDateData {
	if d, ok := dateLocaleData[tag.String()]; ok {
		return d
	}
	base, _ := tag.Base()
	region, _ := tag.Region()
	for _, k := range []string{
		base.String() + "-" + region.String(),
		base.String(),
	} {
		if d, ok := dateLocaleData[k]; ok {
			return d
		}
	}
	return dateLocaleData["en-US"]
}

func stylePattern(d LocaleDateData, style int) string {
	switch style {
	case DateFormatFull:
		return d.DateFull
	case DateFormatLong:
		return d.DateLong
	case DateFormatMedium:
		return d.DateMedium
	case DateFormatShort:
		return d.DateShort
	case TimeFormatFull:
		return d.TimeFull
	case TimeFormatLong:
		return d.TimeLong
	case TimeFormatMedium:
		return d.TimeMedium
	case TimeFormatShort:
		return d.TimeShort
	case DateTimeFormatFull:
		return d.DateTimeFull
	case DateTimeFormatLong:
		return d.DateTimeLong
	case DateTimeFormatMedium:
		return d.DateTimeMedium
	case DateTimeFormatShort:
		return d.DateTimeShort
	}
	return d.DateMedium
}

const (
	sentinelMonthFull    = "\x01\x02M\x02\x01"
	sentinelMonthShort   = "\x01\x02m\x02\x01"
	sentinelWeekdayFull  = "\x01\x02W\x02\x01"
	sentinelWeekdayShort = "\x01\x02w\x02\x01"
)

// localizedFormat formats t with pattern then swaps the English month/weekday
// names that time.Format produced for the localized equivalents in d.
//
// The substitution uses sentinels to avoid the classic collision where a long
// localized month contains the English short prefix (e.g. German "Januar"
// would otherwise re-match "Jan").
func localizedFormat(t time.Time, pattern string, d LocaleDateData) string {
	s := t.Format(pattern)

	enM := englishMonths[t.Month()-1]
	enMS := englishMonthsShort[t.Month()-1]
	enW := englishWeekdays[t.Weekday()]
	enWS := englishWeekdaysShort[t.Weekday()]

	// Replace long forms first; their English values contain the short forms.
	if strings.Contains(s, enM) {
		s = strings.Replace(s, enM, sentinelMonthFull, 1)
	}
	if strings.Contains(s, enW) {
		s = strings.Replace(s, enW, sentinelWeekdayFull, 1)
	}
	if strings.Contains(s, enMS) {
		s = strings.Replace(s, enMS, sentinelMonthShort, 1)
	}
	if strings.Contains(s, enWS) {
		s = strings.Replace(s, enWS, sentinelWeekdayShort, 1)
	}

	s = strings.Replace(s, sentinelMonthFull, pickName(d.Months[t.Month()-1], enM), 1)
	s = strings.Replace(s, sentinelMonthShort, pickName(d.MonthsShort[t.Month()-1], enMS), 1)
	s = strings.Replace(s, sentinelWeekdayFull, pickName(d.Weekdays[t.Weekday()], enW), 1)
	s = strings.Replace(s, sentinelWeekdayShort, pickName(d.WeekdaysShort[t.Weekday()], enWS), 1)
	return s
}

func pickName(localized, english string) string {
	if localized == "" {
		return english
	}
	return localized
}
