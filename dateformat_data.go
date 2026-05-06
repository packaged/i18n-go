package i18n

// Localized name slices used by entries in dateLocaleData.
// Patterns themselves live inline in the map literal below so each locale
// reads top-to-bottom without jumping to a separate definition.
var (
	monthsFR = [12]string{
		"janvier", "février", "mars", "avril", "mai", "juin",
		"juillet", "août", "septembre", "octobre", "novembre", "décembre",
	}
	monthsShortFR = [12]string{
		"janv.", "févr.", "mars", "avr.", "mai", "juin",
		"juil.", "août", "sept.", "oct.", "nov.", "déc.",
	}
	weekdaysFR      = [7]string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"}
	weekdaysShortFR = [7]string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."}

	monthsDE = [12]string{
		"Januar", "Februar", "März", "April", "Mai", "Juni",
		"Juli", "August", "September", "Oktober", "November", "Dezember",
	}
	monthsShortDE = [12]string{
		"Jan.", "Feb.", "März", "Apr.", "Mai", "Juni",
		"Juli", "Aug.", "Sept.", "Okt.", "Nov.", "Dez.",
	}
	weekdaysDE      = [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"}
	weekdaysShortDE = [7]string{"So.", "Mo.", "Di.", "Mi.", "Do.", "Fr.", "Sa."}

	monthsES = [12]string{
		"enero", "febrero", "marzo", "abril", "mayo", "junio",
		"julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
	}
	monthsShortES = [12]string{
		"ene", "feb", "mar", "abr", "may", "jun",
		"jul", "ago", "sept", "oct", "nov", "dic",
	}
	weekdaysES      = [7]string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"}
	weekdaysShortES = [7]string{"dom", "lun", "mar", "mié", "jue", "vie", "sáb"}

	monthsIT = [12]string{
		"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno",
		"luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre",
	}
	monthsShortIT = [12]string{
		"gen", "feb", "mar", "apr", "mag", "giu",
		"lug", "ago", "set", "ott", "nov", "dic",
	}
	weekdaysIT      = [7]string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"}
	weekdaysShortIT = [7]string{"dom", "lun", "mar", "mer", "gio", "ven", "sab"}

	monthsPT = [12]string{
		"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
	}
	monthsShortPT = [12]string{
		"jan.", "fev.", "mar.", "abr.", "mai.", "jun.",
		"jul.", "ago.", "set.", "out.", "nov.", "dez.",
	}
	weekdaysPT      = [7]string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"}
	weekdaysShortPT = [7]string{"dom.", "seg.", "ter.", "qua.", "qui.", "sex.", "sáb."}

	monthsNL = [12]string{
		"januari", "februari", "maart", "april", "mei", "juni",
		"juli", "augustus", "september", "oktober", "november", "december",
	}
	monthsShortNL = [12]string{
		"jan", "feb", "mrt", "apr", "mei", "jun",
		"jul", "aug", "sep", "okt", "nov", "dec",
	}
	weekdaysNL      = [7]string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"}
	weekdaysShortNL = [7]string{"zo", "ma", "di", "wo", "do", "vr", "za"}

	// Russian: month names use the genitive case (the date-context form).
	monthsRU = [12]string{
		"января", "февраля", "марта", "апреля", "мая", "июня",
		"июля", "августа", "сентября", "октября", "ноября", "декабря",
	}
	monthsShortRU = [12]string{
		"янв.", "февр.", "мар.", "апр.", "мая", "июн.",
		"июл.", "авг.", "сент.", "окт.", "нояб.", "дек.",
	}
	weekdaysRU      = [7]string{"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"}
	weekdaysShortRU = [7]string{"вс", "пн", "вт", "ср", "чт", "пт", "сб"}

	// Polish: month names use the genitive case.
	monthsPL = [12]string{
		"stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca",
		"lipca", "sierpnia", "września", "października", "listopada", "grudnia",
	}
	monthsShortPL = [12]string{
		"sty", "lut", "mar", "kwi", "maj", "cze",
		"lip", "sie", "wrz", "paź", "lis", "gru",
	}
	weekdaysPL      = [7]string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"}
	weekdaysShortPL = [7]string{"niedz.", "pon.", "wt.", "śr.", "czw.", "pt.", "sob."}

	monthsSV = [12]string{
		"januari", "februari", "mars", "april", "maj", "juni",
		"juli", "augusti", "september", "oktober", "november", "december",
	}
	monthsShortSV = [12]string{
		"jan.", "feb.", "mars", "apr.", "maj", "juni",
		"juli", "aug.", "sep.", "okt.", "nov.", "dec.",
	}
	weekdaysSV      = [7]string{"söndag", "måndag", "tisdag", "onsdag", "torsdag", "fredag", "lördag"}
	weekdaysShortSV = [7]string{"sön", "mån", "tis", "ons", "tors", "fre", "lör"}

	monthsDA = [12]string{
		"januar", "februar", "marts", "april", "maj", "juni",
		"juli", "august", "september", "oktober", "november", "december",
	}
	monthsShortDA = [12]string{
		"jan.", "feb.", "mar.", "apr.", "maj", "juni",
		"juli", "aug.", "sep.", "okt.", "nov.", "dec.",
	}
	weekdaysDA      = [7]string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}
	weekdaysShortDA = [7]string{"søn.", "man.", "tir.", "ons.", "tor.", "fre.", "lør."}

	monthsNB = [12]string{
		"januar", "februar", "mars", "april", "mai", "juni",
		"juli", "august", "september", "oktober", "november", "desember",
	}
	monthsShortNB = [12]string{
		"jan.", "feb.", "mars", "apr.", "mai", "juni",
		"juli", "aug.", "sep.", "okt.", "nov.", "des.",
	}
	weekdaysNB      = [7]string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}
	weekdaysShortNB = [7]string{"søn.", "man.", "tir.", "ons.", "tor.", "fre.", "lør."}

	monthsFI = [12]string{
		"tammikuuta", "helmikuuta", "maaliskuuta", "huhtikuuta", "toukokuuta", "kesäkuuta",
		"heinäkuuta", "elokuuta", "syyskuuta", "lokakuuta", "marraskuuta", "joulukuuta",
	}
	monthsShortFI = [12]string{
		"tammik.", "helmik.", "maalisk.", "huhtik.", "toukok.", "kesäk.",
		"heinäk.", "elok.", "syysk.", "lokak.", "marrask.", "jouluk.",
	}
	weekdaysFI      = [7]string{"sunnuntai", "maanantai", "tiistai", "keskiviikko", "torstai", "perjantai", "lauantai"}
	weekdaysShortFI = [7]string{"su", "ma", "ti", "ke", "to", "pe", "la"}

	monthsTR = [12]string{
		"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran",
		"Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık",
	}
	monthsShortTR = [12]string{
		"Oca", "Şub", "Mar", "Nis", "May", "Haz",
		"Tem", "Ağu", "Eyl", "Eki", "Kas", "Ara",
	}
	weekdaysTR      = [7]string{"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi"}
	weekdaysShortTR = [7]string{"Paz", "Pzt", "Sal", "Çar", "Per", "Cum", "Cmt"}

	// Asian locales: patterns use numeric months with "月" / "월" suffix, so
	// the month-name slices are not consulted. Only weekdays are localized.
	weekdaysJA      = [7]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
	weekdaysShortJA = [7]string{"日", "月", "火", "水", "木", "金", "土"}

	weekdaysZHCN      = [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	weekdaysShortZHCN = [7]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}

	weekdaysZHTW      = [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	weekdaysShortZHTW = [7]string{"週日", "週一", "週二", "週三", "週四", "週五", "週六"}

	weekdaysKO      = [7]string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}
	weekdaysShortKO = [7]string{"일", "월", "화", "수", "목", "금", "토"}
)

