// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var nl_Locale = merge(nil, LocaleData{
	Name:      "nl",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefgijklnorstuvwz`),
	Translations: map[string]string{
		"september": "september",
		"donderdag": "thursday",
		"augustus":  "august",
		"december":  "december",
		"februari":  "february",
		"november":  "november",
		"zaterdag":  "saturday",
		"seconden":  "second",
		"woensdag":  "wednesday",
		"geleden":   "ago",
		"vrijdag":   "friday",
		"januari":   "january",
		"minuten":   "minute",
		"maandag":   "monday",
		"maanden":   "month",
		"oktober":   "october",
		"seconde":   "second",
		"dinsdag":   "tuesday",
		"minuut":    "minute",
		"zondag":    "sunday",
		"april":     "april",
		"dagen":     "day",
		"maart":     "march",
		"maand":     "month",
		"weken":     "week",
		"juli":      "july",
		"juni":      "june",
		"week":      "week",
		"jaar":      "year",
		"apr":       "april",
		"aug":       "august",
		"dag":       "day",
		"dec":       "december",
		"feb":       "february",
		"gmt":       "gmt",
		"uur":       "hour",
		"jan":       "january",
		"jul":       "july",
		"jun":       "june",
		"mrt":       "march",
		"mei":       "may",
		"min":       "minute",
		"mnd":       "month",
		"nov":       "november",
		"okt":       "october",
		"sec":       "second",
		"sep":       "september",
		"utc":       "utc",
		"om":        "",
		"am":        "am",
		"vr":        "friday",
		"in":        "in",
		"ma":        "monday",
		"pm":        "pm",
		"za":        "saturday",
		"zo":        "sunday",
		"do":        "thursday",
		"di":        "tuesday",
		"wo":        "wednesday",
		"wk":        "week",
		"jr":        "year",
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
		"s":         "second",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"binnen een minuut": "0 minute ago",
		"binnen een uur":    "0 hour ago",
		"volgende maand":    "in 1 month",
		"volgende week":     "in 1 week",
		"vorige maand":      "1 month ago",
		"volgend jaar":      "in 1 year",
		"vorige week":       "1 week ago",
		"vorige jaar":       "1 year ago",
		"eergisteren":       "2 day ago",
		"deze maand":        "0 month ago",
		"vorig jaar":        "1 year ago",
		"overmorgen":        "in 2 day",
		"deze week":         "0 week ago",
		"dit jaar":          "0 year ago",
		"gisteren":          "1 day ago",
		"vandaag":           "0 day ago",
		"morgen":            "in 1 day",
		"nu":                "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) seconden geleden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) minuten geleden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) maanden geleden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) seconde geleden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) minuut geleden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) dagen geleden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) maand geleden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) weken geleden`), "$1 week ago"},
		{regexp.MustCompile(`(?i)over (\d+) seconden`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) week geleden`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) jaar geleden`), "$1 year ago"},
		{regexp.MustCompile(`(?i)over (\d+) minuten`), "in $1 minute"},
		{regexp.MustCompile(`(?i)over (\d+) maanden`), "in $1 month"},
		{regexp.MustCompile(`(?i)over (\d+) seconde`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) dag geleden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) dgn geleden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) uur geleden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) min geleden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sec geleden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)over (\d+) minuut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)over (\d+) dagen`), "in $1 day"},
		{regexp.MustCompile(`(?i)over (\d+) maand`), "in $1 month"},
		{regexp.MustCompile(`(?i)over (\d+) weken`), "in $1 week"},
		{regexp.MustCompile(`(?i)over (\d+) week`), "in $1 week"},
		{regexp.MustCompile(`(?i)over (\d+) jaar`), "in $1 year"},
		{regexp.MustCompile(`(?i)over (\d+) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)over (\d+) dgn`), "in $1 day"},
		{regexp.MustCompile(`(?i)over (\d+) uur`), "in $1 hour"},
		{regexp.MustCompile(`(?i)over (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)over (\d+) sec`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var nl_AW_Locale = merge(&nl_Locale, LocaleData{
	Name:            "nl-AW",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var nl_BE_Locale = merge(&nl_Locale, LocaleData{
	Name:            "nl-BE",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var nl_BQ_Locale = merge(&nl_Locale, LocaleData{
	Name:            "nl-BQ",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var nl_CW_Locale = merge(&nl_Locale, LocaleData{
	Name:            "nl-CW",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var nl_SR_Locale = merge(&nl_Locale, LocaleData{
	Name:            "nl-SR",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var nl_SX_Locale = merge(&nl_Locale, LocaleData{
	Name:            "nl-SX",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ seconden geleden|\d+ maanden geleden|\d+ minuten geleden|\d+ seconde geleden|\d+ minuut geleden|\d+ dagen geleden|\d+ maand geleden|\d+ weken geleden|over \d+ seconden|\d+ jaar geleden|\d+ week geleden|over \d+ maanden|over \d+ minuten|over \d+ seconde|\d+ dag geleden|\d+ dgn geleden|\d+ min geleden|\d+ sec geleden|\d+ uur geleden|over \d+ minuut|over \d+ dagen|over \d+ maand|over \d+ weken|over \d+ jaar|over \d+ week|over \d+ dag|over \d+ dgn|over \d+ min|over \d+ sec|over \d+ uur)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(binnen een minuut|binnen een uur|volgende maand|volgende week|volgend jaar|vorige maand|eergisteren|vorige jaar|vorige week|deze maand|overmorgen|vorig jaar|deze week|donderdag|september|augustus|december|dit jaar|februari|gisteren|november|seconden|woensdag|zaterdag|dinsdag|geleden|januari|maandag|maanden|minuten|oktober|seconde|vandaag|vrijdag|minuut|morgen|zondag|april|dagen|maand|maart|weken|jaar|juli|juni|week|apr|aug|dag|dec|feb|gmt|jan|jul|jun|mei|min|mnd|mrt|nov|okt|sec|sep|utc|uur|\+|\.|\[|\]|\||am|di|do|in|jr|ma|nu|om|pm|vr|wk|wo|za|zo| |'|,|-|/|:|;|@|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
