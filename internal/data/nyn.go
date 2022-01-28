// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var nyn_Locale = merge(nil, LocaleData{
	Name:      "nyn",
	DateOrder: "DMY",
	Charset:   []rune(`/bcdeghijklnorstuwxyz`),
	Translations: map[string]string{
		"okwaikumi na ibiri": "december",
		"okwaikumi na kumwe": "november",
		"obucweka/esekendi":  "second",
		"okwamushanju":       "july",
		"okwamunaana":        "august",
		"orwakataano":        "friday",
		"okwokubanza":        "january",
		"okwamukaaga":        "june",
		"okwakashatu":        "march",
		"okwakataana":        "may",
		"orwokubanza":        "monday",
		"orwamukaaga":        "saturday",
		"orwakashatu":        "wednesday",
		"okwakabiri":         "february",
		"okwamwenda":         "september",
		"orwakabiri":         "tuesday",
		"okwaikumi":          "october",
		"okwakana":           "april",
		"edakiika":           "minute",
		"orwakana":           "thursday",
		"eizooba":            "day",
		"shaaha":             "hour",
		"omwezi":             "month",
		"esande":             "week",
		"omwaka":             "year",
		"sande":              "sunday",
		"kkn":                "april",
		"kmn":                "august",
		"knb":                "december",
		"kbr":                "february",
		"okt":                "friday",
		"gmt":                "gmt",
		"kbz":                "january",
		"kms":                "july",
		"kmk":                "june",
		"kst":                "march",
		"ktn":                "may",
		"ork":                "monday",
		"knk":                "november",
		"kkm":                "october",
		"omk":                "saturday",
		"kmw":                "september",
		"san":                "sunday",
		"okn":                "thursday",
		"okb":                "tuesday",
		"utc":                "utc",
		"oks":                "wednesday",
		"am":                 "am",
		"pm":                 "pm",
		"'":                  "",
		",":                  "",
		";":                  "",
		"@":                  "",
		"[":                  "",
		"]":                  "",
		"|":                  "",
		" ":                  " ",
		"+":                  "+",
		"-":                  "-",
		".":                  ".",
		"/":                  "/",
		":":                  ":",
		"z":                  "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"nyomwabazyo": "1 day ago",
		"nyenkyakare": "in 1 day",
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
		"erizooba":    "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(okwaikumi na ibiri|okwaikumi na kumwe|obucweka/esekendi|okwamushanju|nyenkyakare|nyomwabazyo|okwakashatu|okwakataana|okwamukaaga|okwamunaana|okwokubanza|orwakashatu|orwakataano|orwamukaaga|orwokubanza|this minute|last month|next month|okwakabiri|okwamwenda|orwakabiri|this month|last week|last year|next week|next year|okwaikumi|this hour|this week|this year|edakiika|erizooba|okwakana|orwakana|eizooba|esande|omwaka|omwezi|shaaha|sande|gmt|kbr|kbz|kkm|kkn|kmk|kmn|kms|kmw|knb|knk|kst|ktn|now|okb|okn|oks|okt|omk|ork|san|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
