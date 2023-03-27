// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fr_Locale LocaleData

func init() {
	fr_Locale = merge(nil, LocaleData{
		Name:                  "fr",
		DateOrder:             "DMY",
		Charset:               []rune(`-bcdefghijlnorstuvyzèéû`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)d'une(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)d'un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)une(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d*)\s+h\s+(\d+[.,]?\d*)\s+min(\z|[^\pL\pM\d]|_)`), "${1}${2}h${3}m${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d*)h(\d+[.,]?\d*)m?(\z|[^\pL\pM\d]|_)`), "${1}${2}:${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+[.,]?\d*)\s?(minute|seconde|heure)(\z|[^\pL\pM\d]|_)`), "${1}${2} ${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+[.,]?\d*)\s?s(\z|[^\pL\pM\d]|_)`), "${1}${2} seconde${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+[.,]?\d*)\s?m(\z|[^\pL\pM\d]|_)`), "${1}${2} minute${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)moins\s(?:de\s)?(\d+[.,]?\d*)\s?h(\z|[^\pL\pM\d]|_)`), "${1}${2} heure${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)deux(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)trois(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)quatre(\z|[^\pL\pM\d]|_)`), "${1}4${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)cinq(\z|[^\pL\pM\d]|_)`), "${1}5${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)six(\z|[^\pL\pM\d]|_)`), "${1}6${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)sept(\z|[^\pL\pM\d]|_)`), "${1}7${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)huit(\z|[^\pL\pM\d]|_)`), "${1}8${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)neuf(\z|[^\pL\pM\d]|_)`), "${1}9${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)dix(\z|[^\pL\pM\d]|_)`), "${1}10${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)onze(\z|[^\pL\pM\d]|_)`), "${1}11${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)douze(\z|[^\pL\pM\d]|_)`), "${1}12${2}"},
		},
		Translations: map[string][]string{
			"septembre": {"september"},
			"decembre":  {"december"},
			"dimanche":  {"sunday"},
			"mercredi":  {"wednesday"},
			"novembre":  {"november"},
			"secondes":  {"second"},
			"semaines":  {"week"},
			"vendredi":  {"friday"},
			"environ":   {""},
			"fevrier":   {"february"},
			"janvier":   {"january"},
			"juillet":   {"july"},
			"minutes":   {"minute"},
			"octobre":   {"october"},
			"seconde":   {"second"},
			"semaine":   {"week"},
			"annees":    {"year"},
			"heures":    {"hour"},
			"il y a":    {"ago"},
			"minute":    {"minute"},
			"samedi":    {"saturday"},
			"annee":     {"year"},
			"apres":     {"in"},
			"avril":     {"april"},
			"heure":     {"hour"},
			"il ya":     {"ago"},
			"jeudi":     {"thursday"},
			"jours":     {"day"},
			"lundi":     {"monday"},
			"mardi":     {"tuesday"},
			"aout":      {"august"},
			"dans":      {"in"},
			"fevr":      {"february"},
			"janv":      {"january"},
			"jour":      {"day"},
			"juil":      {"july"},
			"juin":      {"june"},
			"mars":      {"march"},
			"mois":      {"month"},
			"sept":      {"september"},
			"ans":       {"year"},
			"aou":       {"august"},
			"avr":       {"april"},
			"dec":       {"december"},
			"dim":       {"sunday"},
			"fev":       {"february"},
			"gmt":       {"gmt"},
			"jan":       {"january"},
			"jeu":       {"thursday"},
			"jul":       {"july"},
			"jun":       {"june"},
			"lun":       {"monday"},
			"mai":       {"may"},
			"mar":       {"tuesday"},
			"mer":       {"wednesday"},
			"min":       {"minute"},
			"nov":       {"november"},
			"oct":       {"october"},
			"sam":       {"saturday"},
			"sem":       {"week"},
			"sep":       {"september"},
			"utc":       {"utc"},
			"ven":       {"friday"},
			"am":        {"am"},
			"an":        {"year"},
			"di":        {"sunday"},
			"en":        {"in"},
			"er":        {""},
			"et":        {""},
			"je":        {"thursday"},
			"le":        {""},
			"lu":        {"monday"},
			"ma":        {"tuesday"},
			"me":        {"wednesday"},
			"pm":        {"pm"},
			"sa":        {"saturday"},
			"ve":        {"friday"},
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
			"a":         {""},
			"h":         {"hour"},
			"j":         {"day"},
			"m":         {"month"},
			"s":         {"second"},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"la semaine prochaine": "in 1 week",
			"la semaine derniere":  "1 week ago",
			"l'annee prochaine":    "in 1 year",
			"l'annee derniere":     "1 year ago",
			"le mois prochain":     "in 1 month",
			"cette minute-ci":      "0 minute ago",
			"le mois dernier":      "1 month ago",
			"cette heure-ci":       "0 hour ago",
			"cette semaine":        "0 week ago",
			"apres-demain":         "in 2 day",
			"aujourd'hui":          "0 day ago",
			"cette annee":          "0 year ago",
			"avant-hier":           "2 day ago",
			"ce mois-ci":           "0 month ago",
			"maintenant":           "0 second ago",
			"demain":               "in 1 day",
			"hier":                 "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) secondes`), "$1 second ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) semaines`), "$1 week ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) minutes`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) seconde`), "$1 second ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) semaine`), "$1 week ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) secondes`), "in $1 second"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) semaines`), "in $1 week"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) heures`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) minute`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) minutes`), "in $1 minute"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) seconde`), "in $1 second"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) semaine`), "in $1 week"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) heure`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) jours`), "$1 day ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) heures`), "in $1 hour"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) minute`), "in $1 minute"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) jour`), "$1 day ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) mois`), "$1 month ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) heure`), "in $1 hour"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) jours`), "in $1 day"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) ans`), "$1 year ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) sem`), "$1 week ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) jour`), "in $1 day"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) mois`), "in $1 month"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) an`), "$1 year ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*)min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) ans`), "in $1 year"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) sem`), "in $1 week"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) a`), "$1 year ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) j`), "$1 day ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) m`), "$1 month ago"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) an`), "in $1 year"},
			{regexp.MustCompile(`(?i)il y a (\d+[.,]?\d*)h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) a`), "in $1 year"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) j`), "in $1 day"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) m`), "in $1 month"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*) s`), "in $1 second"},
			{regexp.MustCompile(`(?i)dans (\d+[.,]?\d*)h`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(il y a \d+[.,]?\d* secondes|il y a \d+[.,]?\d* semaines|il y a \d+[.,]?\d* minutes|il y a \d+[.,]?\d* seconde|il y a \d+[.,]?\d* semaine|dans \d+[.,]?\d* secondes|dans \d+[.,]?\d* semaines|il y a \d+[.,]?\d* heures|il y a \d+[.,]?\d* minute|dans \d+[.,]?\d* minutes|dans \d+[.,]?\d* seconde|dans \d+[.,]?\d* semaine|il y a \d+[.,]?\d* heure|il y a \d+[.,]?\d* jours|dans \d+[.,]?\d* heures|dans \d+[.,]?\d* minute|il y a \d+[.,]?\d* jour|il y a \d+[.,]?\d* mois|dans \d+[.,]?\d* heure|dans \d+[.,]?\d* jours|il y a \d+[.,]?\d* ans|il y a \d+[.,]?\d* min|il y a \d+[.,]?\d* sem|dans \d+[.,]?\d* jour|dans \d+[.,]?\d* mois|il y a \d+[.,]?\d* an|il y a \d+[.,]?\d*min|dans \d+[.,]?\d* ans|dans \d+[.,]?\d* min|dans \d+[.,]?\d* sem|il y a \d+[.,]?\d* a|il y a \d+[.,]?\d* h|il y a \d+[.,]?\d* j|il y a \d+[.,]?\d* m|il y a \d+[.,]?\d* s|dans \d+[.,]?\d* an|il y a \d+[.,]?\d*h|dans \d+[.,]?\d* a|dans \d+[.,]?\d* h|dans \d+[.,]?\d* j|dans \d+[.,]?\d* m|dans \d+[.,]?\d* s|dans \d+[.,]?\d*h)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(il y a \d+[.,]?\d* secondes|il y a \d+[.,]?\d* semaines|il y a \d+[.,]?\d* minutes|il y a \d+[.,]?\d* seconde|il y a \d+[.,]?\d* semaine|dans \d+[.,]?\d* secondes|dans \d+[.,]?\d* semaines|il y a \d+[.,]?\d* heures|il y a \d+[.,]?\d* minute|dans \d+[.,]?\d* minutes|dans \d+[.,]?\d* seconde|dans \d+[.,]?\d* semaine|il y a \d+[.,]?\d* heure|il y a \d+[.,]?\d* jours|dans \d+[.,]?\d* heures|dans \d+[.,]?\d* minute|il y a \d+[.,]?\d* jour|il y a \d+[.,]?\d* mois|dans \d+[.,]?\d* heure|dans \d+[.,]?\d* jours|il y a \d+[.,]?\d* ans|il y a \d+[.,]?\d* min|il y a \d+[.,]?\d* sem|dans \d+[.,]?\d* jour|dans \d+[.,]?\d* mois|il y a \d+[.,]?\d* an|il y a \d+[.,]?\d*min|dans \d+[.,]?\d* ans|dans \d+[.,]?\d* min|dans \d+[.,]?\d* sem|il y a \d+[.,]?\d* a|il y a \d+[.,]?\d* h|il y a \d+[.,]?\d* j|il y a \d+[.,]?\d* m|il y a \d+[.,]?\d* s|dans \d+[.,]?\d* an|il y a \d+[.,]?\d*h|dans \d+[.,]?\d* a|dans \d+[.,]?\d* h|dans \d+[.,]?\d* j|dans \d+[.,]?\d* m|dans \d+[.,]?\d* s|dans \d+[.,]?\d*h)$`),
		KnownWords:      []string{"la semaine prochaine", "la semaine derniere", "l'annee prochaine", "l'annee derniere", "le mois prochain", "cette minute-ci", "le mois dernier", "cette heure-ci", "cette semaine", "apres-demain", "aujourd'hui", "cette annee", "avant-hier", "ce mois-ci", "maintenant", "septembre", "decembre", "dimanche", "mercredi", "novembre", "secondes", "semaines", "vendredi", "environ", "fevrier", "janvier", "juillet", "minutes", "octobre", "seconde", "semaine", "annees", "demain", "heures", "il y a", "minute", "samedi", "annee", "apres", "avril", "heure", "il ya", "jeudi", "jours", "lundi", "mardi", "aout", "dans", "fevr", "hier", "janv", "jour", "juil", "juin", "mars", "mois", "sept", "ans", "aou", "avr", "dec", "dim", "fev", "gmt", "jan", "jeu", "jul", "jun", "lun", "mai", "mar", "mer", "min", "nov", "oct", "sam", "sem", "sep", "utc", "ven", "am", "an", "di", "en", "er", "et", "je", "le", "lu", "ma", "me", "pm", "sa", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "h", "j", "m", "s", "z", "|"},
	})
}

var fr_BE_Locale LocaleData

func init() {
	fr_BE_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-BE",
		DateOrder: "DMY",
	})
}

var fr_BF_Locale LocaleData

func init() {
	fr_BF_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-BF",
		DateOrder: "DMY",
	})
}

var fr_BI_Locale LocaleData

func init() {
	fr_BI_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-BI",
		DateOrder: "DMY",
	})
}

var fr_BJ_Locale LocaleData

func init() {
	fr_BJ_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-BJ",
		DateOrder: "DMY",
	})
}

var fr_BL_Locale LocaleData

func init() {
	fr_BL_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-BL",
		DateOrder: "DMY",
	})
}

var fr_CA_Locale LocaleData

func init() {
	fr_CA_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CA",
		DateOrder: "YMD",
		Translations: map[string][]string{
			"juill": {"july"},
		},
		KnownWords: []string{"la semaine prochaine", "la semaine derniere", "l'annee prochaine", "l'annee derniere", "le mois prochain", "cette minute-ci", "le mois dernier", "cette heure-ci", "cette semaine", "apres-demain", "aujourd'hui", "cette annee", "avant-hier", "ce mois-ci", "maintenant", "septembre", "decembre", "dimanche", "mercredi", "novembre", "secondes", "semaines", "vendredi", "environ", "fevrier", "janvier", "juillet", "minutes", "octobre", "seconde", "semaine", "annees", "demain", "heures", "il y a", "minute", "samedi", "annee", "apres", "avril", "heure", "il ya", "jeudi", "jours", "juill", "lundi", "mardi", "aout", "dans", "fevr", "hier", "janv", "jour", "juil", "juin", "mars", "mois", "sept", "ans", "aou", "avr", "dec", "dim", "fev", "gmt", "jan", "jeu", "jul", "jun", "lun", "mai", "mar", "mer", "min", "nov", "oct", "sam", "sem", "sep", "utc", "ven", "am", "an", "di", "en", "er", "et", "je", "le", "lu", "ma", "me", "pm", "sa", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "h", "j", "m", "s", "z", "|"},
	})
}

var fr_CD_Locale LocaleData

func init() {
	fr_CD_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CD",
		DateOrder: "DMY",
	})
}

var fr_CF_Locale LocaleData

func init() {
	fr_CF_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CF",
		DateOrder: "DMY",
	})
}

var fr_CG_Locale LocaleData

func init() {
	fr_CG_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CG",
		DateOrder: "DMY",
	})
}

var fr_CH_Locale LocaleData

func init() {
	fr_CH_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CH",
		DateOrder: "DMY",
	})
}

var fr_CI_Locale LocaleData

func init() {
	fr_CI_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CI",
		DateOrder: "DMY",
	})
}

var fr_CM_Locale LocaleData

func init() {
	fr_CM_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-CM",
		DateOrder: "DMY",
		Translations: map[string][]string{
			"matin": {"am"},
			"soir":  {"pm"},
			"mat":   {"am"},
		},
		KnownWords: []string{"la semaine prochaine", "la semaine derniere", "l'annee prochaine", "l'annee derniere", "le mois prochain", "cette minute-ci", "le mois dernier", "cette heure-ci", "cette semaine", "apres-demain", "aujourd'hui", "cette annee", "avant-hier", "ce mois-ci", "maintenant", "septembre", "decembre", "dimanche", "mercredi", "novembre", "secondes", "semaines", "vendredi", "environ", "fevrier", "janvier", "juillet", "minutes", "octobre", "seconde", "semaine", "annees", "demain", "heures", "il y a", "minute", "samedi", "annee", "apres", "avril", "heure", "il ya", "jeudi", "jours", "lundi", "mardi", "matin", "aout", "dans", "fevr", "hier", "janv", "jour", "juil", "juin", "mars", "mois", "sept", "soir", "ans", "aou", "avr", "dec", "dim", "fev", "gmt", "jan", "jeu", "jul", "jun", "lun", "mai", "mar", "mat", "mer", "min", "nov", "oct", "sam", "sem", "sep", "utc", "ven", "am", "an", "di", "en", "er", "et", "je", "le", "lu", "ma", "me", "pm", "sa", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "h", "j", "m", "s", "z", "|"},
	})
}

var fr_DJ_Locale LocaleData

func init() {
	fr_DJ_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-DJ",
		DateOrder: "DMY",
	})
}

var fr_DZ_Locale LocaleData

func init() {
	fr_DZ_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-DZ",
		DateOrder: "DMY",
	})
}

var fr_GA_Locale LocaleData

func init() {
	fr_GA_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-GA",
		DateOrder: "DMY",
	})
}

var fr_GF_Locale LocaleData

func init() {
	fr_GF_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-GF",
		DateOrder: "DMY",
	})
}

var fr_GN_Locale LocaleData

func init() {
	fr_GN_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-GN",
		DateOrder: "DMY",
	})
}

var fr_GP_Locale LocaleData

func init() {
	fr_GP_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-GP",
		DateOrder: "DMY",
	})
}

var fr_GQ_Locale LocaleData

func init() {
	fr_GQ_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-GQ",
		DateOrder: "DMY",
	})
}

var fr_HT_Locale LocaleData

func init() {
	fr_HT_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-HT",
		DateOrder: "DMY",
		Translations: map[string][]string{
			"sec": {"second"},
			"hr":  {"hour"},
			"jr":  {"day"},
		},
		KnownWords: []string{"la semaine prochaine", "la semaine derniere", "l'annee prochaine", "l'annee derniere", "le mois prochain", "cette minute-ci", "le mois dernier", "cette heure-ci", "cette semaine", "apres-demain", "aujourd'hui", "cette annee", "avant-hier", "ce mois-ci", "maintenant", "septembre", "decembre", "dimanche", "mercredi", "novembre", "secondes", "semaines", "vendredi", "environ", "fevrier", "janvier", "juillet", "minutes", "octobre", "seconde", "semaine", "annees", "demain", "heures", "il y a", "minute", "samedi", "annee", "apres", "avril", "heure", "il ya", "jeudi", "jours", "lundi", "mardi", "aout", "dans", "fevr", "hier", "janv", "jour", "juil", "juin", "mars", "mois", "sept", "ans", "aou", "avr", "dec", "dim", "fev", "gmt", "jan", "jeu", "jul", "jun", "lun", "mai", "mar", "mer", "min", "nov", "oct", "sam", "sec", "sem", "sep", "utc", "ven", "am", "an", "di", "en", "er", "et", "hr", "je", "jr", "le", "lu", "ma", "me", "pm", "sa", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "h", "j", "m", "s", "z", "|"},
	})
}

var fr_KM_Locale LocaleData

func init() {
	fr_KM_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-KM",
		DateOrder: "DMY",
	})
}

var fr_LU_Locale LocaleData

func init() {
	fr_LU_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-LU",
		DateOrder: "DMY",
	})
}

var fr_MA_Locale LocaleData

func init() {
	fr_MA_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MA",
		DateOrder: "DMY",
		Translations: map[string][]string{
			"jui": {"june"},
			"mar": {"march"},
		},
		KnownWords: []string{"la semaine prochaine", "la semaine derniere", "l'annee prochaine", "l'annee derniere", "le mois prochain", "cette minute-ci", "le mois dernier", "cette heure-ci", "cette semaine", "apres-demain", "aujourd'hui", "cette annee", "avant-hier", "ce mois-ci", "maintenant", "septembre", "decembre", "dimanche", "mercredi", "novembre", "secondes", "semaines", "vendredi", "environ", "fevrier", "janvier", "juillet", "minutes", "octobre", "seconde", "semaine", "annees", "demain", "heures", "il y a", "minute", "samedi", "annee", "apres", "avril", "heure", "il ya", "jeudi", "jours", "lundi", "mardi", "aout", "dans", "fevr", "hier", "janv", "jour", "juil", "juin", "mars", "mois", "sept", "ans", "aou", "avr", "dec", "dim", "fev", "gmt", "jan", "jeu", "jui", "jul", "jun", "lun", "mai", "mar", "mer", "min", "nov", "oct", "sam", "sem", "sep", "utc", "ven", "am", "an", "di", "en", "er", "et", "je", "le", "lu", "ma", "me", "pm", "sa", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "h", "j", "m", "s", "z", "|"},
	})
}

var fr_MC_Locale LocaleData

func init() {
	fr_MC_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MC",
		DateOrder: "DMY",
	})
}

var fr_MF_Locale LocaleData

func init() {
	fr_MF_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MF",
		DateOrder: "DMY",
	})
}

var fr_MG_Locale LocaleData

func init() {
	fr_MG_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MG",
		DateOrder: "DMY",
	})
}

var fr_ML_Locale LocaleData

func init() {
	fr_ML_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-ML",
		DateOrder: "DMY",
	})
}

var fr_MQ_Locale LocaleData

func init() {
	fr_MQ_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MQ",
		DateOrder: "DMY",
	})
}

var fr_MR_Locale LocaleData

func init() {
	fr_MR_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MR",
		DateOrder: "DMY",
	})
}

var fr_MU_Locale LocaleData

func init() {
	fr_MU_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-MU",
		DateOrder: "DMY",
	})
}

var fr_NC_Locale LocaleData

func init() {
	fr_NC_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-NC",
		DateOrder: "DMY",
	})
}

var fr_NE_Locale LocaleData

func init() {
	fr_NE_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-NE",
		DateOrder: "DMY",
	})
}

var fr_PF_Locale LocaleData

func init() {
	fr_PF_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-PF",
		DateOrder: "DMY",
	})
}

var fr_PM_Locale LocaleData

func init() {
	fr_PM_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-PM",
		DateOrder: "DMY",
	})
}

var fr_RE_Locale LocaleData

func init() {
	fr_RE_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-RE",
		DateOrder: "DMY",
	})
}

var fr_RW_Locale LocaleData

func init() {
	fr_RW_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-RW",
		DateOrder: "DMY",
	})
}

var fr_SC_Locale LocaleData

func init() {
	fr_SC_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-SC",
		DateOrder: "DMY",
	})
}

var fr_SN_Locale LocaleData

func init() {
	fr_SN_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-SN",
		DateOrder: "DMY",
	})
}

var fr_SY_Locale LocaleData

func init() {
	fr_SY_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-SY",
		DateOrder: "DMY",
	})
}

var fr_TD_Locale LocaleData

func init() {
	fr_TD_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-TD",
		DateOrder: "DMY",
	})
}

var fr_TG_Locale LocaleData

func init() {
	fr_TG_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-TG",
		DateOrder: "DMY",
	})
}

var fr_TN_Locale LocaleData

func init() {
	fr_TN_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-TN",
		DateOrder: "DMY",
	})
}

var fr_VU_Locale LocaleData

func init() {
	fr_VU_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-VU",
		DateOrder: "DMY",
	})
}

var fr_WF_Locale LocaleData

func init() {
	fr_WF_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-WF",
		DateOrder: "DMY",
	})
}

var fr_YT_Locale LocaleData

func init() {
	fr_YT_Locale = merge(&fr_Locale, LocaleData{
		Name:      "fr-YT",
		DateOrder: "DMY",
	})
}
