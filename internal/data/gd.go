// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var gd_Locale = merge(nil, LocaleData{
	Name:      "gd",
	DateOrder: "DMY",
	Charset:   []rune(`-bcdefghilnorstuz`),
	Translations: map[string]string{
		"dhen fhaoilleach": "january",
		"dhen dubhlachd":   "december",
		"dhen t-samhain":   "november",
		"dhen t-sultain":   "september",
		"dhen ghiblean":    "april",
		"dhen lunastal":    "august",
		"dhen ghearran":    "february",
		"am faoilleach":    "january",
		"dhen cheitean":    "may",
		"an dubhlachd":     "december",
		"uair a thide":     "hour",
		"an t-ogmhios":     "june",
		"dhen ogmhios":     "june",
		"an t-samhain":     "november",
		"dhen damhair":     "october",
		"an t-sultain":     "september",
		"an lunastal":      "august",
		"an t-iuchar":      "july",
		"dhen iuchar":      "july",
		"disathairne":      "saturday",
		"didomhnaich":      "sunday",
		"an giblean":       "april",
		"an gearran":       "february",
		"dhen mhart":       "march",
		"an ceitean":       "may",
		"an damhair":       "october",
		"diardaoin":        "thursday",
		"diciadain":        "wednesday",
		"seachdain":        "week",
		"dihaoine":         "friday",
		"bliadhna":         "year",
		"am mart":          "march",
		"mionaid":          "minute",
		"diluain":          "monday",
		"dimairt":          "tuesday",
		"seachd":           "week",
		"latha":            "day",
		"gearr":            "february",
		"gibl":             "april",
		"luna":             "august",
		"dubh":             "december",
		"uair":             "hour",
		"faoi":             "january",
		"iuch":             "july",
		"ogmh":             "june",
		"mart":             "march",
		"ceit":             "may",
		"mion":             "minute",
		"mios":             "month",
		"samh":             "november",
		"damh":             "october",
		"diog":             "second",
		"sult":             "september",
		"blia":             "year",
		"dih":              "friday",
		"gmt":              "gmt",
		"dil":              "monday",
		"dis":              "saturday",
		"did":              "sunday",
		"dia":              "thursday",
		"dim":              "tuesday",
		"utc":              "utc",
		"dic":              "wednesday",
		"am":               "am",
		"la":               "day",
		"mi":               "month",
		"pm":               "pm",
		"sn":               "week",
		"bl":               "year",
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
		"m":                "am",
		"u":                "hour",
		"f":                "pm",
		"d":                "second",
		"z":                "z",
	},
	RelativeType: map[string]string{
		"an t-seachdain seo chaidh": "1 week ago",
		"an t-seachdain seo":        "0 week ago",
		"am mios seo chaidh":        "1 month ago",
		"am mios sa chaidh":         "1 month ago",
		"an ath-sheachdain":         "in 1 week",
		"seachd sa chaidh":          "1 week ago",
		"an ath-bhliadhna":          "in 1 year",
		"an t-seachd seo":           "0 week ago",
		"an ath-sheachd":            "in 1 week",
		"an ath-mhios":              "in 1 month",
		"this minute":               "0 minute ago",
		"am mios seo":               "0 month ago",
		"an t-sn seo":               "0 week ago",
		"am bliadhna":               "0 year ago",
		"a-maireach":                "in 1 day",
		"an ath-bhl":                "in 1 year",
		"this hour":                 "0 hour ago",
		"am mi seo":                 "0 month ago",
		"an-drasta":                 "0 second ago",
		"an-uiridh":                 "1 year ago",
		"an-diugh":                  "0 day ago",
		"ath-mhi":                   "in 1 month",
		"ath-shn":                   "in 1 week",
		"an-uir":                    "1 year ago",
		"am bl":                     "0 year ago",
		"an-de":                     "1 day ago",
		"mi ch":                     "1 month ago",
		"sn ch":                     "1 week ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)an ceann (\d+) uair a thide`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) uair a thide air ais`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)an ceann (\d+) seachdain`), "in $1 week"},
		{regexp.MustCompile(`(?i)an ceann (\d+) bhliadhna`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) seachdain air ais`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) bhliadhna air ais`), "$1 year ago"},
		{regexp.MustCompile(`(?i)an ceann (\d+) mhionaid`), "in $1 minute"},
		{regexp.MustCompile(`(?i)an ceann (\d+) bliadhna`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) mhionaid air ais`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) bliadhna air ais`), "$1 year ago"},
		{regexp.MustCompile(`(?i)an ceann (\d+) mionaid`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) mionaid air ais`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)an ceann (\d+) mhiosa`), "in $1 month"},
		{regexp.MustCompile(`(?i)an ceann (\d+) latha`), "in $1 day"},
		{regexp.MustCompile(`(?i)an ceann (\d+) miosa`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) latha air ais`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) mhios air ais`), "$1 month ago"},
		{regexp.MustCompile(`(?i)an ceann (\d+) diog`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) mios air ais`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) diog air ais`), "$1 second ago"},
		{regexp.MustCompile(`(?i)an (\d+) sheachd`), "in $1 week"},
		{regexp.MustCompile(`(?i)o (\d+) sheachd`), "$1 week ago"},
		{regexp.MustCompile(`(?i)an (\d+) seachd`), "in $1 week"},
		{regexp.MustCompile(`(?i)o (\d+) seachd`), "$1 week ago"},
		{regexp.MustCompile(`(?i)an (\d+) mhion`), "in $1 minute"},
		{regexp.MustCompile(`(?i)an (\d+) mhios`), "in $1 month"},
		{regexp.MustCompile(`(?i)an (\d+) bhlia`), "in $1 year"},
		{regexp.MustCompile(`(?i)o (\d+) mhion`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)o (\d+) mhios`), "$1 month ago"},
		{regexp.MustCompile(`(?i)o (\d+) bhlia`), "$1 year ago"},
		{regexp.MustCompile(`(?i)an (\d+) uair`), "in $1 hour"},
		{regexp.MustCompile(`(?i)an (\d+) mion`), "in $1 minute"},
		{regexp.MustCompile(`(?i)an (\d+) mios`), "in $1 month"},
		{regexp.MustCompile(`(?i)an (\d+) diog`), "in $1 second"},
		{regexp.MustCompile(`(?i)an (\d+) blia`), "in $1 year"},
		{regexp.MustCompile(`(?i)o (\d+) uair`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)o (\d+) mion`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)o (\d+) mios`), "$1 month ago"},
		{regexp.MustCompile(`(?i)o (\d+) diog`), "$1 second ago"},
		{regexp.MustCompile(`(?i)o (\d+) blia`), "$1 year ago"},
		{regexp.MustCompile(`(?i)an (\d+) la`), "in $1 day"},
		{regexp.MustCompile(`(?i)o (\d+) la`), "$1 day ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(an ceann \d+ uair a thide|\d+ uair a thide air ais|an ceann \d+ bhliadhna|an ceann \d+ seachdain|\d+ bhliadhna air ais|\d+ seachdain air ais|an ceann \d+ bliadhna|an ceann \d+ mhionaid|\d+ bliadhna air ais|\d+ mhionaid air ais|an ceann \d+ mionaid|\d+ mionaid air ais|an ceann \d+ mhiosa|an ceann \d+ latha|an ceann \d+ miosa|\d+ latha air ais|\d+ mhios air ais|an ceann \d+ diog|\d+ diog air ais|\d+ mios air ais|an \d+ sheachd|an \d+ seachd|o \d+ sheachd|an \d+ bhlia|an \d+ mhion|an \d+ mhios|o \d+ seachd|an \d+ blia|an \d+ diog|an \d+ mion|an \d+ mios|an \d+ uair|o \d+ bhlia|o \d+ mhion|o \d+ mhios|o \d+ blia|o \d+ diog|o \d+ mion|o \d+ mios|o \d+ uair|an \d+ la|o \d+ la)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(an ceann \d+ uair a thide|\d+ uair a thide air ais|an ceann \d+ bhliadhna|an ceann \d+ seachdain|\d+ bhliadhna air ais|\d+ seachdain air ais|an ceann \d+ bliadhna|an ceann \d+ mhionaid|\d+ bliadhna air ais|\d+ mhionaid air ais|an ceann \d+ mionaid|\d+ mionaid air ais|an ceann \d+ mhiosa|an ceann \d+ latha|an ceann \d+ miosa|\d+ latha air ais|\d+ mhios air ais|an ceann \d+ diog|\d+ diog air ais|\d+ mios air ais|an \d+ sheachd|an \d+ seachd|o \d+ sheachd|an \d+ bhlia|an \d+ mhion|an \d+ mhios|o \d+ seachd|an \d+ blia|an \d+ diog|an \d+ mion|an \d+ mios|an \d+ uair|o \d+ bhlia|o \d+ mhion|o \d+ mhios|o \d+ blia|o \d+ diog|o \d+ mion|o \d+ mios|o \d+ uair|an \d+ la|o \d+ la)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(an t-seachdain seo chaidh|am mios seo chaidh|an t-seachdain seo|am mios sa chaidh|an ath-sheachdain|an ath-bhliadhna|dhen fhaoilleach|seachd sa chaidh|an t-seachd seo|an ath-sheachd|dhen dubhlachd|dhen t-samhain|dhen t-sultain|am faoilleach|dhen cheitean|dhen ghearran|dhen ghiblean|dhen lunastal|an ath-mhios|an dubhlachd|an t-ogmhios|an t-samhain|an t-sultain|dhen damhair|dhen ogmhios|uair a thide|am bliadhna|am mios seo|an lunastal|an t-iuchar|an t-sn seo|dhen iuchar|didomhnaich|disathairne|this minute|a-maireach|an ath-bhl|an ceitean|an damhair|an gearran|an giblean|dhen mhart|am mi seo|an-drasta|an-uiridh|diardaoin|diciadain|seachdain|this hour|an-diugh|bliadhna|dihaoine|am mart|ath-mhi|ath-shn|diluain|dimairt|mionaid|an-uir|seachd|am bl|an-de|gearr|latha|mi ch|sn ch|blia|ceit|damh|diog|dubh|faoi|gibl|iuch|luna|mart|mion|mios|ogmh|samh|sult|uair|dia|dic|did|dih|dil|dim|dis|gmt|utc|\+|\.|\[|\]|\||am|bl|la|mi|pm|sn| |'|,|-|/|:|;|@|d|f|m|u|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
