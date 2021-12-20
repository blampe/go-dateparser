// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var pa_Locale = merge(nil, LocaleData{
	Name:      "pa",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"ਸਨਿਚਰਵਾਰ": "saturday",
		"ਸਕਰਵਾਰ":   "friday",
		"ਵੀਰਵਾਰ":   "thursday",
		"ਮਗਲਵਾਰ":   "tuesday",
		"ਫਰਵਰੀ":    "february",
		"ਜਨਵਰੀ":    "january",
		"ਸਮਵਾਰ":    "monday",
		"ਮਹੀਨਾ":    "month",
		"ਅਕਤਬਰ":    "october",
		"ਸਨਿਚਰ":    "saturday",
		"ਐਤਵਾਰ":    "sunday",
		"ਬਧਵਾਰ":    "wednesday",
		"ਅਪਰਲ":     "april",
		"ਅਗਸਤ":     "august",
		"ਦਸਬਰ":     "december",
		"ਜਲਾਈ":     "july",
		"ਮਾਰਚ":     "march",
		"ਨਵਬਰ":     "november",
		"ਸਕਿਟ":     "second",
		"ਸਤਬਰ":     "september",
		"ਹਫਤਾ":     "week",
		"ਅਪਰ":      "april",
		"ਦਿਨ":      "day",
		"ਸਕਰ":      "friday",
		"gmt":      "gmt",
		"ਘਟਾ":      "hour",
		"ਜਲਾ":      "july",
		"ਮਿਟ":      "minute",
		"ਅਕਤ":      "october",
		"ਬਾਦ":      "pm",
		"ਵੀਰ":      "thursday",
		"ਮਗਲ":      "tuesday",
		"utc":      "utc",
		"ਸਾਲ":      "year",
		"am":       "am",
		"ਪਦ":       "am",
		"ਅਗ":       "august",
		"ਦਸ":       "december",
		"ਫਰ":       "february",
		"ਜਨ":       "june",
		"ਮਈ":       "may",
		"ਸਮ":       "monday",
		"ਨਵ":       "november",
		"pm":       "pm",
		"ਸਤ":       "september",
		"ਐਤ":       "sunday",
		"ਬਧ":       "wednesday",
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
		"ਘ":        "hour",
		"z":        "z",
	},
	RelativeType: map[string]string{
		"ਪਿਛਲਾ ਮਹੀਨਾ": "1 month ago",
		"ਪਿਛਲਾ ਹਫਤਾ":  "1 week ago",
		"ਅਗਲਾ ਮਹੀਨਾ":  "in 1 month",
		"ਬੀਤਿਆ ਕਲਹ":   "1 day ago",
		"ਪਿਛਲਾ ਸਾਲ":   "1 year ago",
		"ਅਗਲਾ ਹਫਤਾ":   "in 1 week",
		"ਇਹ ਮਹੀਨਾ":    "0 month ago",
		"ਅਗਲਾ ਸਾਲ":    "in 1 year",
		"ਇਹ ਹਫਤਾ":     "0 week ago",
		"ਇਸ ਮਿਟ":      "0 minute ago",
		"ਇਹ ਸਾਲ":      "0 year ago",
		"ਇਸ ਘਟ":       "0 hour ago",
		"ਭਲਕ":         "in 1 day",
		"ਅਜ":          "0 day ago",
		"ਹਣ":          "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) ਮਹੀਨਾ ਪਹਿਲਾ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ਮਹੀਨ ਪਹਿਲਾ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ਸਕਿਟ ਪਹਿਲਾ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) ਹਫਤਾ ਪਹਿਲਾ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ਮਹੀਨਿਆ ਵਿਚ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ਦਿਨ ਪਹਿਲਾ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) ਘਟਾ ਪਹਿਲਾ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ਮਿਟ ਪਹਿਲਾ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ਹਫਤ ਪਹਿਲਾ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ਸਾਲ ਪਹਿਲਾ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ਸਕਿਟਾ ਵਿਚ`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ਹਫਤਿਆ ਵਿਚ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ਘਟ ਪਹਿਲਾ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ਦਿਨਾ ਵਿਚ`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ਘਟਿਆ ਵਿਚ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) ਮਿਟਾ ਵਿਚ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ਮਹੀਨ ਵਿਚ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ਸਕਿਟ ਵਿਚ`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ਸਾਲਾ ਵਿਚ`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ਦਿਨ ਵਿਚ`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ਮਿਟ ਵਿਚ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ਹਫਤ ਵਿਚ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ਸਾਲ ਵਿਚ`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ਘਟ ਵਿਚ`), "in $1 hour"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ ਮਹੀਨਾ ਪਹਿਲਾ|\d+ ਮਹੀਨ ਪਹਿਲਾ|\d+ ਮਹੀਨਿਆ ਵਿਚ|\d+ ਸਕਿਟ ਪਹਿਲਾ|\d+ ਹਫਤਾ ਪਹਿਲਾ|\d+ ਘਟਾ ਪਹਿਲਾ|\d+ ਦਿਨ ਪਹਿਲਾ|\d+ ਮਿਟ ਪਹਿਲਾ|\d+ ਸਕਿਟਾ ਵਿਚ|\d+ ਸਾਲ ਪਹਿਲਾ|\d+ ਹਫਤ ਪਹਿਲਾ|\d+ ਹਫਤਿਆ ਵਿਚ|\d+ ਘਟ ਪਹਿਲਾ|\d+ ਘਟਿਆ ਵਿਚ|\d+ ਦਿਨਾ ਵਿਚ|\d+ ਮਹੀਨ ਵਿਚ|\d+ ਮਿਟਾ ਵਿਚ|\d+ ਸਕਿਟ ਵਿਚ|\d+ ਸਾਲਾ ਵਿਚ|\d+ ਦਿਨ ਵਿਚ|\d+ ਮਿਟ ਵਿਚ|\d+ ਸਾਲ ਵਿਚ|\d+ ਹਫਤ ਵਿਚ|\d+ ਘਟ ਵਿਚ)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ ਮਹੀਨਾ ਪਹਿਲਾ|\d+ ਮਹੀਨ ਪਹਿਲਾ|\d+ ਮਹੀਨਿਆ ਵਿਚ|\d+ ਸਕਿਟ ਪਹਿਲਾ|\d+ ਹਫਤਾ ਪਹਿਲਾ|\d+ ਘਟਾ ਪਹਿਲਾ|\d+ ਦਿਨ ਪਹਿਲਾ|\d+ ਮਿਟ ਪਹਿਲਾ|\d+ ਸਕਿਟਾ ਵਿਚ|\d+ ਸਾਲ ਪਹਿਲਾ|\d+ ਹਫਤ ਪਹਿਲਾ|\d+ ਹਫਤਿਆ ਵਿਚ|\d+ ਘਟ ਪਹਿਲਾ|\d+ ਘਟਿਆ ਵਿਚ|\d+ ਦਿਨਾ ਵਿਚ|\d+ ਮਹੀਨ ਵਿਚ|\d+ ਮਿਟਾ ਵਿਚ|\d+ ਸਕਿਟ ਵਿਚ|\d+ ਸਾਲਾ ਵਿਚ|\d+ ਦਿਨ ਵਿਚ|\d+ ਮਿਟ ਵਿਚ|\d+ ਸਾਲ ਵਿਚ|\d+ ਹਫਤ ਵਿਚ|\d+ ਘਟ ਵਿਚ)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ਪਿਛਲਾ ਮਹੀਨਾ|ਅਗਲਾ ਮਹੀਨਾ|ਪਿਛਲਾ ਹਫਤਾ|ਅਗਲਾ ਹਫਤਾ|ਪਿਛਲਾ ਸਾਲ|ਬੀਤਿਆ ਕਲਹ|ਅਗਲਾ ਸਾਲ|ਇਹ ਮਹੀਨਾ|ਸਨਿਚਰਵਾਰ|ਇਹ ਹਫਤਾ|ਇਸ ਮਿਟ|ਇਹ ਸਾਲ|ਮਗਲਵਾਰ|ਵੀਰਵਾਰ|ਸਕਰਵਾਰ|ਅਕਤਬਰ|ਇਸ ਘਟ|ਐਤਵਾਰ|ਜਨਵਰੀ|ਫਰਵਰੀ|ਬਧਵਾਰ|ਮਹੀਨਾ|ਸਨਿਚਰ|ਸਮਵਾਰ|ਅਗਸਤ|ਅਪਰਲ|ਜਲਾਈ|ਦਸਬਰ|ਨਵਬਰ|ਮਾਰਚ|ਸਕਿਟ|ਸਤਬਰ|ਹਫਤਾ|gmt|utc|ਅਕਤ|ਅਪਰ|ਘਟਾ|ਜਲਾ|ਦਿਨ|ਬਾਦ|ਭਲਕ|ਮਗਲ|ਮਿਟ|ਵੀਰ|ਸਕਰ|ਸਾਲ|\+|\.|\[|\]|\||am|pm|ਅਗ|ਅਜ|ਐਤ|ਜਨ|ਦਸ|ਨਵ|ਪਦ|ਫਰ|ਬਧ|ਮਈ|ਸਤ|ਸਮ|ਹਣ| |'|,|-|/|:|;|@|z|ਘ)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
