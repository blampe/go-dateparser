// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ko_Locale = merge(nil, LocaleData{
	Name:      "ko",
	DateOrder: "YMD.",
	Charset:   []rune(`cgtuz간금난내년늘다달목번분수시어오올요월음이일작재전제주지초토해현화후`),
	Translations: map[string]string{
		"12월": "december",
		"금요일": "friday",
		"gmt": "gmt",
		"월요일": "monday",
		"11월": "november",
		"10월": "october",
		"토요일": "saturday",
		"일요일": "sunday",
		"목요일": "thursday",
		"화요일": "tuesday",
		"utc": "utc",
		"수요일": "wednesday",
		"am":  "am",
		"오전":  "am",
		"4월":  "april",
		"8월":  "august",
		"2월":  "february",
		"1월":  "january",
		"7월":  "july",
		"6월":  "june",
		"3월":  "march",
		"5월":  "may",
		"pm":  "pm",
		"오후":  "pm",
		"9월":  "september",
		"'":   "",
		",":   "",
		";":   "",
		"@":   "",
		"[":   "",
		"]":   "",
		"|":   "",
		" ":   " ",
		"+":   "+",
		"-":   "-",
		".":   ".",
		"/":   "/",
		":":   ":",
		"일":   "day",
		"금":   "friday",
		"시":   "hour",
		"분":   "minute",
		"월":   "month",
		"토":   "saturday",
		"초":   "second",
		"목":   "thursday",
		"화":   "tuesday",
		"수":   "wednesday",
		"주":   "week",
		"년":   "year",
		"z":   "z",
	},
	RelativeType: map[string]string{
		"현재 시간": "0 hour ago",
		"현재 분":  "0 minute ago",
		"이번 달":  "0 month ago",
		"이번 주":  "0 week ago",
		"다음 달":  "in 1 month",
		"다음 주":  "in 1 week",
		"지난달":   "1 month ago",
		"지난주":   "1 week ago",
		"오늘":    "0 day ago",
		"지금":    "0 second ago",
		"올해":    "0 year ago",
		"어제":    "1 day ago",
		"작년":    "1 year ago",
		"내일":    "in 1 day",
		"내년":    "in 1 year",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+)시간 전`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+)개월 전`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+)시간 후`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+)개월 후`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+)일 전`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+)분 전`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+)초 전`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+)주 전`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+)년 전`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+)일 후`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+)분 후`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+)초 후`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+)주 후`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+)년 후`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+개월 전|\d+개월 후|\d+시간 전|\d+시간 후|\d+년 전|\d+년 후|\d+분 전|\d+분 후|\d+일 전|\d+일 후|\d+주 전|\d+주 후|\d+초 전|\d+초 후)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+개월 전|\d+개월 후|\d+시간 전|\d+시간 후|\d+년 전|\d+년 후|\d+분 전|\d+분 후|\d+일 전|\d+일 후|\d+주 전|\d+주 후|\d+초 전|\d+초 후)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(현재 시간|다음 달|다음 주|이번 달|이번 주|현재 분|10월|11월|12월|gmt|utc|금요일|목요일|수요일|월요일|일요일|지난달|지난주|토요일|화요일|1월|2월|3월|4월|5월|6월|7월|8월|9월|\+|\.|\[|\]|\||am|pm|내년|내일|어제|오늘|오전|오후|올해|작년|지금| |'|,|-|/|:|;|@|z|금|년|목|분|수|시|월|일|주|초|토|화)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ko_KP_Locale = merge(&ko_Locale, LocaleData{
	Name:            "ko-KP",
	DateOrder:       "YMD.",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+개월 전|\d+개월 후|\d+시간 전|\d+시간 후|\d+년 전|\d+년 후|\d+분 전|\d+분 후|\d+일 전|\d+일 후|\d+주 전|\d+주 후|\d+초 전|\d+초 후)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+개월 전|\d+개월 후|\d+시간 전|\d+시간 후|\d+년 전|\d+년 후|\d+분 전|\d+분 후|\d+일 전|\d+일 후|\d+주 전|\d+주 후|\d+초 전|\d+초 후)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(현재 시간|다음 달|다음 주|이번 달|이번 주|현재 분|10월|11월|12월|gmt|utc|금요일|목요일|수요일|월요일|일요일|지난달|지난주|토요일|화요일|1월|2월|3월|4월|5월|6월|7월|8월|9월|\+|\.|\[|\]|\||am|pm|내년|내일|어제|오늘|오전|오후|올해|작년|지금| |'|,|-|/|:|;|@|z|금|년|목|분|수|시|월|일|주|초|토|화)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
