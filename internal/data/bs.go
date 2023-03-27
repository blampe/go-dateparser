// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bs_Locale LocaleData

func init() {
	bs_Locale = merge(nil, LocaleData{
		Name:      "bs",
		DateOrder: "DMY.",
		Charset:   []rune(`bcdefghijklnorstuvzćčš`),
		Translations: map[string][]string{
			"ponedjeljak": {"monday"},
			"prijepodne":  {"am"},
			"septembar":   {"september"},
			"cetvrtak":    {"thursday"},
			"decembar":    {"december"},
			"nedjelja":    {"sunday"},
			"novembar":    {"november"},
			"februar":     {"february"},
			"oktobar":     {"october"},
			"popodne":     {"pm"},
			"sedmica":     {"week"},
			"sekunda":     {"second"},
			"srijeda":     {"wednesday"},
			"avgust":      {"august"},
			"godina":      {"year"},
			"januar":      {"january"},
			"minuta":      {"minute"},
			"mjesec":      {"month"},
			"subota":      {"saturday"},
			"utorak":      {"tuesday"},
			"april":       {"april"},
			"petak":       {"friday"},
			"juli":        {"july"},
			"juni":        {"june"},
			"mart":        {"march"},
			"apr":         {"april"},
			"avg":         {"august"},
			"cet":         {"thursday"},
			"dan":         {"day"},
			"dec":         {"december"},
			"feb":         {"february"},
			"gmt":         {"gmt"},
			"god":         {"year"},
			"jan":         {"january"},
			"jul":         {"july"},
			"jun":         {"june"},
			"maj":         {"may"},
			"mar":         {"march"},
			"min":         {"minute"},
			"ned":         {"sunday"},
			"nov":         {"november"},
			"okt":         {"october"},
			"pet":         {"friday"},
			"pon":         {"monday"},
			"sat":         {"hour"},
			"sed":         {"week"},
			"sek":         {"second"},
			"sep":         {"september"},
			"sri":         {"wednesday"},
			"sub":         {"saturday"},
			"utc":         {"utc"},
			"uto":         {"tuesday"},
			"am":          {"am"},
			"mj":          {"month"},
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
			"g":           {"year"},
			"h":           {"hour"},
			"s":           {"second"},
			"z":           {"z"},
			"|":           {""},
		},
		RelativeType: map[string]string{
			"sljedece sedmice": "in 1 week",
			"sljedece godine":  "in 1 year",
			"sljedeci mjesec":  "in 1 month",
			"prosle sedmice":   "1 week ago",
			"prosle godine":    "1 year ago",
			"prosli mjesec":    "1 month ago",
			"ovaj mjesec":      "0 month ago",
			"ove sedmice":      "0 week ago",
			"ova minuta":       "0 minute ago",
			"ove godine":       "0 year ago",
			"ovaj sat":         "0 hour ago",
			"danas":            "0 day ago",
			"jucer":            "1 day ago",
			"sutra":            "in 1 day",
			"sada":             "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) mjeseci`), "$1 month ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sedmica`), "$1 week ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sedmicu`), "$1 week ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sekundi`), "$1 second ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sekundu`), "$1 second ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) godina`), "$1 year ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) godinu`), "$1 year ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) minuta`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) minutu`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) mjesec`), "$1 month ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) dana`), "$1 day ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sati`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mjeseci`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sedmica`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sedmicu`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundi`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundu`), "in $1 second"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) dan`), "$1 day ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) god`), "$1 year ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sat`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sed`), "$1 week ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) sek`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) godina`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) godinu`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minuta`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minutu`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mjesec`), "in $1 month"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) mj`), "$1 month ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)prije (\d+[.,]?\d*) g`), "$1 year ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dana`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sati`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dan`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) god`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sat`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sed`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mj`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) g`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(prije \d+[.,]?\d* mjeseci|prije \d+[.,]?\d* sedmica|prije \d+[.,]?\d* sedmicu|prije \d+[.,]?\d* sekundi|prije \d+[.,]?\d* sekundu|prije \d+[.,]?\d* godina|prije \d+[.,]?\d* godinu|prije \d+[.,]?\d* minuta|prije \d+[.,]?\d* minutu|prije \d+[.,]?\d* mjesec|prije \d+[.,]?\d* dana|prije \d+[.,]?\d* sati|za \d+[.,]?\d* mjeseci|za \d+[.,]?\d* sedmica|za \d+[.,]?\d* sedmicu|za \d+[.,]?\d* sekundi|za \d+[.,]?\d* sekundu|prije \d+[.,]?\d* dan|prije \d+[.,]?\d* god|prije \d+[.,]?\d* min|prije \d+[.,]?\d* sat|prije \d+[.,]?\d* sed|prije \d+[.,]?\d* sek|za \d+[.,]?\d* godina|za \d+[.,]?\d* godinu|za \d+[.,]?\d* minuta|za \d+[.,]?\d* minutu|za \d+[.,]?\d* mjesec|prije \d+[.,]?\d* mj|prije \d+[.,]?\d* d|prije \d+[.,]?\d* g|za \d+[.,]?\d* dana|za \d+[.,]?\d* sati|za \d+[.,]?\d* dan|za \d+[.,]?\d* god|za \d+[.,]?\d* min|za \d+[.,]?\d* sat|za \d+[.,]?\d* sed|za \d+[.,]?\d* sek|za \d+[.,]?\d* mj|za \d+[.,]?\d* d|za \d+[.,]?\d* g)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(prije \d+[.,]?\d* mjeseci|prije \d+[.,]?\d* sedmica|prije \d+[.,]?\d* sedmicu|prije \d+[.,]?\d* sekundi|prije \d+[.,]?\d* sekundu|prije \d+[.,]?\d* godina|prije \d+[.,]?\d* godinu|prije \d+[.,]?\d* minuta|prije \d+[.,]?\d* minutu|prije \d+[.,]?\d* mjesec|prije \d+[.,]?\d* dana|prije \d+[.,]?\d* sati|za \d+[.,]?\d* mjeseci|za \d+[.,]?\d* sedmica|za \d+[.,]?\d* sedmicu|za \d+[.,]?\d* sekundi|za \d+[.,]?\d* sekundu|prije \d+[.,]?\d* dan|prije \d+[.,]?\d* god|prije \d+[.,]?\d* min|prije \d+[.,]?\d* sat|prije \d+[.,]?\d* sed|prije \d+[.,]?\d* sek|za \d+[.,]?\d* godina|za \d+[.,]?\d* godinu|za \d+[.,]?\d* minuta|za \d+[.,]?\d* minutu|za \d+[.,]?\d* mjesec|prije \d+[.,]?\d* mj|prije \d+[.,]?\d* d|prije \d+[.,]?\d* g|za \d+[.,]?\d* dana|za \d+[.,]?\d* sati|za \d+[.,]?\d* dan|za \d+[.,]?\d* god|za \d+[.,]?\d* min|za \d+[.,]?\d* sat|za \d+[.,]?\d* sed|za \d+[.,]?\d* sek|za \d+[.,]?\d* mj|za \d+[.,]?\d* d|za \d+[.,]?\d* g)$`),
		KnownWords:      []string{"sljedece sedmice", "sljedece godine", "sljedeci mjesec", "prosle sedmice", "prosle godine", "prosli mjesec", "ovaj mjesec", "ove sedmice", "ponedjeljak", "ova minuta", "ove godine", "prijepodne", "septembar", "cetvrtak", "decembar", "nedjelja", "novembar", "ovaj sat", "februar", "oktobar", "popodne", "sedmica", "sekunda", "srijeda", "avgust", "godina", "januar", "minuta", "mjesec", "subota", "utorak", "april", "danas", "jucer", "petak", "sutra", "juli", "juni", "mart", "sada", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sed", "sek", "sep", "sri", "sub", "utc", "uto", "am", "mj", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "g", "h", "s", "z", "|"},
	})
}
