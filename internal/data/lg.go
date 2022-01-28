// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lg_Locale = merge(nil, LocaleData{
	Name:      "lg",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvwxyz`),
	Translations: map[string]string{
		"lwakutaano": "friday",
		"lwamukaaga": "saturday",
		"sebuttemba": "september",
		"febwaliyo":  "february",
		"janwaliyo":  "january",
		"kasikonda":  "second",
		"lwakubiri":  "tuesday",
		"lwakusatu":  "wednesday",
		"okitobba":   "october",
		"sabbiiti":   "week",
		"agusito":    "august",
		"desemba":    "december",
		"julaayi":    "july",
		"dakiika":    "minute",
		"novemba":    "november",
		"lwakuna":    "thursday",
		"lunaku":     "day",
		"marisi":     "march",
		"balaza":     "monday",
		"apuli":      "april",
		"saawa":      "hour",
		"juuni":      "june",
		"maayi":      "may",
		"mwezi":      "month",
		"mwaka":      "year",
		"apu":        "april",
		"agu":        "august",
		"des":        "december",
		"feb":        "february",
		"lw5":        "friday",
		"gmt":        "gmt",
		"jan":        "january",
		"jul":        "july",
		"juu":        "june",
		"mar":        "march",
		"maa":        "may",
		"bal":        "monday",
		"nov":        "november",
		"oki":        "october",
		"lw6":        "saturday",
		"seb":        "september",
		"sab":        "sunday",
		"lw4":        "thursday",
		"lw2":        "tuesday",
		"utc":        "utc",
		"lw3":        "wednesday",
		"am":         "am",
		"pm":         "pm",
		"'":          "",
		",":          "",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"|":          "",
		" ":          " ",
		"+":          "+",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		"z":          "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"lwaleero":    "0 day ago",
		"ggulo":       "1 day ago",
		"nkya":        "in 1 day",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|lwakutaano|lwamukaaga|next month|sebuttemba|this month|febwaliyo|janwaliyo|kasikonda|last week|last year|lwakubiri|lwakusatu|next week|next year|this hour|this week|this year|lwaleero|okitobba|sabbiiti|agusito|dakiika|desemba|julaayi|lwakuna|novemba|balaza|lunaku|marisi|apuli|ggulo|juuni|maayi|mwaka|mwezi|saawa|nkya|agu|apu|bal|des|feb|gmt|jan|jul|juu|lw2|lw3|lw4|lw5|lw6|maa|mar|nov|now|oki|sab|seb|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
