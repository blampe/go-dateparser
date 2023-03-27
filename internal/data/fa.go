// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fa_Locale LocaleData

func init() {
	fa_Locale = merge(nil, LocaleData{
		Name:                  "fa",
		DateOrder:             "YMD",
		Charset:               []rune(`cgtuzآئابتثجدذرزسشظعفقلمنهوئپچژکگی‌`),
		SentenceSplitterGroup: 6,
		Translations: map[string][]string{
			"قبل‌ازظهر": {"am"},
			"چهار شنبه": {"wednesday"},
			"بعدازظهر":  {"pm"},
			"روز شنبه":  {"saturday"},
			"پنج شنبه":  {"thursday"},
			"چهارشنبه":  {"wednesday"},
			"دو شنبه":   {"saturday"},
			"سه‌شنبه":   {"tuesday"},
			"سپتامبر":   {"september"},
			"پنجشنبه":   {"thursday"},
			"اکتوبر":    {"october"},
			"دسامبر":    {"december"},
			"دوشنبه":    {"monday"},
			"سهشنبه":    {"tuesday"},
			"سپتمبر":    {"september"},
			"فبروری":    {"february"},
			"نوامبر":    {"november"},
			"ژانویه":    {"january"},
			"یکشنبه":    {"sunday"},
			"اوریل":     {"april"},
			"اپریل":     {"april"},
			"اکتبر":     {"october"},
			"ثانیه":     {"second"},
			"جنوری":     {"january"},
			"جولای":     {"july"},
			"دسمبر":     {"december"},
			"دقیقه":     {"minute"},
			"فوریه":     {"february"},
			"نومبر":     {"november"},
			"ژويیه":     {"july"},
			"اگست":      {"august"},
			"جمعه":      {"friday"},
			"ساعت":      {"hour"},
			"شنبه":      {"saturday"},
			"مارس":      {"march"},
			"مارچ":      {"march"},
			"هفته":      {"week"},
			"ژوين":      {"june"},
			"gmt":       {"gmt"},
			"utc":       {"utc"},
			"اوت":       {"august"},
			"جون":       {"june"},
			"دوم":       {"second"},
			"روز":       {"day"},
			"سال":       {"year"},
			"ماه":       {"month"},
			"پیش":       {"ago"},
			"am":        {"am"},
			"pm":        {"pm"},
			"بظ":        {"pm"},
			"در":        {"in"},
			"قظ":        {"am"},
			"مه":        {"may"},
			"می":        {"may"},
			" ":         {" "},
			"'":         {""},
			"+":         {"+"},
			",":         {""},
			"-":         {"-"},
			".":         {"."},
			"/":         {"/"},
			":":         {":"},
			";":         {""},
			"@":         {""},
			"[":         {""},
			"]":         {""},
			"z":         {"z"},
			"|":         {""},
			"د":         {"saturday"},
		},
		RelativeType: map[string]string{
			"هفته اینده": "in 1 week",
			"هفته گذشته": "1 week ago",
			"همین دقیقه": "0 minute ago",
			"سال اینده":  "in 1 year",
			"سال گذشته":  "1 year ago",
			"ماه اینده":  "in 1 month",
			"ماه گذشته":  "1 month ago",
			"همین ساعت":  "0 hour ago",
			"این هفته":   "0 week ago",
			"این ماه":    "0 month ago",
			"ماه پیش":    "1 month ago",
			"امروز":      "0 day ago",
			"امسال":      "0 year ago",
			"اکنون":      "0 second ago",
			"دیروز":      "1 day ago",
			"فردا":       "in 1 day",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ثانیه بعد`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ثانیه پیش`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دقیقه بعد`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دقیقه پیش`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ساعت بعد`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ساعت پیش`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) هفته بعد`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) هفته پیش`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) روز بعد`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) روز پیش`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سال بعد`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سال پیش`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ماه بعد`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ماه پیش`), "$1 month ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* ثانیه بعد|\d+[.,]?\d* ثانیه پیش|\d+[.,]?\d* دقیقه بعد|\d+[.,]?\d* دقیقه پیش|\d+[.,]?\d* ساعت بعد|\d+[.,]?\d* ساعت پیش|\d+[.,]?\d* هفته بعد|\d+[.,]?\d* هفته پیش|\d+[.,]?\d* روز بعد|\d+[.,]?\d* روز پیش|\d+[.,]?\d* سال بعد|\d+[.,]?\d* سال پیش|\d+[.,]?\d* ماه بعد|\d+[.,]?\d* ماه پیش)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* ثانیه بعد|\d+[.,]?\d* ثانیه پیش|\d+[.,]?\d* دقیقه بعد|\d+[.,]?\d* دقیقه پیش|\d+[.,]?\d* ساعت بعد|\d+[.,]?\d* ساعت پیش|\d+[.,]?\d* هفته بعد|\d+[.,]?\d* هفته پیش|\d+[.,]?\d* روز بعد|\d+[.,]?\d* روز پیش|\d+[.,]?\d* سال بعد|\d+[.,]?\d* سال پیش|\d+[.,]?\d* ماه بعد|\d+[.,]?\d* ماه پیش)$`),
		KnownWords:      []string{"هفته اینده", "هفته گذشته", "همین دقیقه", "سال اینده", "سال گذشته", "قبل‌ازظهر", "ماه اینده", "ماه گذشته", "همین ساعت", "چهار شنبه", "این هفته", "بعدازظهر", "روز شنبه", "پنج شنبه", "چهارشنبه", "این ماه", "دو شنبه", "سه‌شنبه", "سپتامبر", "ماه پیش", "پنجشنبه", "اکتوبر", "دسامبر", "دوشنبه", "سهشنبه", "سپتمبر", "فبروری", "نوامبر", "ژانویه", "یکشنبه", "امروز", "امسال", "اوریل", "اپریل", "اکتبر", "اکنون", "ثانیه", "جنوری", "جولای", "دسمبر", "دقیقه", "دیروز", "فوریه", "نومبر", "ژويیه", "اگست", "جمعه", "ساعت", "شنبه", "فردا", "مارس", "مارچ", "هفته", "ژوين", "gmt", "utc", "اوت", "جون", "دوم", "روز", "سال", "ماه", "پیش", "am", "pm", "بظ", "در", "قظ", "مه", "می", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "د"},
	})
}

