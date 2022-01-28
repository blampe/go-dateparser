// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var se_Locale = merge(nil, LocaleData{
	Name:      "se",
	DateOrder: "YMD",
	Charset:   []rune(`bcdeghijklnorstuvwxyzđŋ`),
	Translations: map[string]string{
		"ođđajagemannu": "january",
		"eahketbeaivet": "pm",
		"suoidnemannu":  "july",
		"golggotmannu":  "october",
		"eahketbeaivi":  "pm",
		"iđitbeaivet":   "am",
		"juovlamannu":   "december",
		"guovvamannu":   "february",
		"geassemannu":   "june",
		"njukcamannu":   "march",
		"miessemannu":   "may",
		"skabmamannu":   "november",
		"sotnabeaivi":   "sunday",
		"gaskavahkku":   "wednesday",
		"iđitbeaivi":    "am",
		"cuoŋomannu":    "april",
		"borgemannu":    "august",
		"cakcamannu":    "september",
		"maŋŋebarga":    "tuesday",
		"bearjadat":     "friday",
		"vuossarga":     "monday",
		"lavvardat":     "saturday",
		"duorasdat":     "thursday",
		"minuhtta":      "minute",
		"sekunda":       "second",
		"beaivi":        "day",
		"diibmu":        "hour",
		"vahkku":        "week",
		"mannu":         "month",
		"jahki":         "year",
		"borg":          "august",
		"juov":          "december",
		"guov":          "february",
		"bear":          "friday",
		"ođđj":          "january",
		"suoi":          "july",
		"geas":          "june",
		"njuk":          "march",
		"mies":          "may",
		"vuos":          "monday",
		"skab":          "november",
		"golg":          "october",
		"cakc":          "september",
		"sotn":          "sunday",
		"duor":          "thursday",
		"gask":          "wednesday",
		"cuo":           "april",
		"gmt":           "gmt",
		"lav":           "saturday",
		"maŋ":           "tuesday",
		"utc":           "utc",
		"am":            "am",
		"ib":            "am",
		"eb":            "pm",
		"pm":            "pm",
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
		"ihttin":      "in 1 day",
		"odne":        "0 day ago",
		"ikte":        "1 day ago",
		"now":         "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) manotbadji maŋŋilit`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) minuhtta maŋŋilit`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) sekundda maŋŋilit`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) jandora maŋŋilit`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) diibmur maŋŋilit`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) minuhta maŋŋilit`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) sekunda maŋŋilit`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) manotbadji arat`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) jandor maŋŋilit`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) diibmu maŋŋilit`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) vahkku maŋŋilit`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) jahkki maŋŋilit`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) vahku maŋŋilit`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) jahki maŋŋilit`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) minuhtta arat`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sekundda arat`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) jandora arat`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) diibmur arat`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) minuhta arat`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sekunda arat`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) jandor arat`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) diibmu arat`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) vahkku arat`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) jahkki arat`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) vahku arat`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) jahki arat`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ manotbadji maŋŋilit|\d+ minuhtta maŋŋilit|\d+ sekundda maŋŋilit|\d+ diibmur maŋŋilit|\d+ jandora maŋŋilit|\d+ minuhta maŋŋilit|\d+ sekunda maŋŋilit|\d+ diibmu maŋŋilit|\d+ jahkki maŋŋilit|\d+ jandor maŋŋilit|\d+ manotbadji arat|\d+ vahkku maŋŋilit|\d+ jahki maŋŋilit|\d+ vahku maŋŋilit|\d+ minuhtta arat|\d+ sekundda arat|\d+ diibmur arat|\d+ jandora arat|\d+ minuhta arat|\d+ sekunda arat|\d+ diibmu arat|\d+ jahkki arat|\d+ jandor arat|\d+ vahkku arat|\d+ jahki arat|\d+ vahku arat)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ manotbadji maŋŋilit|\d+ minuhtta maŋŋilit|\d+ sekundda maŋŋilit|\d+ diibmur maŋŋilit|\d+ jandora maŋŋilit|\d+ minuhta maŋŋilit|\d+ sekunda maŋŋilit|\d+ diibmu maŋŋilit|\d+ jahkki maŋŋilit|\d+ jandor maŋŋilit|\d+ manotbadji arat|\d+ vahkku maŋŋilit|\d+ jahki maŋŋilit|\d+ vahku maŋŋilit|\d+ minuhtta arat|\d+ sekundda arat|\d+ diibmur arat|\d+ jandora arat|\d+ minuhta arat|\d+ sekunda arat|\d+ diibmu arat|\d+ jahkki arat|\d+ jandor arat|\d+ vahkku arat|\d+ jahki arat|\d+ vahku arat)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(eahketbeaivet|ođđajagemannu|eahketbeaivi|golggotmannu|suoidnemannu|gaskavahkku|geassemannu|guovvamannu|iđitbeaivet|juovlamannu|miessemannu|njukcamannu|skabmamannu|sotnabeaivi|this minute|borgemannu|cakcamannu|cuoŋomannu|iđitbeaivi|last month|maŋŋebarga|next month|this month|bearjadat|duorasdat|last week|last year|lavvardat|next week|next year|this hour|this week|this year|vuossarga|minuhtta|sekunda|beaivi|diibmu|ihttin|vahkku|jahki|mannu|bear|borg|cakc|duor|gask|geas|golg|guov|ikte|juov|mies|njuk|odne|ođđj|skab|sotn|suoi|vuos|cuo|gmt|lav|maŋ|now|utc|\+|\.|\[|\]|\||am|eb|ib|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var se_FI_Locale = merge(&se_Locale, LocaleData{
	Name:      "se-FI",
	DateOrder: "YMD",
	Translations: map[string]string{
		"maŋŋebargga": "tuesday",
		"bearjadaga":  "friday",
		"vuossargga":  "monday",
		"lavvardaga":  "saturday",
		"duorastaga":  "thursday",
		"gaskavahku":  "wednesday",
		"j":           "year",
	},
	RelativeType: map[string]string{
		"boahtte jagi": "in 1 year",
		"mannan jagi":  "1 year ago",
		"dan jagi":     "0 year ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) jagi siste`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) jagi arat`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ manotbadji maŋŋilit|\d+ minuhtta maŋŋilit|\d+ sekundda maŋŋilit|\d+ diibmur maŋŋilit|\d+ jandora maŋŋilit|\d+ minuhta maŋŋilit|\d+ sekunda maŋŋilit|\d+ diibmu maŋŋilit|\d+ jahkki maŋŋilit|\d+ jandor maŋŋilit|\d+ manotbadji arat|\d+ vahkku maŋŋilit|\d+ jahki maŋŋilit|\d+ vahku maŋŋilit|\d+ minuhtta arat|\d+ sekundda arat|\d+ diibmur arat|\d+ jandora arat|\d+ minuhta arat|\d+ sekunda arat|\d+ diibmu arat|\d+ jahkki arat|\d+ jandor arat|\d+ vahkku arat|\d+ jagi siste|\d+ jahki arat|\d+ vahku arat|\d+ jagi arat)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ manotbadji maŋŋilit|\d+ minuhtta maŋŋilit|\d+ sekundda maŋŋilit|\d+ diibmur maŋŋilit|\d+ jandora maŋŋilit|\d+ minuhta maŋŋilit|\d+ sekunda maŋŋilit|\d+ diibmu maŋŋilit|\d+ jahkki maŋŋilit|\d+ jandor maŋŋilit|\d+ manotbadji arat|\d+ vahkku maŋŋilit|\d+ jahki maŋŋilit|\d+ vahku maŋŋilit|\d+ minuhtta arat|\d+ sekundda arat|\d+ diibmur arat|\d+ jandora arat|\d+ minuhta arat|\d+ sekunda arat|\d+ diibmu arat|\d+ jahkki arat|\d+ jandor arat|\d+ vahkku arat|\d+ jagi siste|\d+ jahki arat|\d+ vahku arat|\d+ jagi arat)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(eahketbeaivet|ođđajagemannu|boahtte jagi|eahketbeaivi|golggotmannu|suoidnemannu|gaskavahkku|geassemannu|guovvamannu|iđitbeaivet|juovlamannu|mannan jagi|maŋŋebargga|miessemannu|njukcamannu|skabmamannu|sotnabeaivi|this minute|bearjadaga|borgemannu|cakcamannu|cuoŋomannu|duorastaga|gaskavahku|iđitbeaivi|last month|lavvardaga|maŋŋebarga|next month|this month|vuossargga|bearjadat|duorasdat|last week|last year|lavvardat|next week|next year|this hour|this week|this year|vuossarga|dan jagi|minuhtta|sekunda|beaivi|diibmu|ihttin|vahkku|jahki|mannu|bear|borg|cakc|duor|gask|geas|golg|guov|ikte|juov|mies|njuk|odne|ođđj|skab|sotn|suoi|vuos|cuo|gmt|lav|maŋ|now|utc|\+|\.|\[|\]|\||am|eb|ib|pm| |'|,|-|/|:|;|@|j|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var se_SE_Locale = merge(&se_Locale, LocaleData{
	Name:            "se-SE",
	DateOrder:       "YMD",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ manotbadji maŋŋilit|\d+ minuhtta maŋŋilit|\d+ sekundda maŋŋilit|\d+ diibmur maŋŋilit|\d+ jandora maŋŋilit|\d+ minuhta maŋŋilit|\d+ sekunda maŋŋilit|\d+ diibmu maŋŋilit|\d+ jahkki maŋŋilit|\d+ jandor maŋŋilit|\d+ manotbadji arat|\d+ vahkku maŋŋilit|\d+ jahki maŋŋilit|\d+ vahku maŋŋilit|\d+ minuhtta arat|\d+ sekundda arat|\d+ diibmur arat|\d+ jandora arat|\d+ minuhta arat|\d+ sekunda arat|\d+ diibmu arat|\d+ jahkki arat|\d+ jandor arat|\d+ vahkku arat|\d+ jahki arat|\d+ vahku arat)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ manotbadji maŋŋilit|\d+ minuhtta maŋŋilit|\d+ sekundda maŋŋilit|\d+ diibmur maŋŋilit|\d+ jandora maŋŋilit|\d+ minuhta maŋŋilit|\d+ sekunda maŋŋilit|\d+ diibmu maŋŋilit|\d+ jahkki maŋŋilit|\d+ jandor maŋŋilit|\d+ manotbadji arat|\d+ vahkku maŋŋilit|\d+ jahki maŋŋilit|\d+ vahku maŋŋilit|\d+ minuhtta arat|\d+ sekundda arat|\d+ diibmur arat|\d+ jandora arat|\d+ minuhta arat|\d+ sekunda arat|\d+ diibmu arat|\d+ jahkki arat|\d+ jandor arat|\d+ vahkku arat|\d+ jahki arat|\d+ vahku arat)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(eahketbeaivet|ođđajagemannu|eahketbeaivi|golggotmannu|suoidnemannu|gaskavahkku|geassemannu|guovvamannu|iđitbeaivet|juovlamannu|miessemannu|njukcamannu|skabmamannu|sotnabeaivi|this minute|borgemannu|cakcamannu|cuoŋomannu|iđitbeaivi|last month|maŋŋebarga|next month|this month|bearjadat|duorasdat|last week|last year|lavvardat|next week|next year|this hour|this week|this year|vuossarga|minuhtta|sekunda|beaivi|diibmu|ihttin|vahkku|jahki|mannu|bear|borg|cakc|duor|gask|geas|golg|guov|ikte|juov|mies|njuk|odne|ođđj|skab|sotn|suoi|vuos|cuo|gmt|lav|maŋ|now|utc|\+|\.|\[|\]|\||am|eb|ib|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
