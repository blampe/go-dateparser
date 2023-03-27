// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lkt_Locale LocaleData

func init() {
	lkt_Locale = merge(nil, LocaleData{
		Name:      "lkt",
		DateOrder: "YMD",
		Charset:   []rune(`-ceghiklnorstuwyzáéíóúčŋšžǧȟ`),
		Translations: map[string][]string{
			"wipazukha-waste wi": {"june"},
			"chaŋwape-kasna wi":  {"october"},
			"istawichayazaŋ wi":  {"march"},
			"owaphe oh'aŋkho":    {"minute"},
			"thiyoheyuŋka wi":    {"february"},
			"chaŋphasapa wi":     {"july"},
			"chaŋwapetho wi":     {"may"},
			"thahekapsuŋ wi":     {"december"},
			"chaŋwapegi wi":      {"september"},
			"owaŋgyuzazapi":      {"saturday"},
			"wiothehika wi":      {"january"},
			"aŋpetuwakhaŋ":       {"sunday"},
			"aŋpetuzaptaŋ":       {"friday"},
			"aŋpetunuŋpa":        {"tuesday"},
			"aŋpetuwaŋzi":        {"monday"},
			"aŋpetuyamni":        {"wednesday"},
			"phezitho wi":        {"april"},
			"waniyetu wi":        {"november"},
			"wasuthuŋ wi":        {"august"},
			"aŋpetutopa":         {"thursday"},
			"aŋpetu":             {"day"},
			"omakha":             {"year"},
			"owaphe":             {"hour"},
			"okpi":               {"second"},
			"gmt":                {"gmt"},
			"oko":                {"week"},
			"utc":                {"utc"},
			"am":                 {"am"},
			"pm":                 {"pm"},
			"wi":                 {"month"},
			" ":                  {" "},
			"'":                  {""},
			"+":                  {"+"},
			",":                  {""},
			"-":                  {"-"},
			".":                  {"."},
			"/":                  {"/"},
			":":                  {":"},
			";":                  {""},
			"@":                  {""},
			"[":                  {""},
			"]":                  {""},
			"z":                  {"z"},
			"|":                  {""},
		},
		RelativeType: map[string]string{
			"thokata omakha kiŋhaŋ": "in 1 year",
			"thokata oko kiŋhaŋ":    "in 1 week",
			"omakha k'uŋ hehaŋ":     "1 year ago",
			"thokata wi kiŋhaŋ":     "in 1 month",
			"hiŋhaŋni kiŋhaŋ":       "in 1 day",
			"oko k'uŋ hehaŋ":        "1 week ago",
			"le aŋpetu kiŋ":         "0 day ago",
			"le omakha kiŋ":         "0 year ago",
			"wi k'uŋ hehaŋ":         "1 month ago",
			"this minute":           "0 minute ago",
			"le oko kiŋ":            "0 week ago",
			"le wi kiŋ":             "0 month ago",
			"this hour":             "0 hour ago",
			"htalehaŋ":              "1 day ago",
			"now":                   "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)hekta oh'aŋkho (\d+[.,]?\d*) k'uŋ hehaŋ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)hekta wiyawapi (\d+[.,]?\d*) k'uŋ hehaŋ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)letaŋhaŋ oh'aŋkho (\d+[.,]?\d*) kiŋhaŋ`), "in $1 minute"},
			{regexp.MustCompile(`(?i)letaŋhaŋ wiyawapi (\d+[.,]?\d*) kiŋhaŋ`), "in $1 month"},
			{regexp.MustCompile(`(?i)hekta omakha (\d+[.,]?\d*) k'uŋ hehaŋ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)hekta owaphe (\d+[.,]?\d*) k'uŋ hehaŋ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)letaŋhaŋ omakha (\d+[.,]?\d*) kiŋhaŋ`), "in $1 year"},
			{regexp.MustCompile(`(?i)letaŋhaŋ owaphe (\d+[.,]?\d*) kiŋhaŋ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)hekta (\d+[.,]?\d*)-chaŋ k'uŋ hehaŋ`), "$1 day ago"},
			{regexp.MustCompile(`(?i)hekta okpi (\d+[.,]?\d*) k'uŋ hehaŋ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)hekta oko (\d+[.,]?\d*) k'uŋ hehaŋ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)letaŋhaŋ (\d+[.,]?\d*)-chaŋ kiŋhaŋ`), "in $1 day"},
			{regexp.MustCompile(`(?i)letaŋhaŋ okpi (\d+[.,]?\d*) kiŋhaŋ`), "in $1 second"},
			{regexp.MustCompile(`(?i)letaŋhaŋ oko (\d+[.,]?\d*) kiŋhaŋ`), "in $1 week"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(hekta oh'aŋkho \d+[.,]?\d* k'uŋ hehaŋ|hekta wiyawapi \d+[.,]?\d* k'uŋ hehaŋ|letaŋhaŋ oh'aŋkho \d+[.,]?\d* kiŋhaŋ|letaŋhaŋ wiyawapi \d+[.,]?\d* kiŋhaŋ|hekta omakha \d+[.,]?\d* k'uŋ hehaŋ|hekta owaphe \d+[.,]?\d* k'uŋ hehaŋ|letaŋhaŋ omakha \d+[.,]?\d* kiŋhaŋ|letaŋhaŋ owaphe \d+[.,]?\d* kiŋhaŋ|hekta \d+[.,]?\d*-chaŋ k'uŋ hehaŋ|hekta okpi \d+[.,]?\d* k'uŋ hehaŋ|hekta oko \d+[.,]?\d* k'uŋ hehaŋ|letaŋhaŋ \d+[.,]?\d*-chaŋ kiŋhaŋ|letaŋhaŋ okpi \d+[.,]?\d* kiŋhaŋ|letaŋhaŋ oko \d+[.,]?\d* kiŋhaŋ)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(hekta oh'aŋkho \d+[.,]?\d* k'uŋ hehaŋ|hekta wiyawapi \d+[.,]?\d* k'uŋ hehaŋ|letaŋhaŋ oh'aŋkho \d+[.,]?\d* kiŋhaŋ|letaŋhaŋ wiyawapi \d+[.,]?\d* kiŋhaŋ|hekta omakha \d+[.,]?\d* k'uŋ hehaŋ|hekta owaphe \d+[.,]?\d* k'uŋ hehaŋ|letaŋhaŋ omakha \d+[.,]?\d* kiŋhaŋ|letaŋhaŋ owaphe \d+[.,]?\d* kiŋhaŋ|hekta \d+[.,]?\d*-chaŋ k'uŋ hehaŋ|hekta okpi \d+[.,]?\d* k'uŋ hehaŋ|hekta oko \d+[.,]?\d* k'uŋ hehaŋ|letaŋhaŋ \d+[.,]?\d*-chaŋ kiŋhaŋ|letaŋhaŋ okpi \d+[.,]?\d* kiŋhaŋ|letaŋhaŋ oko \d+[.,]?\d* kiŋhaŋ)$`),
		KnownWords:      []string{"thokata omakha kiŋhaŋ", "thokata oko kiŋhaŋ", "wipazukha-waste wi", "chaŋwape-kasna wi", "istawichayazaŋ wi", "omakha k'uŋ hehaŋ", "thokata wi kiŋhaŋ", "hiŋhaŋni kiŋhaŋ", "owaphe oh'aŋkho", "thiyoheyuŋka wi", "chaŋphasapa wi", "chaŋwapetho wi", "oko k'uŋ hehaŋ", "thahekapsuŋ wi", "chaŋwapegi wi", "le aŋpetu kiŋ", "le omakha kiŋ", "owaŋgyuzazapi", "wi k'uŋ hehaŋ", "wiothehika wi", "aŋpetuwakhaŋ", "aŋpetuzaptaŋ", "aŋpetunuŋpa", "aŋpetuwaŋzi", "aŋpetuyamni", "phezitho wi", "this minute", "waniyetu wi", "wasuthuŋ wi", "aŋpetutopa", "le oko kiŋ", "le wi kiŋ", "this hour", "htalehaŋ", "aŋpetu", "omakha", "owaphe", "okpi", "gmt", "now", "oko", "utc", "am", "pm", "wi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
