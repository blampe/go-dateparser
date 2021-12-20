// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var nus_Locale = merge(nil, LocaleData{
	Name:      "nus",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"tiop thar pɛt": "january",
		"tiop in diit":  "december",
		"dhieec latni":  "friday",
		"bakɛl latni":   "saturday",
		"ŋuaan latni":   "thursday",
		"pay yietni":    "july",
		"diɔk latni":    "wednesday",
		"caŋ kuɔth":     "sunday",
		"rɛw latni":     "tuesday",
		"kornyoot":      "june",
		"jiec lat":      "monday",
		"thɛkɛni":       "second",
		"dhieec":        "friday",
		"thoor":         "august",
		"thaak":         "hour",
		"duɔɔŋ":         "march",
		"minit":         "minute",
		"laath":         "october",
		"bakɛl":         "saturday",
		"ŋuaan":         "thursday",
		"guak":          "april",
		"thoo":          "august",
		"tiop":          "january",
		"duɔɔ":          "march",
		"duat":          "may",
		"jiec":          "monday",
		"tɛɛr":          "september",
		"diɔk":          "wednesday",
		"jiɔk":          "week",
		"ruɔn":          "year",
		"caŋ":           "day",
		"tid":           "december",
		"pɛt":           "february",
		"gmt":           "gmt",
		"kor":           "june",
		"dua":           "may",
		"pay":           "month",
		"kur":           "november",
		"laa":           "october",
		"tɛɛ":           "september",
		"rɛw":           "tuesday",
		"utc":           "utc",
		"am":            "am",
		"rw":            "am",
		"pm":            "pm",
		"tŋ":            "pm",
		"'":             "",
		",":             "",
		";":             "",
		"@":             "",
		"[":             "",
		"]":             "",
		"|":             "",
		" ":             " ",
		"+":             "+",
		"-":             "-",
		".":             ".",
		"/":             "/",
		":":             ":",
		"z":             "z",
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
		"walɛ":        "0 day ago",
		"ruun":        "in 1 day",
		"now":         "0 second ago",
		"pan":         "1 day ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(tiop thar pɛt|dhieec latni|tiop in diit|bakɛl latni|this minute|ŋuaan latni|diɔk latni|last month|next month|pay yietni|this month|caŋ kuɔth|last week|last year|next week|next year|rɛw latni|this hour|this week|this year|jiec lat|kornyoot|thɛkɛni|dhieec|bakɛl|duɔɔŋ|laath|minit|thaak|thoor|ŋuaan|diɔk|duat|duɔɔ|guak|jiec|jiɔk|ruun|ruɔn|thoo|tiop|tɛɛr|walɛ|caŋ|dua|gmt|kor|kur|laa|now|pan|pay|pɛt|rɛw|tid|tɛɛ|utc|\+|\.|\[|\]|\||am|pm|rw|tŋ| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
