// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var zu_Locale LocaleData

func init() {
	zu_Locale = merge(nil, LocaleData{
		Name:      "zu",
		DateOrder: "MDY",
		Charset:   []rune(`bcdefghijklnorstuvwyz`),
		Translations: map[string][]string{
			"ulwesithathu": {"wednesday"},
			"ulwesihlanu":  {"friday"},
			"umsombuluko":  {"monday"},
			"ulwesibili":   {"tuesday"},
			"umasingana":   {"january"},
			"februwari":    {"february"},
			"isekhondi":    {"second"},
			"septhemba":    {"september"},
			"umgqibelo":    {"saturday"},
			"iminithi":     {"minute"},
			"januwari":     {"january"},
			"ulwesine":     {"thursday"},
			"disemba":      {"december"},
			"ephreli":      {"april"},
			"inyanga":      {"month"},
			"novemba":      {"november"},
			"okthoba":      {"october"},
			"agasti":       {"august"},
			"isonto":       {"sunday"},
			"julayi":       {"july"},
			"unyaka":       {"year"},
			"ihora":        {"hour"},
			"iviki":        {"week"},
			"mashi":        {"march"},
			"usuku":        {"day"},
			"juni":         {"june"},
			"meyi":         {"may"},
			"aga":          {"august"},
			"bil":          {"tuesday"},
			"dis":          {"december"},
			"eph":          {"april"},
			"feb":          {"february"},
			"gmt":          {"gmt"},
			"hla":          {"friday"},
			"jan":          {"january"},
			"jul":          {"july"},
			"jun":          {"june"},
			"mas":          {"march"},
			"mey":          {"may"},
			"mgq":          {"saturday"},
			"mso":          {"monday"},
			"nov":          {"november"},
			"okt":          {"october"},
			"sep":          {"september"},
			"sin":          {"thursday"},
			"son":          {"sunday"},
			"tha":          {"wednesday"},
			"utc":          {"utc"},
			"am":           {"am"},
			"pm":           {"pm"},
			" ":            {" "},
			"'":            {""},
			"+":            {"+"},
			",":            {""},
			"-":            {"-"},
			".":            {"."},
			"/":            {"/"},
			":":            {":"},
			";":            {""},
			"@":            {""},
			"[":            {""},
			"]":            {""},
			"z":            {"z"},
			"|":            {""},
		},
		RelativeType: map[string]string{
			"onyakeni odlule": "1 year ago",
			"inyanga edlule":  "1 month ago",
			"iviki eledlule":  "1 week ago",
			"inyanga ezayo":   "in 1 month",
			"iviki elizayo":   "in 1 week",
			"leli minithi":    "0 minute ago",
			"unyaka ozayo":    "in 1 year",
			"kulo nyaka":      "0 year ago",
			"le nyanga":       "0 month ago",
			"leli hora":       "0 hour ago",
			"leli viki":       "0 week ago",
			"namhlanje":       "0 day ago",
			"kusasa":          "in 1 day",
			"izolo":           "1 day ago",
			"manje":           "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) amasekhondi edlule`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) isekhondi eledlule`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) izinyanga ezedlule`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) amaminithi edlule`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) iminithi eledlule`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) izinsuku ezedlule`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) iminyaka edlule`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) usuku olwedlule`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) amahora edlule`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ihora eledlule`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) inyanga edlule`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) unyaka odlule`), "$1 year ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* amasekhondi edlule|\d+[.,]?\d* isekhondi eledlule|\d+[.,]?\d* izinyanga ezedlule|\d+[.,]?\d* amaminithi edlule|\d+[.,]?\d* iminithi eledlule|\d+[.,]?\d* izinsuku ezedlule|\d+[.,]?\d* iminyaka edlule|\d+[.,]?\d* usuku olwedlule|\d+[.,]?\d* amahora edlule|\d+[.,]?\d* ihora eledlule|\d+[.,]?\d* inyanga edlule|\d+[.,]?\d* unyaka odlule)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* amasekhondi edlule|\d+[.,]?\d* isekhondi eledlule|\d+[.,]?\d* izinyanga ezedlule|\d+[.,]?\d* amaminithi edlule|\d+[.,]?\d* iminithi eledlule|\d+[.,]?\d* izinsuku ezedlule|\d+[.,]?\d* iminyaka edlule|\d+[.,]?\d* usuku olwedlule|\d+[.,]?\d* amahora edlule|\d+[.,]?\d* ihora eledlule|\d+[.,]?\d* inyanga edlule|\d+[.,]?\d* unyaka odlule)$`),
		KnownWords:      []string{"onyakeni odlule", "inyanga edlule", "iviki eledlule", "inyanga ezayo", "iviki elizayo", "leli minithi", "ulwesithathu", "unyaka ozayo", "ulwesihlanu", "umsombuluko", "kulo nyaka", "ulwesibili", "umasingana", "februwari", "isekhondi", "le nyanga", "leli hora", "leli viki", "namhlanje", "septhemba", "umgqibelo", "iminithi", "januwari", "ulwesine", "disemba", "ephreli", "inyanga", "novemba", "okthoba", "agasti", "isonto", "julayi", "kusasa", "unyaka", "ihora", "iviki", "izolo", "manje", "mashi", "usuku", "juni", "meyi", "aga", "bil", "dis", "eph", "feb", "gmt", "hla", "jan", "jul", "jun", "mas", "mey", "mgq", "mso", "nov", "okt", "sep", "sin", "son", "tha", "utc", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