// dateLocaleData holds CLDR-derived patterns and localized names per BCP-47 tag.
//
// Patterns are Go time layouts. Where a locale leaves Months/Weekdays empty,
// time.Format's English output is preserved.
//
// Coverage is intentionally focused on the locales rwd-go's country map needs.
// To add a locale, add a map entry plus (optionally) name slices above. The
// init() at the bottom registers language-only and regional aliases.
var dateLocaleData = map[string]LocaleDateData{
	"en-US": {
		DateShort:      "1/2/06",
		DateMedium:     "Jan 2, 2006",
		DateLong:       "January 2, 2006",
		DateFull:       "Monday, January 2, 2006",
		TimeShort:      "3:04 PM",
		TimeMedium:     "3:04:05 PM",
		TimeLong:       "3:04:05 PM MST",
		TimeFull:       "3:04:05 PM MST",
		DateTimeShort:  "1/2/06, 3:04 PM",
		DateTimeMedium: "Jan 2, 2006, 3:04:05 PM",
		DateTimeLong:   "January 2, 2006 at 3:04:05 PM MST",
		DateTimeFull:   "Monday, January 2, 2006 at 3:04:05 PM MST",
	},
	"en-GB": {
		DateShort:      "02/01/2006",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday, 2 January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02/01/2006, 15:04",
		DateTimeMedium: "2 Jan 2006, 15:04:05",
		DateTimeLong:   "2 January 2006 at 15:04:05 MST",
		DateTimeFull:   "Monday, 2 January 2006 at 15:04:05 MST",
	},
	"en-CA": {
		DateShort:      "2006-01-02",
		DateMedium:     "Jan 2, 2006",
		DateLong:       "January 2, 2006",
		DateFull:       "Monday, January 2, 2006",
		TimeShort:      "3:04 PM",
		TimeMedium:     "3:04:05 PM",
		TimeLong:       "3:04:05 PM MST",
		TimeFull:       "3:04:05 PM MST",
		DateTimeShort:  "2006-01-02, 3:04 PM",
		DateTimeMedium: "Jan 2, 2006, 3:04:05 PM",
		DateTimeLong:   "January 2, 2006 at 3:04:05 PM MST",
		DateTimeFull:   "Monday, January 2, 2006 at 3:04:05 PM MST",
	},
	"en-AU": {
		DateShort:      "2/1/2006",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday, 2 January 2006",
		TimeShort:      "3:04 PM",
		TimeMedium:     "3:04:05 PM",
		TimeLong:       "3:04:05 PM MST",
		TimeFull:       "3:04:05 PM MST",
		DateTimeShort:  "2/1/2006, 3:04 PM",
		DateTimeMedium: "2 Jan 2006, 3:04:05 PM",
		DateTimeLong:   "2 January 2006 at 3:04:05 PM MST",
		DateTimeFull:   "Monday, 2 January 2006 at 3:04:05 PM MST",
	},

	"fr-FR": {
		DateShort:      "02/01/2006",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday 2 January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02/01/2006 15:04",
		DateTimeMedium: "2 Jan 2006, 15:04:05",
		DateTimeLong:   "2 January 2006 à 15:04:05 MST",
		DateTimeFull:   "Monday 2 January 2006 à 15:04:05 MST",
		Months:         monthsFR,
		MonthsShort:    monthsShortFR,
		Weekdays:       weekdaysFR,
		WeekdaysShort:  weekdaysShortFR,
	},

	"de-DE": {
		DateShort:      "02.01.06",
		DateMedium:     "02.01.2006",
		DateLong:       "2. January 2006",
		DateFull:       "Monday, 2. January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02.01.06, 15:04",
		DateTimeMedium: "02.01.2006, 15:04:05",
		DateTimeLong:   "2. January 2006 um 15:04:05 MST",
		DateTimeFull:   "Monday, 2. January 2006 um 15:04:05 MST",
		Months:         monthsDE,
		MonthsShort:    monthsShortDE,
		Weekdays:       weekdaysDE,
		WeekdaysShort:  weekdaysShortDE,
	},

	"es-ES": {
		DateShort:      "2/1/06",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 de January de 2006",
		DateFull:       "Monday, 2 de January de 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2/1/06 15:04",
		DateTimeMedium: "2 Jan 2006, 15:04:05",
		DateTimeLong:   "2 de January de 2006, 15:04:05 MST",
		DateTimeFull:   "Monday, 2 de January de 2006, 15:04:05 MST",
		Months:         monthsES,
		MonthsShort:    monthsShortES,
		Weekdays:       weekdaysES,
		WeekdaysShort:  weekdaysShortES,
	},

	"es-MX": {
		DateShort:      "2/1/06",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 de January de 2006",
		DateFull:       "Monday, 2 de January de 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2/1/06, 15:04",
		DateTimeMedium: "2 Jan 2006, 15:04:05",
		DateTimeLong:   "2 de January de 2006, 15:04:05 MST",
		DateTimeFull:   "Monday, 2 de January de 2006, 15:04:05 MST",
		Months:         monthsES,
		MonthsShort:    monthsShortES,
		Weekdays:       weekdaysES,
		WeekdaysShort:  weekdaysShortES,
	},

	"it-IT": {
		DateShort:      "02/01/06",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday 2 January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02/01/06, 15:04",
		DateTimeMedium: "2 Jan 2006, 15:04:05",
		DateTimeLong:   "2 January 2006 alle 15:04:05 MST",
		DateTimeFull:   "Monday 2 January 2006 alle 15:04:05 MST",
		Months:         monthsIT,
		MonthsShort:    monthsShortIT,
		Weekdays:       weekdaysIT,
		WeekdaysShort:  weekdaysShortIT,
	},

	"pt-BR": {
		DateShort:      "02/01/2006",
		DateMedium:     "2 de Jan de 2006",
		DateLong:       "2 de January de 2006",
		DateFull:       "Monday, 2 de January de 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02/01/2006 15:04",
		DateTimeMedium: "2 de Jan de 2006 15:04:05",
		DateTimeLong:   "2 de January de 2006 às 15:04:05 MST",
		DateTimeFull:   "Monday, 2 de January de 2006 às 15:04:05 MST",
		Months:         monthsPT,
		MonthsShort:    monthsShortPT,
		Weekdays:       weekdaysPT,
		WeekdaysShort:  weekdaysShortPT,
	},

	"pt-PT": {
		DateShort:      "02/01/06",
		DateMedium:     "2 de Jan de 2006",
		DateLong:       "2 de January de 2006",
		DateFull:       "Monday, 2 de January de 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02/01/06, 15:04",
		DateTimeMedium: "2 de Jan de 2006, 15:04:05",
		DateTimeLong:   "2 de January de 2006 às 15:04:05 MST",
		DateTimeFull:   "Monday, 2 de January de 2006 às 15:04:05 MST",
		Months:         monthsPT,
		MonthsShort:    monthsShortPT,
		Weekdays:       weekdaysPT,
		WeekdaysShort:  weekdaysShortPT,
	},

	"nl-NL": {
		DateShort:      "02-01-2006",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday 2 January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02-01-2006 15:04",
		DateTimeMedium: "2 Jan 2006 15:04:05",
		DateTimeLong:   "2 January 2006 om 15:04:05 MST",
		DateTimeFull:   "Monday 2 January 2006 om 15:04:05 MST",
		Months:         monthsNL,
		MonthsShort:    monthsShortNL,
		Weekdays:       weekdaysNL,
		WeekdaysShort:  weekdaysShortNL,
	},

	"ru-RU": {
		DateShort:      "02.01.2006",
		DateMedium:     "2 Jan 2006 г.",
		DateLong:       "2 January 2006 г.",
		DateFull:       "Monday, 2 January 2006 г.",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02.01.2006, 15:04",
		DateTimeMedium: "2 Jan 2006 г., 15:04:05",
		DateTimeLong:   "2 January 2006 г., 15:04:05 MST",
		DateTimeFull:   "Monday, 2 January 2006 г., 15:04:05 MST",
		Months:         monthsRU,
		MonthsShort:    monthsShortRU,
		Weekdays:       weekdaysRU,
		WeekdaysShort:  weekdaysShortRU,
	},

	"pl-PL": {
		DateShort:      "2.01.2006",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday, 2 January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2.01.2006, 15:04",
		DateTimeMedium: "2 Jan 2006, 15:04:05",
		DateTimeLong:   "2 January 2006, 15:04:05 MST",
		DateTimeFull:   "Monday, 2 January 2006, 15:04:05 MST",
		Months:         monthsPL,
		MonthsShort:    monthsShortPL,
		Weekdays:       weekdaysPL,
		WeekdaysShort:  weekdaysShortPL,
	},

	"sv-SE": {
		DateShort:      "2006-01-02",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "Monday 2 January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2006-01-02 15:04",
		DateTimeMedium: "2 Jan 2006 15:04:05",
		DateTimeLong:   "2 January 2006 kl. 15:04:05 MST",
		DateTimeFull:   "Monday 2 January 2006 kl. 15:04:05 MST",
		Months:         monthsSV,
		MonthsShort:    monthsShortSV,
		Weekdays:       weekdaysSV,
		WeekdaysShort:  weekdaysShortSV,
	},

	"da-DK": {
		DateShort:      "02.01.2006",
		DateMedium:     "2. Jan 2006",
		DateLong:       "2. January 2006",
		DateFull:       "Monday den 2. January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02.01.2006 15:04",
		DateTimeMedium: "2. Jan 2006 15:04:05",
		DateTimeLong:   "2. January 2006 kl. 15:04:05 MST",
		DateTimeFull:   "Monday den 2. January 2006 kl. 15:04:05 MST",
		Months:         monthsDA,
		MonthsShort:    monthsShortDA,
		Weekdays:       weekdaysDA,
		WeekdaysShort:  weekdaysShortDA,
	},

	"nb-NO": {
		DateShort:      "02.01.2006",
		DateMedium:     "2. Jan 2006",
		DateLong:       "2. January 2006",
		DateFull:       "Monday 2. January 2006",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "02.01.2006, 15:04",
		DateTimeMedium: "2. Jan 2006, 15:04:05",
		DateTimeLong:   "2. January 2006 kl. 15:04:05 MST",
		DateTimeFull:   "Monday 2. January 2006 kl. 15:04:05 MST",
		Months:         monthsNB,
		MonthsShort:    monthsShortNB,
		Weekdays:       weekdaysNB,
		WeekdaysShort:  weekdaysShortNB,
	},

	"fi-FI": {
		DateShort:      "2.1.2006",
		DateMedium:     "2. Jan 2006",
		DateLong:       "2. January 2006",
		DateFull:       "Monday 2. January 2006",
		TimeShort:      "15.04",
		TimeMedium:     "15.04.05",
		TimeLong:       "15.04.05 MST",
		TimeFull:       "15.04.05 MST",
		DateTimeShort:  "2.1.2006 15.04",
		DateTimeMedium: "2. Jan 2006 15.04.05",
		DateTimeLong:   "2. January 2006 klo 15.04.05 MST",
		DateTimeFull:   "Monday 2. January 2006 klo 15.04.05 MST",
		Months:         monthsFI,
		MonthsShort:    monthsShortFI,
		Weekdays:       weekdaysFI,
		WeekdaysShort:  weekdaysShortFI,
	},

	"tr-TR": {
		DateShort:      "2.01.2006",
		DateMedium:     "2 Jan 2006",
		DateLong:       "2 January 2006",
		DateFull:       "2 January 2006 Monday",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2.01.2006 15:04",
		DateTimeMedium: "2 Jan 2006 15:04:05",
		DateTimeLong:   "2 January 2006 15:04:05 MST",
		DateTimeFull:   "2 January 2006 Monday 15:04:05 MST",
		Months:         monthsTR,
		MonthsShort:    monthsShortTR,
		Weekdays:       weekdaysTR,
		WeekdaysShort:  weekdaysShortTR,
	},

	"ja-JP": {
		DateShort:      "2006/01/02",
		DateMedium:     "2006/01/02",
		DateLong:       "2006年1月2日",
		DateFull:       "2006年1月2日 Monday",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2006/01/02 15:04",
		DateTimeMedium: "2006/01/02 15:04:05",
		DateTimeLong:   "2006年1月2日 15:04:05 MST",
		DateTimeFull:   "2006年1月2日 Monday 15:04:05 MST",
		Weekdays:       weekdaysJA,
		WeekdaysShort:  weekdaysShortJA,
	},

	"zh-CN": {
		DateShort:      "2006/1/2",
		DateMedium:     "2006年1月2日",
		DateLong:       "2006年1月2日",
		DateFull:       "2006年1月2日 Monday",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2006/1/2 15:04",
		DateTimeMedium: "2006年1月2日 15:04:05",
		DateTimeLong:   "2006年1月2日 15:04:05 MST",
		DateTimeFull:   "2006年1月2日 Monday 15:04:05 MST",
		Weekdays:       weekdaysZHCN,
		WeekdaysShort:  weekdaysShortZHCN,
	},

	"zh-TW": {
		DateShort:      "2006/1/2",
		DateMedium:     "2006年1月2日",
		DateLong:       "2006年1月2日",
		DateFull:       "2006年1月2日 Monday",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2006/1/2 15:04",
		DateTimeMedium: "2006年1月2日 15:04:05",
		DateTimeLong:   "2006年1月2日 15:04:05 MST",
		DateTimeFull:   "2006年1月2日 Monday 15:04:05 MST",
		Weekdays:       weekdaysZHTW,
		WeekdaysShort:  weekdaysShortZHTW,
	},

	"ko-KR": {
		DateShort:      "2006. 1. 2.",
		DateMedium:     "2006. 1. 2.",
		DateLong:       "2006년 1월 2일",
		DateFull:       "2006년 1월 2일 Monday",
		TimeShort:      "15:04",
		TimeMedium:     "15:04:05",
		TimeLong:       "15:04:05 MST",
		TimeFull:       "15:04:05 MST",
		DateTimeShort:  "2006. 1. 2. 15:04",
		DateTimeMedium: "2006. 1. 2. 15:04:05",
		DateTimeLong:   "2006년 1월 2일 15:04:05 MST",
		DateTimeFull:   "2006년 1월 2일 Monday 15:04:05 MST",
		Weekdays:       weekdaysKO,
		WeekdaysShort:  weekdaysShortKO,
	},
}

