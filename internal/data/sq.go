// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sq_Locale = merge(nil, LocaleData{
	Name:      "sq",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"muajin e ardhshem": "in 1 month",
		"javen e ardhshme":  "in 1 week",
		"vitin e ardhshem":  "in 1 year",
		"muajin e kaluar":   "1 month ago",
		"javen e kaluar":    "1 week ago",
		"vitin e kaluar":    "1 year ago",
		"kete minute":       "0 minute ago",
		"e paradites":       "am",
		"e pasdites":        "pm",
		"kete muaj":         "0 month ago",
		"kete jave":         "0 week ago",
		"e merkure":         "wednesday",
		"kete ore":          "0 hour ago",
		"kete vit":          "0 year ago",
		"paradite":          "am",
		"e premte":          "friday",
		"e shtune":          "saturday",
		"dhjetor":           "december",
		"qershor":           "june",
		"pasdite":           "pm",
		"sekonde":           "second",
		"shtator":           "september",
		"e enjte":           "thursday",
		"e marte":           "tuesday",
		"shkurt":            "february",
		"korrik":            "july",
		"minute":            "minute",
		"e hene":            "monday",
		"nentor":            "november",
		"e diel":            "sunday",
		"prill":             "april",
		"gusht":             "august",
		"neser":             "in 1 day",
		"janar":             "january",
		"tetor":             "october",
		"tani":              "0 second ago",
		"dite":              "day",
		"mars":              "march",
		"muaj":              "month",
		"jave":              "week",
		"sot":               "0 day ago",
		"dje":               "1 day ago",
		"pri":               "april",
		"gsh":               "august",
		"dhj":               "december",
		"shk":               "february",
		"pre":               "friday",
		"gmt":               "gmt",
		"ore":               "hour",
		"jan":               "january",
		"kor":               "july",
		"qer":               "june",
		"mar":               "march",
		"maj":               "may",
		"min":               "minute",
		"hen":               "monday",
		"nen":               "november",
		"tet":               "october",
		"sek":               "second",
		"sht":               "september",
		"die":               "sunday",
		"enj":               "thursday",
		"utc":               "utc",
		"mer":               "wednesday",
		"vit":               "year",
		"am":                "am",
		"pm":                "pm",
		"'":                 "",
		",":                 "",
		";":                 "",
		"@":                 "",
		"[":                 "",
		"]":                 "",
		"|":                 "",
		" ":                 " ",
		"+":                 "+",
		"-":                 "-",
		".":                 ".",
		"/":                 "/",
		":":                 ":",
		"z":                 "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) sekonda me pare`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) sekonde me pare`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) minuta me pare`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) minute me pare`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pas (\d+) sekondash`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) dite me pare`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) muaj me pare`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) jave me pare`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) vjet me pare`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pas (\d+) minutash`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ore me pare`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) min me pare`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sek me pare`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) vit me pare`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pas (\d+) sekonde`), "in $1 second"},
		{regexp.MustCompile(`(?i)pas (\d+) vjetesh`), "in $1 year"},
		{regexp.MustCompile(`(?i)pas (\d+) ditesh`), "in $1 day"},
		{regexp.MustCompile(`(?i)pas (\d+) minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pas (\d+) muajsh`), "in $1 month"},
		{regexp.MustCompile(`(?i)pas (\d+) javesh`), "in $1 week"},
		{regexp.MustCompile(`(?i)pas (\d+) oresh`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pas (\d+) muaji`), "in $1 month"},
		{regexp.MustCompile(`(?i)pas (\d+) dite`), "in $1 day"},
		{regexp.MustCompile(`(?i)pas (\d+) jave`), "in $1 week"},
		{regexp.MustCompile(`(?i)pas (\d+) viti`), "in $1 year"},
		{regexp.MustCompile(`(?i)pas (\d+) ore`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pas (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pas (\d+) sek`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(\d+ sekonda me pare|\d+ sekonde me pare|\d+ minuta me pare|\d+ minute me pare|pas \d+ sekondash|\d+ dite me pare|\d+ jave me pare|\d+ muaj me pare|\d+ vjet me pare|pas \d+ minutash|\d+ min me pare|\d+ ore me pare|\d+ sek me pare|\d+ vit me pare|pas \d+ sekonde|pas \d+ vjetesh|pas \d+ ditesh|pas \d+ javesh|pas \d+ minute|pas \d+ muajsh|pas \d+ muaji|pas \d+ oresh|pas \d+ dite|pas \d+ jave|pas \d+ viti|pas \d+ min|pas \d+ ore|pas \d+ sek)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ sekonda me pare|\d+ sekonde me pare|\d+ minuta me pare|\d+ minute me pare|pas \d+ sekondash|\d+ dite me pare|\d+ jave me pare|\d+ muaj me pare|\d+ vjet me pare|pas \d+ minutash|\d+ min me pare|\d+ ore me pare|\d+ sek me pare|\d+ vit me pare|pas \d+ sekonde|pas \d+ vjetesh|pas \d+ ditesh|pas \d+ javesh|pas \d+ minute|pas \d+ muajsh|pas \d+ muaji|pas \d+ oresh|pas \d+ dite|pas \d+ jave|pas \d+ viti|pas \d+ min|pas \d+ ore|pas \d+ sek)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(muajin e ardhshem|javen e ardhshme|vitin e ardhshem|muajin e kaluar|javen e kaluar|vitin e kaluar|e paradites|kete minute|e pasdites|e merkure|kete jave|kete muaj|e premte|e shtune|kete ore|kete vit|paradite|dhjetor|e enjte|e marte|pasdite|qershor|sekonde|shtator|e diel|e hene|korrik|minute|nentor|shkurt|gusht|janar|neser|prill|tetor|dite|jave|mars|muaj|tani|dhj|die|dje|enj|gmt|gsh|hen|jan|kor|maj|mar|mer|min|nen|ore|pre|pri|qer|sek|shk|sht|sot|tet|utc|vit|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})

var sq_MK_Locale = merge(&sq_Locale, LocaleData{
	Name:      "sq-MK",
	DateOrder: "DMY",
})

var sq_XK_Locale = merge(&sq_Locale, LocaleData{
	Name:      "sq-XK",
	DateOrder: "DMY",
})
