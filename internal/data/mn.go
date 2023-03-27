// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mn_Locale LocaleData

func init() {
	mn_Locale = merge(nil, LocaleData{
		Name:      "mn",
		DateOrder: "YMD",
		Charset:   []rune(`-cgtuzабвгдежзийклмнопрстухцчшэяёүө`),
		Translations: map[string][]string{
			"арван хоердугаар сар": {"december"},
			"арван нэгдүгээр сар":  {"november"},
			"гуравдугаар сар":      {"march"},
			"дөрөвдүгээр сар":      {"april"},
			"зургадугаар сар":      {"june"},
			"аравдугаар сар":       {"october"},
			"наимдугаар сар":       {"august"},
			"хоердугаар сар":       {"february"},
			"долдугаар сар":        {"july"},
			"нэгдүгээр сар":        {"january"},
			"тавдугаар сар":        {"may"},
			"есдүгээр сар":         {"september"},
			"долоо хоног":          {"week"},
			"10-р сар":             {"october"},
			"11-р сар":             {"november"},
			"12-р сар":             {"december"},
			"1-р сар":              {"january"},
			"2-р сар":              {"february"},
			"3-р сар":              {"march"},
			"4-р сар":              {"april"},
			"5-р сар":              {"may"},
			"6-р сар":              {"june"},
			"7-р сар":              {"july"},
			"8-р сар":              {"august"},
			"9-р сар":              {"september"},
			"баасан":               {"friday"},
			"лхагва":               {"wednesday"},
			"мягмар":               {"tuesday"},
			"секунд":               {"second"},
			"бямба":                {"saturday"},
			"даваа":                {"monday"},
			"минут":                {"minute"},
			"пүрэв":                {"thursday"},
			"өдөр":                 {"day"},
			"gmt":                  {"gmt"},
			"utc":                  {"utc"},
			"жил":                  {"year"},
			"мин":                  {"minute"},
			"ням":                  {"sunday"},
			"сар":                  {"month"},
			"сек":                  {"second"},
			"цаг":                  {"hour"},
			"7х":                   {"week"},
			"am":                   {"am"},
			"pm":                   {"pm"},
			"ба":                   {"friday"},
			"бя":                   {"saturday"},
			"да":                   {"monday"},
			"лх":                   {"wednesday"},
			"мя":                   {"tuesday"},
			"ня":                   {"sunday"},
			"пү":                   {"thursday"},
			"үх":                   {"pm"},
			"үө":                   {"am"},
			" ":                    {" "},
			"'":                    {""},
			"+":                    {"+"},
			",":                    {""},
			"-":                    {"-"},
			".":                    {"."},
			"/":                    {"/"},
			":":                    {":"},
			";":                    {""},
			"@":                    {""},
			"[":                    {""},
			"]":                    {""},
			"z":                    {"z"},
			"|":                    {""},
			"ц":                    {"hour"},
		},
		RelativeType: map[string]string{
			"өнгөрсөн долоо хоног": "1 week ago",
			"ирэх долоо хоног":     "in 1 week",
			"энэ долоо хоног":      "0 week ago",
			"өнгөрсөн жил":         "1 year ago",
			"өнгөрсөн сар":         "1 month ago",
			"энэ минут":            "0 minute ago",
			"ирэх жил":             "in 1 year",
			"ирэх сар":             "in 1 month",
			"маргааш":              "in 1 day",
			"энэ жил":              "0 year ago",
			"энэ сар":              "0 month ago",
			"энэ цаг":              "0 hour ago",
			"өнөөдөр":              "0 day ago",
			"өчигдөр":              "1 day ago",
			"одоо":                 "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секундын дараа`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) минутын дараа`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секундын өмнө`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 7х-иин дараа`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) жилиин дараа`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) минутын өмнө`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) цагиин дараа`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) өдриин дараа`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 7х-иин өмнө`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) жилиин өмнө`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сарын дараа`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) цагиин өмнө`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) өдриин өмнө`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сарын өмнө`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мин дараа`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сек дараа`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мин өмнө`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сек өмнө`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ц дараа`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ц өмнө`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) өдөрт`), "in $1 day"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* секундын дараа|\d+[.,]?\d* минутын дараа|\d+[.,]?\d* секундын өмнө|\d+[.,]?\d* 7х-иин дараа|\d+[.,]?\d* жилиин дараа|\d+[.,]?\d* минутын өмнө|\d+[.,]?\d* цагиин дараа|\d+[.,]?\d* өдриин дараа|\d+[.,]?\d* 7х-иин өмнө|\d+[.,]?\d* жилиин өмнө|\d+[.,]?\d* сарын дараа|\d+[.,]?\d* цагиин өмнө|\d+[.,]?\d* өдриин өмнө|\d+[.,]?\d* сарын өмнө|\d+[.,]?\d* мин дараа|\d+[.,]?\d* сек дараа|\d+[.,]?\d* мин өмнө|\d+[.,]?\d* сек өмнө|\d+[.,]?\d* ц дараа|\d+[.,]?\d* ц өмнө|\d+[.,]?\d* өдөрт)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* секундын дараа|\d+[.,]?\d* минутын дараа|\d+[.,]?\d* секундын өмнө|\d+[.,]?\d* 7х-иин дараа|\d+[.,]?\d* жилиин дараа|\d+[.,]?\d* минутын өмнө|\d+[.,]?\d* цагиин дараа|\d+[.,]?\d* өдриин дараа|\d+[.,]?\d* 7х-иин өмнө|\d+[.,]?\d* жилиин өмнө|\d+[.,]?\d* сарын дараа|\d+[.,]?\d* цагиин өмнө|\d+[.,]?\d* өдриин өмнө|\d+[.,]?\d* сарын өмнө|\d+[.,]?\d* мин дараа|\d+[.,]?\d* сек дараа|\d+[.,]?\d* мин өмнө|\d+[.,]?\d* сек өмнө|\d+[.,]?\d* ц дараа|\d+[.,]?\d* ц өмнө|\d+[.,]?\d* өдөрт)$`),
		KnownWords:      []string{"арван хоердугаар сар", "өнгөрсөн долоо хоног", "арван нэгдүгээр сар", "ирэх долоо хоног", "гуравдугаар сар", "дөрөвдүгээр сар", "зургадугаар сар", "энэ долоо хоног", "аравдугаар сар", "наимдугаар сар", "хоердугаар сар", "долдугаар сар", "нэгдүгээр сар", "тавдугаар сар", "есдүгээр сар", "өнгөрсөн жил", "өнгөрсөн сар", "долоо хоног", "энэ минут", "10-р сар", "11-р сар", "12-р сар", "ирэх жил", "ирэх сар", "1-р сар", "2-р сар", "3-р сар", "4-р сар", "5-р сар", "6-р сар", "7-р сар", "8-р сар", "9-р сар", "маргааш", "энэ жил", "энэ сар", "энэ цаг", "өнөөдөр", "өчигдөр", "баасан", "лхагва", "мягмар", "секунд", "бямба", "даваа", "минут", "пүрэв", "одоо", "өдөр", "gmt", "utc", "жил", "мин", "ням", "сар", "сек", "цаг", "7х", "am", "pm", "ба", "бя", "да", "лх", "мя", "ня", "пү", "үх", "үө", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ц"},
	})
}