// init wires up region and language-only aliases so lookups for tags like
// "en", "fr-CA", "de-AT", or "no-NO" resolve to a sensible base entry without
// duplicating data.
func init() {
	aliases := map[string][]string{
		"en-US": {"en"},
		"en-GB": {"en-IE", "en-NZ", "en-IN", "en-ZA", "en-SG", "en-PH", "en-MT", "en-HK"},
		"fr-FR": {"fr", "fr-CA", "fr-BE", "fr-CH", "fr-LU"},
		"de-DE": {"de", "de-AT", "de-CH", "de-LI", "de-LU"},
		"es-ES": {"es", "es-AR", "es-CL", "es-CO", "es-PE", "es-UY", "es-VE", "es-CR", "es-DO"},
		"it-IT": {"it", "it-CH"},
		"pt-PT": {"pt"},
		"nl-NL": {"nl", "nl-BE"},
		"ru-RU": {"ru"},
		"pl-PL": {"pl"},
		"sv-SE": {"sv"},
		"da-DK": {"da"},
		"nb-NO": {"no", "no-NO", "nb", "nn-NO", "nn"},
		"fi-FI": {"fi"},
		"tr-TR": {"tr"},
		"ja-JP": {"ja"},
		"zh-CN": {"zh", "zh-Hans", "zh-Hans-CN", "zh-SG"},
		"zh-TW": {"zh-Hant", "zh-Hant-TW", "zh-HK", "zh-MO"},
		"ko-KR": {"ko"},
	}
	for src, targets := range aliases {
		data, ok := dateLocaleData[src]
		if !ok {
			continue
		}
		for _, t := range targets {
			if _, exists := dateLocaleData[t]; !exists {
				dateLocaleData[t] = data
			}
		}
	}
}
