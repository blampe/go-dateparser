// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var my_Locale = merge(nil, LocaleData{
	Name:      "my",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"ဖေဖောဝါရ": "february",
		"ကြာသပတေး": "thursday",
		"အောကတဘာ":  "october",
		"သောကြာ":   "friday",
		"ဇနနဝါရ":   "january",
		"စကတငဘာ":   "september",
		"တနငဂနေ":   "sunday",
		"ဒဇငဘာ":    "december",
		"တနငလာ":    "monday",
		"နဝငဘာ":    "november",
		"ဗဒဓဟး":    "wednesday",
		"အောက":     "october",
		"စကကန":     "second",
		"အငဂါ":     "tuesday",
		"နနက":      "am",
		"ဧပြ":      "april",
		"ဩဂတ":      "august",
		"gmt":      "gmt",
		"နာရ":      "hour",
		"ဇလင":      "july",
		"မနစ":      "minute",
		"ညနေ":      "pm",
		"စနေ":      "saturday",
		"utc":      "utc",
		"am":       "am",
		"ရက":       "day",
		"ဖေ":       "february",
		"ဇန":       "june",
		"မတ":       "march",
		"မေ":       "may",
		"pm":       "pm",
		"စက":       "september",
		"ပတ":       "week",
		"နစ":       "year",
		"'":        "",
		",":        "",
		";":        "",
		"@":        "",
		"[":        "",
		"]":        "",
		"|":        "",
		" ":        " ",
		"+":        "+",
		"-":        "-",
		".":        ".",
		"/":        "/",
		":":        ":",
		"ဧ":        "april",
		"ဩ":        "august",
		"ဒ":        "december",
		"ဇ":        "july",
		"လ":        "month",
		"န":        "november",
		"z":        "z",
	},
	RelativeType: map[string]string{
		"ပြးခသည သတငးပတ": "1 week ago",
		"လာမည သတငးပတ":   "in 1 week",
		"ယခ သတငးပတ":     "0 week ago",
		"ပြးခသညလ":       "1 month ago",
		"မနကဖြန":        "in 1 day",
		"လာမညနစ":        "in 1 year",
		"ဤအချန":         "0 hour ago",
		"ယမနနစ":         "1 year ago",
		"လာမညလ":         "in 1 month",
		"ဤမနစ":          "0 minute ago",
		"ယခနစ":          "0 year ago",
		"မနေက":          "1 day ago",
		"ယနေ":           "0 day ago",
		"ယခလ":           "0 month ago",
		"ယခ":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) စကကန`), "$1 second ago"},
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) နာရ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) မနစ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) ရက`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) ပတ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) နစ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)ပြးခသည (\d+) လ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) စကကနအတငး`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) နာရအတငး`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) မနစအတငး`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ရကအတငး`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ပတအတငး`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) နစအတငး`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) လအတငး`), "in $1 month"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(ပြးခသည \d+ စကကန|ပြးခသည \d+ နာရ|ပြးခသည \d+ မနစ|ပြးခသည \d+ နစ|ပြးခသည \d+ ပတ|ပြးခသည \d+ ရက|\d+ စကကနအတငး|ပြးခသည \d+ လ|\d+ နာရအတငး|\d+ မနစအတငး|\d+ နစအတငး|\d+ ပတအတငး|\d+ ရကအတငး|\d+ လအတငး)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(ပြးခသည \d+ စကကန|ပြးခသည \d+ နာရ|ပြးခသည \d+ မနစ|ပြးခသည \d+ နစ|ပြးခသည \d+ ပတ|ပြးခသည \d+ ရက|\d+ စကကနအတငး|ပြးခသည \d+ လ|\d+ နာရအတငး|\d+ မနစအတငး|\d+ နစအတငး|\d+ ပတအတငး|\d+ ရကအတငး|\d+ လအတငး)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ပြးခသည သတငးပတ|လာမည သတငးပတ|ယခ သတငးပတ|ကြာသပတေး|ဖေဖောဝါရ|ပြးခသညလ|အောကတဘာ|စကတငဘာ|ဇနနဝါရ|တနငဂနေ|မနကဖြန|လာမညနစ|သောကြာ|တနငလာ|ဒဇငဘာ|နဝငဘာ|ဗဒဓဟး|ယမနနစ|လာမညလ|ဤအချန|စကကန|မနေက|ယခနစ|အငဂါ|အောက|ဤမနစ|gmt|utc|စနေ|ဇလင|ညနေ|နနက|နာရ|မနစ|ယခလ|ယနေ|ဧပြ|ဩဂတ|\+|\.|\[|\]|\||am|pm|စက|ဇန|နစ|ပတ|ဖေ|မတ|မေ|ယခ|ရက| |'|,|-|/|:|;|@|z|ဇ|ဒ|န|လ|ဧ|ဩ)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
