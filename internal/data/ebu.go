// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ebu_Locale = merge(nil, LocaleData{
	Name:      "ebu",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"mweri wa ikumi na kairi": "december",
		"mweri wa ikumi na umwe":  "november",
		"mweri wa gatantatu":      "june",
		"mweri wa mugwanja":       "july",
		"mweri wa kathatu":        "march",
		"mweri wa kanana":         "august",
		"mweri wa gatano":         "may",
		"mweri wa kairi":          "february",
		"mweri wa mbere":          "january",
		"mweri wa ikumi":          "october",
		"mweri wa kenda":          "september",
		"mweri wa kana":           "april",
		"njumamothii":             "saturday",
		"njumatatu":               "monday",
		"njumatano":               "wednesday",
		"muthenya":                "day",
		"aramithi":                "thursday",
		"njumaine":                "tuesday",
		"ndagika":                 "minute",
		"sekondi":                 "second",
		"njumaa":                  "friday",
		"kiumia":                  "week",
		"ithaa":                   "hour",
		"mweri":                   "month",
		"mwaka":                   "year",
		"kan":                     "april",
		"knn":                     "august",
		"igi":                     "december",
		"kai":                     "february",
		"maa":                     "friday",
		"gmt":                     "gmt",
		"mbe":                     "january",
		"mug":                     "july",
		"gan":                     "june",
		"kat":                     "march",
		"gat":                     "may",
		"tat":                     "monday",
		"imw":                     "november",
		"iku":                     "october",
		"nmm":                     "saturday",
		"ken":                     "september",
		"kma":                     "sunday",
		"arm":                     "thursday",
		"ine":                     "tuesday",
		"utc":                     "utc",
		"tan":                     "wednesday",
		"am":                      "am",
		"ki":                      "am",
		"pm":                      "pm",
		"ut":                      "pm",
		"'":                       "",
		",":                       "",
		";":                       "",
		"@":                       "",
		"[":                       "",
		"]":                       "",
		"|":                       "",
		" ":                       " ",
		"+":                       "+",
		"-":                       "-",
		".":                       ".",
		"/":                       "/",
		":":                       ":",
		"z":                       "z",
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
		"umunthi":     "0 day ago",
		"igoro":       "1 day ago",
		"ruciu":       "in 1 day",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(mweri wa ikumi na kairi|mweri wa ikumi na umwe|mweri wa gatantatu|mweri wa mugwanja|mweri wa kathatu|mweri wa gatano|mweri wa kanana|mweri wa ikumi|mweri wa kairi|mweri wa kenda|mweri wa mbere|mweri wa kana|njumamothii|this minute|last month|next month|this month|last week|last year|next week|next year|njumatano|njumatatu|this hour|this week|this year|aramithi|muthenya|njumaine|ndagika|sekondi|umunthi|kiumia|njumaa|igoro|ithaa|mwaka|mweri|ruciu|arm|gan|gat|gmt|igi|iku|imw|ine|kai|kan|kat|ken|kma|knn|maa|mbe|mug|nmm|now|tan|tat|utc|\+|\.|\[|\]|\||am|ki|pm|ut| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
