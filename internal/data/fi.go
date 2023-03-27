// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fi_Locale LocaleData

func init() {
	fi_Locale = merge(nil, LocaleData{
		Name:                  "fi",
		DateOrder:             "DMY",
		Charset:               []rune(`cdeghijklnorstuvyzä`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d*) (sekunnin|sekuntin|minuutin|tunnin|paivan|viikon|kuukauden|vuoden) (paasta|kuluttua)(\z|[^\pL\pM\d]|_)`), "${1}${4} ${2} ${3}${5}"},
		},
		Translations: map[string][]string{
			"keskiviikkona": {"wednesday"},
			"maaliskuussa":  {"march"},
			"marraskuussa":  {"november"},
			"heinakuussa":   {"july"},
			"helmikuussa":   {"february"},
			"huhtikuussa":   {"april"},
			"joulukuussa":   {"december"},
			"keskiviikko":   {"wednesday"},
			"maaliskuuta":   {"march"},
			"maanantaina":   {"monday"},
			"marraskuuta":   {"november"},
			"perjantaina":   {"friday"},
			"sunnuntaina":   {"sunday"},
			"tammikuussa":   {"january"},
			"toukokuussa":   {"may"},
			"heinakuuta":    {"july"},
			"helmikuuta":    {"february"},
			"huhtikuuta":    {"april"},
			"joulukuuta":    {"december"},
			"kesakuussa":    {"june"},
			"lauantaina":    {"saturday"},
			"lokakuussa":    {"october"},
			"syyskuussa":    {"september"},
			"tammikuuta":    {"january"},
			"toukokuuta":    {"may"},
			"elokuussa":     {"august"},
			"kesakuuta":     {"june"},
			"kuukauden":     {"month"},
			"kuukautta":     {"month"},
			"lokakuuta":     {"october"},
			"maaliskuu":     {"march"},
			"maanantai":     {"monday"},
			"marraskuu":     {"november"},
			"minuuttia":     {"minute"},
			"perjantai":     {"friday"},
			"sekunttia":     {"second"},
			"sunnuntai":     {"sunday"},
			"syyskuuta":     {"september"},
			"tiistaina":     {"tuesday"},
			"torstaina":     {"thursday"},
			"elokuuta":      {"august"},
			"heinakuu":      {"july"},
			"helmikuu":      {"february"},
			"huhtikuu":      {"april"},
			"joulukuu":      {"december"},
			"kuluttua":      {"in"},
			"kuukausi":      {"month"},
			"lauantai":      {"saturday"},
			"minuutin":      {"minute"},
			"minuutti":      {"minute"},
			"sekunnin":      {"second"},
			"sekuntia":      {"second"},
			"sekuntin":      {"second"},
			"sekuntti":      {"second"},
			"tammikuu":      {"january"},
			"toukokuu":      {"may"},
			"kesakuu":       {"june"},
			"lokakuu":       {"october"},
			"maalisk":       {"march"},
			"marrask":       {"november"},
			"sekunti":       {"second"},
			"syyskuu":       {"september"},
			"tiistai":       {"tuesday"},
			"torstai":       {"thursday"},
			"viikkoa":       {"week"},
			"elokuu":        {"august"},
			"heinak":        {"july"},
			"helmik":        {"february"},
			"huhtik":        {"april"},
			"jouluk":        {"december"},
			"maalis":        {"march"},
			"marras":        {"november"},
			"paasta":        {"in"},
			"paivaa":        {"day"},
			"paivan":        {"day"},
			"sitten":        {"ago"},
			"tammik":        {"january"},
			"toukok":        {"may"},
			"tunnin":        {"hour"},
			"tuntia":        {"hour"},
			"viikko":        {"week"},
			"viikon":        {"week"},
			"vuoden":        {"year"},
			"vuonna":        {"year"},
			"vuotta":        {"year"},
			"heina":         {"july"},
			"helmi":         {"february"},
			"huhti":         {"april"},
			"joulu":         {"december"},
			"kesak":         {"june"},
			"lokak":         {"october"},
			"paiva":         {"day"},
			"syysk":         {"september"},
			"tammi":         {"january"},
			"touko":         {"may"},
			"tunti":         {"hour"},
			"vuosi":         {"year"},
			"elok":          {"august"},
			"kesa":          {"june"},
			"loka":          {"october"},
			"pvaa":          {"day"},
			"syys":          {"september"},
			"elo":           {"august"},
			"gmt":           {"gmt"},
			"min":           {"minute"},
			"pva":           {"day"},
			"utc":           {"utc"},
			"vko":           {"week"},
			":n":            {""},
			"am":            {"am"},
			"ap":            {"am"},
			"ip":            {"pm"},
			"ke":            {"wednesday"},
			"kk":            {"month"},
			"la":            {"saturday"},
			"ma":            {"monday"},
			"pe":            {"friday"},
			"pm":            {"pm"},
			"pv":            {"day"},
			"su":            {"sunday"},
			"ti":            {"tuesday"},
			"to":            {"thursday"},
			"vk":            {"week"},
			"vv":            {"year"},
			" ":             {" "},
			"'":             {""},
			"+":             {"+"},
			",":             {""},
			"-":             {"-"},
			".":             {"."},
			"/":             {"/"},
			":":             {":"},
			";":             {""},
			"@":             {""},
			"[":             {""},
			"]":             {""},
			"p":             {"day"},
			"s":             {"second"},
			"t":             {"hour"},
			"v":             {"year"},
			"z":             {"z"},
			"|":             {""},
		},
		RelativeType: map[string]string{
			"taman minuutin aikana": "0 minute ago",
			"taman tunnin aikana":   "0 hour ago",
			"minuutin sisalla":      "0 minute ago",
			"toissa viikolla":       "2 week ago",
			"talla viikolla":        "0 week ago",
			"toissa paivana":        "2 day ago",
			"tunnin sisalla":        "0 hour ago",
			"viime viikolla":        "1 week ago",
			"ensi viikolla":         "in 1 week",
			"toissa kuussa":         "2 month ago",
			"toissa vuonna":         "2 year ago",
			"tassa kuussa":          "0 month ago",
			"viime kuussa":          "1 month ago",
			"viime vuonna":          "1 year ago",
			"ensi kuussa":           "in 1 month",
			"ensi vuonna":           "in 1 year",
			"tana vuonna":           "0 year ago",
			"huomenna":              "in 1 day",
			"talla vk":              "0 week ago",
			"tassa kk":              "0 month ago",
			"viime kk":              "1 month ago",
			"viime vk":              "1 week ago",
			"ensi kk":               "in 1 month",
			"ensi vk":               "in 1 week",
			"viime v":               "1 year ago",
			"ensi v":                "in 1 year",
			"tana v":                "0 year ago",
			"tanaan":                "0 day ago",
			"eilen":                 "1 day ago",
			"huom":                  "in 1 day",
			"nyt":                   "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kuukauden paasta`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kuukautta sitten`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuuttia sitten`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kuukausi sitten`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuutin paasta`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuutti sitten`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekunnin paasta`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekuntia sitten`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekunti sitten`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) viikkoa sitten`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) paivaa sitten`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) paivan paasta`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) tunnin paasta`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) tuntia sitten`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) viikko sitten`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) viikon paasta`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vuoden paasta`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vuotta sitten`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) paiva sitten`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) tunti sitten`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vuosi sitten`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min paasta`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min sitten`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kk paasta`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kk sitten`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) pv paasta`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) pv sitten`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vk paasta`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vk sitten`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) s paasta`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) s sitten`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) t paasta`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) t sitten`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) v paasta`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) v sitten`), "$1 year ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* kuukauden paasta|\d+[.,]?\d* kuukautta sitten|\d+[.,]?\d* minuuttia sitten|\d+[.,]?\d* kuukausi sitten|\d+[.,]?\d* minuutin paasta|\d+[.,]?\d* minuutti sitten|\d+[.,]?\d* sekunnin paasta|\d+[.,]?\d* sekuntia sitten|\d+[.,]?\d* sekunti sitten|\d+[.,]?\d* viikkoa sitten|\d+[.,]?\d* paivaa sitten|\d+[.,]?\d* paivan paasta|\d+[.,]?\d* tunnin paasta|\d+[.,]?\d* tuntia sitten|\d+[.,]?\d* viikko sitten|\d+[.,]?\d* viikon paasta|\d+[.,]?\d* vuoden paasta|\d+[.,]?\d* vuotta sitten|\d+[.,]?\d* paiva sitten|\d+[.,]?\d* tunti sitten|\d+[.,]?\d* vuosi sitten|\d+[.,]?\d* min paasta|\d+[.,]?\d* min sitten|\d+[.,]?\d* kk paasta|\d+[.,]?\d* kk sitten|\d+[.,]?\d* pv paasta|\d+[.,]?\d* pv sitten|\d+[.,]?\d* vk paasta|\d+[.,]?\d* vk sitten|\d+[.,]?\d* s paasta|\d+[.,]?\d* s sitten|\d+[.,]?\d* t paasta|\d+[.,]?\d* t sitten|\d+[.,]?\d* v paasta|\d+[.,]?\d* v sitten)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* kuukauden paasta|\d+[.,]?\d* kuukautta sitten|\d+[.,]?\d* minuuttia sitten|\d+[.,]?\d* kuukausi sitten|\d+[.,]?\d* minuutin paasta|\d+[.,]?\d* minuutti sitten|\d+[.,]?\d* sekunnin paasta|\d+[.,]?\d* sekuntia sitten|\d+[.,]?\d* sekunti sitten|\d+[.,]?\d* viikkoa sitten|\d+[.,]?\d* paivaa sitten|\d+[.,]?\d* paivan paasta|\d+[.,]?\d* tunnin paasta|\d+[.,]?\d* tuntia sitten|\d+[.,]?\d* viikko sitten|\d+[.,]?\d* viikon paasta|\d+[.,]?\d* vuoden paasta|\d+[.,]?\d* vuotta sitten|\d+[.,]?\d* paiva sitten|\d+[.,]?\d* tunti sitten|\d+[.,]?\d* vuosi sitten|\d+[.,]?\d* min paasta|\d+[.,]?\d* min sitten|\d+[.,]?\d* kk paasta|\d+[.,]?\d* kk sitten|\d+[.,]?\d* pv paasta|\d+[.,]?\d* pv sitten|\d+[.,]?\d* vk paasta|\d+[.,]?\d* vk sitten|\d+[.,]?\d* s paasta|\d+[.,]?\d* s sitten|\d+[.,]?\d* t paasta|\d+[.,]?\d* t sitten|\d+[.,]?\d* v paasta|\d+[.,]?\d* v sitten)$`),
		KnownWords:      []string{"taman minuutin aikana", "taman tunnin aikana", "minuutin sisalla", "toissa viikolla", "talla viikolla", "toissa paivana", "tunnin sisalla", "viime viikolla", "ensi viikolla", "keskiviikkona", "toissa kuussa", "toissa vuonna", "maaliskuussa", "marraskuussa", "tassa kuussa", "viime kuussa", "viime vuonna", "ensi kuussa", "ensi vuonna", "heinakuussa", "helmikuussa", "huhtikuussa", "joulukuussa", "keskiviikko", "maaliskuuta", "maanantaina", "marraskuuta", "perjantaina", "sunnuntaina", "tammikuussa", "tana vuonna", "toukokuussa", "heinakuuta", "helmikuuta", "huhtikuuta", "joulukuuta", "kesakuussa", "lauantaina", "lokakuussa", "syyskuussa", "tammikuuta", "toukokuuta", "elokuussa", "kesakuuta", "kuukauden", "kuukautta", "lokakuuta", "maaliskuu", "maanantai", "marraskuu", "minuuttia", "perjantai", "sekunttia", "sunnuntai", "syyskuuta", "tiistaina", "torstaina", "elokuuta", "heinakuu", "helmikuu", "huhtikuu", "huomenna", "joulukuu", "kuluttua", "kuukausi", "lauantai", "minuutin", "minuutti", "sekunnin", "sekuntia", "sekuntin", "sekuntti", "talla vk", "tammikuu", "tassa kk", "toukokuu", "viime kk", "viime vk", "ensi kk", "ensi vk", "kesakuu", "lokakuu", "maalisk", "marrask", "sekunti", "syyskuu", "tiistai", "torstai", "viikkoa", "viime v", "elokuu", "ensi v", "heinak", "helmik", "huhtik", "jouluk", "maalis", "marras", "paasta", "paivaa", "paivan", "sitten", "tammik", "tana v", "tanaan", "toukok", "tunnin", "tuntia", "viikko", "viikon", "vuoden", "vuonna", "vuotta", "eilen", "heina", "helmi", "huhti", "joulu", "kesak", "lokak", "paiva", "syysk", "tammi", "touko", "tunti", "vuosi", "elok", "huom", "kesa", "loka", "pvaa", "syys", "elo", "gmt", "min", "nyt", "pva", "utc", "vko", ":n", "am", "ap", "ip", "ke", "kk", "la", "ma", "pe", "pm", "pv", "su", "ti", "to", "vk", "vv", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "p", "s", "t", "v", "z", "|"},
	})
}
