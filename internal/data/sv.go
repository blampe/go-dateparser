// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sv_Locale = merge(nil, LocaleData{
	Name:      "sv",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "den", "pa", "|"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)en(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)tva(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
	},
	Translations: map[string]string{
		"eftermiddag": "pm",
		"formiddag":   "am",
		"september":   "september",
		"december":    "december",
		"februari":    "february",
		"november":    "november",
		"sekunder":    "second",
		"augusti":     "august",
		"fran nu":     "in",
		"januari":     "january",
		"minuter":     "minute",
		"manaden":     "month",
		"manader":     "month",
		"oktober":     "october",
		"torsdag":     "thursday",
		"fredag":      "friday",
		"timmar":      "hour",
		"mandag":      "monday",
		"lordag":      "saturday",
		"sekund":      "second",
		"sondag":      "sunday",
		"tisdag":      "tuesday",
		"onsdag":      "wednesday",
		"veckor":      "week",
		"sedan":       "ago",
		"april":       "april",
		"dagar":       "day",
		"timme":       "hour",
		"minut":       "minute",
		"manad":       "month",
		"torsd":       "thursday",
		"vecka":       "week",
		"fred":        "friday",
		"juli":        "july",
		"juni":        "june",
		"mars":        "march",
		"mand":        "monday",
		"lord":        "saturday",
		"sept":        "september",
		"sond":        "sunday",
		"tors":        "thursday",
		"tisd":        "tuesday",
		"onsd":        "wednesday",
		"aret":        "year",
		"den":         "",
		"apr":         "april",
		"aug":         "august",
		"dag":         "day",
		"dec":         "december",
		"feb":         "february",
		"fre":         "friday",
		"gmt":         "gmt",
		"tim":         "hour",
		"jan":         "january",
		"jul":         "july",
		"jun":         "june",
		"maj":         "may",
		"min":         "minute",
		"man":         "month",
		"nov":         "november",
		"okt":         "october",
		"lor":         "saturday",
		"sek":         "second",
		"sep":         "september",
		"son":         "sunday",
		"tis":         "tuesday",
		"utc":         "utc",
		"ons":         "wednesday",
		"pa":          "",
		"am":          "am",
		"fm":          "am",
		"om":          "in",
		"em":          "pm",
		"pm":          "pm",
		"ar":          "year",
		"'":           "",
		",":           "",
		";":           "",
		"@":           "",
		"[":           "",
		"]":           "",
		"|":           "",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"h":           "hour",
		"t":           "hour",
		"m":           "minute",
		"s":           "second",
		"v":           "week",
		"z":           "z",
	},
	RelativeType: map[string]string{
		"forra manaden": "1 month ago",
		"forra veckan":  "1 week ago",
		"denna timme":   "0 hour ago",
		"denna minut":   "0 minute ago",
		"denna manad":   "0 month ago",
		"denna vecka":   "0 week ago",
		"nasta manad":   "in 1 month",
		"nasta vecka":   "in 1 week",
		"forra aret":    "1 year ago",
		"denna man":     "0 month ago",
		"forra man":     "1 month ago",
		"nasta man":     "in 1 month",
		"i morgon":      "in 1 day",
		"nasta ar":      "in 1 year",
		"denna v":       "0 week ago",
		"forra v":       "1 week ago",
		"forrgar":       "2 day ago",
		"imorgon":       "in 1 day",
		"nasta v":       "in 1 week",
		"i fjol":        "1 year ago",
		"i dag":         "0 day ago",
		"i gar":         "1 day ago",
		"idag":          "0 day ago",
		"i ar":          "0 year ago",
		"igar":          "1 day ago",
		"nu":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)for (\d+) sekunder sedan`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) minuter sedan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) manader sedan`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) timmar sedan`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+) sekund sedan`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) veckor sedan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)for (\d+) dagar sedan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)for (\d+) timme sedan`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+) minut sedan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) manad sedan`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) vecka sedan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)for (\d+) dag sedan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)for (\d+) tim sedan`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+) min sedan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) man sedan`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) sek sedan`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) ar sedan`), "$1 year ago"},
		{regexp.MustCompile(`(?i)for (\d+) d sedan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)for (\d+) v sedan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)om (\d+) sekunder`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) minuter`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) manader`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+) timmar`), "in $1 hour"},
		{regexp.MustCompile(`(?i)om (\d+) sekund`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) veckor`), "in $1 week"},
		{regexp.MustCompile(`(?i)om (\d+) dagar`), "in $1 day"},
		{regexp.MustCompile(`(?i)om (\d+) timme`), "in $1 hour"},
		{regexp.MustCompile(`(?i)om (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) manad`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+) vecka`), "in $1 week"},
		{regexp.MustCompile(`(?i)om (\d+) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)om (\d+) tim`), "in $1 hour"},
		{regexp.MustCompile(`(?i)om (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) man`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) ar`), "in $1 year"},
		{regexp.MustCompile(`(?i)−(\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)−(\d+) man`), "$1 month ago"},
		{regexp.MustCompile(`(?i)om (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)om (\d+) v`), "in $1 week"},
		{regexp.MustCompile(`(?i)−(\d+) ar`), "$1 year ago"},
		{regexp.MustCompile(`(?i)−(\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)−(\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)−(\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)−(\d+) v`), "$1 week ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+ sekunder sedan|for \d+ manader sedan|for \d+ minuter sedan|for \d+ sekund sedan|for \d+ timmar sedan|for \d+ veckor sedan|for \d+ dagar sedan|for \d+ manad sedan|for \d+ minut sedan|for \d+ timme sedan|for \d+ vecka sedan|for \d+ dag sedan|for \d+ man sedan|for \d+ min sedan|for \d+ sek sedan|for \d+ tim sedan|for \d+ ar sedan|for \d+ d sedan|for \d+ v sedan|om \d+ sekunder|om \d+ manader|om \d+ minuter|om \d+ sekund|om \d+ timmar|om \d+ veckor|om \d+ dagar|om \d+ manad|om \d+ minut|om \d+ timme|om \d+ vecka|om \d+ dag|om \d+ man|om \d+ min|om \d+ sek|om \d+ tim|om \d+ ar|om \d+ d|om \d+ v|−\d+ man|−\d+ min|−\d+ ar|−\d+ d|−\d+ h|−\d+ s|−\d+ v)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(for \d+ sekunder sedan|for \d+ manader sedan|for \d+ minuter sedan|for \d+ sekund sedan|for \d+ timmar sedan|for \d+ veckor sedan|for \d+ dagar sedan|for \d+ manad sedan|for \d+ minut sedan|for \d+ timme sedan|for \d+ vecka sedan|for \d+ dag sedan|for \d+ man sedan|for \d+ min sedan|for \d+ sek sedan|for \d+ tim sedan|for \d+ ar sedan|for \d+ d sedan|for \d+ v sedan|om \d+ sekunder|om \d+ manader|om \d+ minuter|om \d+ sekund|om \d+ timmar|om \d+ veckor|om \d+ dagar|om \d+ manad|om \d+ minut|om \d+ timme|om \d+ vecka|om \d+ dag|om \d+ man|om \d+ min|om \d+ sek|om \d+ tim|om \d+ ar|om \d+ d|om \d+ v|−\d+ man|−\d+ min|−\d+ ar|−\d+ d|−\d+ h|−\d+ s|−\d+ v)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(forra manaden|forra veckan|denna manad|denna minut|denna timme|denna vecka|eftermiddag|nasta manad|nasta vecka|forra aret|denna man|formiddag|forra man|nasta man|september|december|februari|i morgon|nasta ar|november|sekunder|augusti|denna v|forra v|forrgar|fran nu|imorgon|januari|manaden|manader|minuter|nasta v|oktober|torsdag|fredag|i fjol|lordag|mandag|onsdag|sekund|sondag|timmar|tisdag|veckor|april|dagar|i dag|i gar|manad|minut|sedan|timme|torsd|vecka|aret|fred|i ar|idag|igar|juli|juni|lord|mand|mars|onsd|sept|sond|tisd|tors|apr|aug|dag|dec|den|feb|fre|gmt|jan|jul|jun|lor|maj|man|min|nov|okt|ons|sek|sep|son|tim|tis|utc|\+|\.|\[|\]|\||am|ar|em|fm|nu|om|pa|pm| |'|,|-|/|:|;|@|h|m|s|t|v|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var sv_AX_Locale = merge(&sv_Locale, LocaleData{
	Name:      "sv-AX",
	DateOrder: "YMD",
})

var sv_FI_Locale = merge(&sv_Locale, LocaleData{
	Name:      "sv-FI",
	DateOrder: "DMY",
})
