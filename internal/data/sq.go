// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sq_Locale LocaleData

func init() {
	sq_Locale = merge(nil, LocaleData{
		Name:      "sq",
		DateOrder: "DMY",
		Charset:   []rune(`cdeghijklnorstuvzë`),
		Translations: map[string][]string{
			"e paradites": {"am"},
			"e pasdites":  {"pm"},
			"e merkure":   {"wednesday"},
			"e premte":    {"friday"},
			"e shtune":    {"saturday"},
			"paradite":    {"am"},
			"dhjetor":     {"december"},
			"e enjte":     {"thursday"},
			"e marte":     {"tuesday"},
			"pasdite":     {"pm"},
			"qershor":     {"june"},
			"sekonde":     {"second"},
			"shtator":     {"september"},
			"e diel":      {"sunday"},
			"e hene":      {"monday"},
			"korrik":      {"july"},
			"minute":      {"minute"},
			"nentor":      {"november"},
			"shkurt":      {"february"},
			"gusht":       {"august"},
			"janar":       {"january"},
			"prill":       {"april"},
			"tetor":       {"october"},
			"dite":        {"day"},
			"jave":        {"week"},
			"mars":        {"march"},
			"muaj":        {"month"},
			"dhj":         {"december"},
			"die":         {"sunday"},
			"enj":         {"thursday"},
			"gmt":         {"gmt"},
			"gsh":         {"august"},
			"hen":         {"monday"},
			"jan":         {"january"},
			"kor":         {"july"},
			"maj":         {"may"},
			"mar":         {"march", "tuesday"},
			"mer":         {"wednesday"},
			"min":         {"minute"},
			"nen":         {"november"},
			"ore":         {"hour"},
			"pre":         {"friday"},
			"pri":         {"april"},
			"qer":         {"june"},
			"sek":         {"second"},
			"shk":         {"february"},
			"sht":         {"saturday", "september"},
			"tet":         {"october"},
			"utc":         {"utc"},
			"vit":         {"year"},
			"am":          {"am"},
			"pm":          {"pm"},
			" ":           {" "},
			"'":           {""},
			"+":           {"+"},
			",":           {""},
			"-":           {"-"},
			".":           {"."},
			"/":           {"/"},
			":":           {":"},
			";":           {""},
			"@":           {""},
			"[":           {""},
			"]":           {""},
			"z":           {"z"},
			"|":           {""},
		},
		RelativeType: map[string]string{
			"muajin e ardhshem": "in 1 month",
			"javen e ardhshme":  "in 1 week",
			"vitin e ardhshem":  "in 1 year",
			"muajin e kaluar":   "1 month ago",
			"javen e kaluar":    "1 week ago",
			"vitin e kaluar":    "1 year ago",
			"kete minute":       "0 minute ago",
			"kete jave":         "0 week ago",
			"kete muaj":         "0 month ago",
			"kete ore":          "0 hour ago",
			"kete vit":          "0 year ago",
			"neser":             "in 1 day",
			"tani":              "0 second ago",
			"dje":               "1 day ago",
			"sot":               "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekonda me pare`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekonde me pare`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuta me pare`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minute me pare`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) sekondash`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dite me pare`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) jave me pare`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) muaj me pare`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vjet me pare`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) minutash`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min me pare`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ore me pare`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sek me pare`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vit me pare`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) sekonde`), "in $1 second"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) vjetesh`), "in $1 year"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) ditesh`), "in $1 day"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) javesh`), "in $1 week"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) minute`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) muajsh`), "in $1 month"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) muaji`), "in $1 month"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) oresh`), "in $1 hour"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) dite`), "in $1 day"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) jave`), "in $1 week"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) viti`), "in $1 year"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) ore`), "in $1 hour"},
			{regexp.MustCompile(`(?i)pas (\d+[.,]?\d*) sek`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* sekonda me pare|\d+[.,]?\d* sekonde me pare|\d+[.,]?\d* minuta me pare|\d+[.,]?\d* minute me pare|pas \d+[.,]?\d* sekondash|\d+[.,]?\d* dite me pare|\d+[.,]?\d* jave me pare|\d+[.,]?\d* muaj me pare|\d+[.,]?\d* vjet me pare|pas \d+[.,]?\d* minutash|\d+[.,]?\d* min me pare|\d+[.,]?\d* ore me pare|\d+[.,]?\d* sek me pare|\d+[.,]?\d* vit me pare|pas \d+[.,]?\d* sekonde|pas \d+[.,]?\d* vjetesh|pas \d+[.,]?\d* ditesh|pas \d+[.,]?\d* javesh|pas \d+[.,]?\d* minute|pas \d+[.,]?\d* muajsh|pas \d+[.,]?\d* muaji|pas \d+[.,]?\d* oresh|pas \d+[.,]?\d* dite|pas \d+[.,]?\d* jave|pas \d+[.,]?\d* viti|pas \d+[.,]?\d* min|pas \d+[.,]?\d* ore|pas \d+[.,]?\d* sek)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* sekonda me pare|\d+[.,]?\d* sekonde me pare|\d+[.,]?\d* minuta me pare|\d+[.,]?\d* minute me pare|pas \d+[.,]?\d* sekondash|\d+[.,]?\d* dite me pare|\d+[.,]?\d* jave me pare|\d+[.,]?\d* muaj me pare|\d+[.,]?\d* vjet me pare|pas \d+[.,]?\d* minutash|\d+[.,]?\d* min me pare|\d+[.,]?\d* ore me pare|\d+[.,]?\d* sek me pare|\d+[.,]?\d* vit me pare|pas \d+[.,]?\d* sekonde|pas \d+[.,]?\d* vjetesh|pas \d+[.,]?\d* ditesh|pas \d+[.,]?\d* javesh|pas \d+[.,]?\d* minute|pas \d+[.,]?\d* muajsh|pas \d+[.,]?\d* muaji|pas \d+[.,]?\d* oresh|pas \d+[.,]?\d* dite|pas \d+[.,]?\d* jave|pas \d+[.,]?\d* viti|pas \d+[.,]?\d* min|pas \d+[.,]?\d* ore|pas \d+[.,]?\d* sek)$`),
		KnownWords:      []string{"muajin e ardhshem", "javen e ardhshme", "vitin e ardhshem", "muajin e kaluar", "javen e kaluar", "vitin e kaluar", "e paradites", "kete minute", "e pasdites", "e merkure", "kete jave", "kete muaj", "e premte", "e shtune", "kete ore", "kete vit", "paradite", "dhjetor", "e enjte", "e marte", "pasdite", "qershor", "sekonde", "shtator", "e diel", "e hene", "korrik", "minute", "nentor", "shkurt", "gusht", "janar", "neser", "prill", "tetor", "dite", "jave", "mars", "muaj", "tani", "dhj", "die", "dje", "enj", "gmt", "gsh", "hen", "jan", "kor", "maj", "mar", "mer", "min", "nen", "ore", "pre", "pri", "qer", "sek", "shk", "sht", "sot", "tet", "utc", "vit", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}

var sq_MK_Locale LocaleData

func init() {
	sq_MK_Locale = merge(&sq_Locale, LocaleData{
		Name:      "sq-MK",
		DateOrder: "DMY",
	})
}

var sq_XK_Locale LocaleData

func init() {
	sq_XK_Locale = merge(&sq_Locale, LocaleData{
		Name:      "sq-XK",
		DateOrder: "DMY",
	})
}
