// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var rwk_Locale = merge(nil, LocaleData{
	Name:      "rwk",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvwxyz`),
	Translations: map[string]string{
		"jumatatuu": "monday",
		"kyiukonyi": "pm",
		"jumapilyi": "sunday",
		"februari":  "february",
		"jumamosi":  "saturday",
		"septemba":  "september",
		"alhamisi":  "thursday",
		"jumatanu":  "wednesday",
		"aprilyi":   "april",
		"desemba":   "december",
		"januari":   "january",
		"dakyika":   "minute",
		"novemba":   "november",
		"sekunde":   "second",
		"jumanne":   "tuesday",
		"agusti":    "august",
		"ijumaa":    "friday",
		"julyai":    "july",
		"oktoba":    "october",
		"wiikyi":    "week",
		"utuko":     "am",
		"mfiri":     "day",
		"junyi":     "june",
		"machi":     "march",
		"mori":      "month",
		"maka":      "year",
		"apr":       "april",
		"ago":       "august",
		"des":       "december",
		"feb":       "february",
		"iju":       "friday",
		"gmt":       "gmt",
		"saa":       "hour",
		"jan":       "january",
		"jul":       "july",
		"jun":       "june",
		"mac":       "march",
		"mei":       "may",
		"jtt":       "monday",
		"nov":       "november",
		"okt":       "october",
		"jmo":       "saturday",
		"sep":       "september",
		"jpi":       "sunday",
		"alh":       "thursday",
		"jnn":       "tuesday",
		"utc":       "utc",
		"jtn":       "wednesday",
		"am":        "am",
		"pm":        "pm",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"z":         "z",
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
		"ngama":       "in 1 day",
		"ukou":        "1 day ago",
		"inu":         "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|jumapilyi|jumatatuu|kyiukonyi|last week|last year|next week|next year|this hour|this week|this year|alhamisi|februari|jumamosi|jumatanu|septemba|aprilyi|dakyika|desemba|januari|jumanne|novemba|sekunde|agusti|ijumaa|julyai|oktoba|wiikyi|junyi|machi|mfiri|ngama|utuko|maka|mori|ukou|ago|alh|apr|des|feb|gmt|iju|inu|jan|jmo|jnn|jpi|jtn|jtt|jul|jun|mac|mei|nov|now|okt|saa|sep|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
