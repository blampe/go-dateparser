// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var jgo_Locale LocaleData

func init() {
	jgo_Locale = merge(nil, LocaleData{
		Name:      "jgo",
		DateOrder: "YMD",
		Charset:   []rune(`bcdefghijklnorstuwxyzáâúŋɔɛ́̌ꞌ`),
		Translations: map[string][]string{
			"pɛsaŋ pɛnɛntuku": {"june"},
			"pɛsaŋ pɛnɛpfuꞌu": {"september"},
			"pɛsaŋ ntsɔpmɔ":   {"november"},
			"pɛsaŋ ntsɔppa":   {"december"},
			"pɛsaŋ pɛnɛfɔm":   {"august"},
			"pɛsaŋ pɛnɛkwa":   {"april"},
			"pɛsaŋ saamba":    {"july"},
			"ŋka mbɔt nji":    {"pm"},
			"nduŋmbi saŋ":     {"january"},
			"pɛsaŋ nɛgɛm":     {"october"},
			"pɛsaŋ pataa":     {"may"},
			"pɛsaŋ pɛtat":     {"march"},
			"apta mɔndi":      {"tuesday"},
			"pɛsaŋ pɛpa":      {"february"},
			"fɛlayɛdɛ":        {"friday"},
			"mba'mba'":        {"am"},
			"wɛnɛsɛdɛ":        {"wednesday"},
			"minute":          {"minute"},
			"sasidɛ":          {"saturday"},
			"second":          {"second"},
			"tɔsɛdɛ":          {"thursday"},
			"month":           {"month"},
			"mɔndi":           {"monday"},
			"sɔndi":           {"sunday"},
			"hour":            {"hour"},
			"week":            {"week"},
			"year":            {"year"},
			"day":             {"day"},
			"gmt":             {"gmt"},
			"utc":             {"utc"},
			"am":              {"am"},
			"pm":              {"pm"},
			" ":               {" "},
			"'":               {""},
			"+":               {"+"},
			",":               {""},
			"-":               {"-"},
			".":               {"."},
			"/":               {"/"},
			":":               {":"},
			";":               {""},
			"@":               {""},
			"[":               {""},
			"]":               {""},
			"z":               {"z"},
			"|":               {""},
		},
		RelativeType: map[string]string{
			"this minute": "0 minute ago",
			"last month":  "1 month ago",
			"next month":  "in 1 month",
			"this month":  "0 month ago",
			"last week":   "1 week ago",
			"last year":   "1 year ago",
			"next week":   "in 1 week",
			"next year":   "in 1 year",
			"this hour":   "0 hour ago",
			"this week":   "0 week ago",
			"this year":   "0 year ago",
			"yesterday":   "1 day ago",
			"tomorrow":    "in 1 day",
			"lɔꞌɔ":        "0 day ago",
			"now":         "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)ɛ gɛ mɔ (\d+[.,]?\d*) ŋgap-mbi`), "$1 week ago"},
			{regexp.MustCompile(`(?i)ɛ gɛ mɔ minut (\d+[.,]?\d*)`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)ɛ gɛ mɔ pɛsaŋ (\d+[.,]?\d*)`), "$1 month ago"},
			{regexp.MustCompile(`(?i)nuu ŋgap-mbi (\d+[.,]?\d*)`), "in $1 week"},
			{regexp.MustCompile(`(?i)ɛ gɛ mɔ (\d+[.,]?\d*) hawa`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ɛ gɛ mɔ lɛꞌ (\d+[.,]?\d*)`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ɛgɛ mɔ ŋguꞌ (\d+[.,]?\d*)`), "$1 year ago"},
			{regexp.MustCompile(`(?i)nuu (\d+[.,]?\d*) minut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)nuu hawa (\d+[.,]?\d*)`), "in $1 hour"},
			{regexp.MustCompile(`(?i)nuu ŋguꞌ (\d+[.,]?\d*)`), "in $1 year"},
			{regexp.MustCompile(`(?i)nuu (\d+[.,]?\d*) saŋ`), "in $1 month"},
			{regexp.MustCompile(`(?i)nuu lɛꞌ (\d+[.,]?\d*)`), "in $1 day"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(ɛ gɛ mɔ \d+[.,]?\d* ŋgap-mbi|ɛ gɛ mɔ minut \d+[.,]?\d*|ɛ gɛ mɔ pɛsaŋ \d+[.,]?\d*|nuu ŋgap-mbi \d+[.,]?\d*|ɛ gɛ mɔ \d+[.,]?\d* hawa|ɛ gɛ mɔ lɛꞌ \d+[.,]?\d*|ɛgɛ mɔ ŋguꞌ \d+[.,]?\d*|nuu \d+[.,]?\d* minut|nuu hawa \d+[.,]?\d*|nuu ŋguꞌ \d+[.,]?\d*|nuu \d+[.,]?\d* saŋ|nuu lɛꞌ \d+[.,]?\d*)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(ɛ gɛ mɔ \d+[.,]?\d* ŋgap-mbi|ɛ gɛ mɔ minut \d+[.,]?\d*|ɛ gɛ mɔ pɛsaŋ \d+[.,]?\d*|nuu ŋgap-mbi \d+[.,]?\d*|ɛ gɛ mɔ \d+[.,]?\d* hawa|ɛ gɛ mɔ lɛꞌ \d+[.,]?\d*|ɛgɛ mɔ ŋguꞌ \d+[.,]?\d*|nuu \d+[.,]?\d* minut|nuu hawa \d+[.,]?\d*|nuu ŋguꞌ \d+[.,]?\d*|nuu \d+[.,]?\d* saŋ|nuu lɛꞌ \d+[.,]?\d*)$`),
		KnownWords:      []string{"pɛsaŋ pɛnɛntuku", "pɛsaŋ pɛnɛpfuꞌu", "pɛsaŋ ntsɔpmɔ", "pɛsaŋ ntsɔppa", "pɛsaŋ pɛnɛfɔm", "pɛsaŋ pɛnɛkwa", "pɛsaŋ saamba", "ŋka mbɔt nji", "nduŋmbi saŋ", "pɛsaŋ nɛgɛm", "pɛsaŋ pataa", "pɛsaŋ pɛtat", "this minute", "apta mɔndi", "last month", "next month", "pɛsaŋ pɛpa", "this month", "last week", "last year", "next week", "next year", "this hour", "this week", "this year", "yesterday", "fɛlayɛdɛ", "mba'mba'", "tomorrow", "wɛnɛsɛdɛ", "minute", "sasidɛ", "second", "tɔsɛdɛ", "month", "mɔndi", "sɔndi", "hour", "lɔꞌɔ", "week", "year", "day", "gmt", "now", "utc", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
