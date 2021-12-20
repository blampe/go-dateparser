// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var nmg_Locale = merge(nil, LocaleData{
	Name:      "nmg",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"sɔndɔ mafu malal": "wednesday",
		"mabaga ma sukul":  "friday",
		"ngwɛn wum navur":  "november",
		"sɔndɔ mafu mana":  "thursday",
		"sɔndɔ mafu maba":  "tuesday",
		"ngwɛn hɛmbuɛri":   "july",
		"ngwɛn matahra":    "january",
		"ngwɛn rɛbvua":     "september",
		"ngwɛn lɔmbi":      "august",
		"ngwɛn nmba":       "february",
		"ngwɛn ntuo":       "june",
		"ngwɛn nlal":       "march",
		"ngwɛn ntan":       "may",
		"ngwɛn nna":        "april",
		"ngwɛn wum":        "october",
		"krisimin":         "december",
		"sasadi":           "saturday",
		"mpala":            "minute",
		"mɔndɔ":            "monday",
		"ngwɛn":            "month",
		"nyiɛl":            "second",
		"sɔndɔ":            "week",
		"mana":             "am",
		"kris":             "december",
		"wula":             "hour",
		"ng11":             "november",
		"ng10":             "october",
		"kugu":             "pm",
		"mbvu":             "year",
		"ng4":              "april",
		"ng8":              "august",
		"duo":              "day",
		"ng2":              "february",
		"mbs":              "friday",
		"gmt":              "gmt",
		"ng1":              "january",
		"ng7":              "july",
		"ng6":              "june",
		"ng3":              "march",
		"ng5":              "may",
		"mɔn":              "monday",
		"sas":              "saturday",
		"ng9":              "september",
		"sɔn":              "sunday",
		"smn":              "thursday",
		"smb":              "tuesday",
		"utc":              "utc",
		"sml":              "wednesday",
		"am":               "am",
		"pm":               "pm",
		"'":                "",
		",":                "",
		";":                "",
		"@":                "",
		"[":                "",
		"]":                "",
		"|":                "",
		" ":                " ",
		"+":                "+",
		"-":                "-",
		".":                ".",
		"/":                "/",
		":":                ":",
		"z":                "z",
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
		"nakugu":      "1 day ago",
		"namana":      "in 1 day",
		"dɔl":         "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(sɔndɔ mafu malal|mabaga ma sukul|ngwɛn wum navur|sɔndɔ mafu maba|sɔndɔ mafu mana|ngwɛn hɛmbuɛri|ngwɛn matahra|ngwɛn rɛbvua|ngwɛn lɔmbi|this minute|last month|next month|ngwɛn nlal|ngwɛn nmba|ngwɛn ntan|ngwɛn ntuo|this month|last week|last year|next week|next year|ngwɛn nna|ngwɛn wum|this hour|this week|this year|krisimin|nakugu|namana|sasadi|mpala|mɔndɔ|ngwɛn|nyiɛl|sɔndɔ|kris|kugu|mana|mbvu|ng10|ng11|wula|duo|dɔl|gmt|mbs|mɔn|ng1|ng2|ng3|ng4|ng5|ng6|ng7|ng8|ng9|now|sas|smb|sml|smn|sɔn|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