var fa_AF_Locale LocaleData

func init() {
	fa_AF_Locale = merge(&fa_Locale, LocaleData{
		Name:      "fa-AF",
		DateOrder: "YMD",
		Translations: map[string][]string{
			"جنو": {"january"},
			"جول": {"july"},
			"دسم": {"december"},
		},
		KnownWords: []string{"هفته اینده", "هفته گذشته", "همین دقیقه", "سال اینده", "سال گذشته", "قبل‌ازظهر", "ماه اینده", "ماه گذشته", "همین ساعت", "چهار شنبه", "این هفته", "بعدازظهر", "روز شنبه", "پنج شنبه", "چهارشنبه", "این ماه", "دو شنبه", "سه‌شنبه", "سپتامبر", "ماه پیش", "پنجشنبه", "اکتوبر", "دسامبر", "دوشنبه", "سهشنبه", "سپتمبر", "فبروری", "نوامبر", "ژانویه", "یکشنبه", "امروز", "امسال", "اوریل", "اپریل", "اکتبر", "اکنون", "ثانیه", "جنوری", "جولای", "دسمبر", "دقیقه", "دیروز", "فوریه", "نومبر", "ژويیه", "اگست", "جمعه", "ساعت", "شنبه", "فردا", "مارس", "مارچ", "هفته", "ژوين", "gmt", "utc", "اوت", "جنو", "جول", "جون", "دسم", "دوم", "روز", "سال", "ماه", "پیش", "am", "pm", "بظ", "در", "قظ", "مه", "می", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "د"},
	})
}
