// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sr_Locale LocaleData

func init() {
	sr_Locale = merge(nil, LocaleData{
		Name:      "sr",
		DateOrder: "DMY.",
		Charset:   []rune(`cgtuzабвгдеиклмнопрстуфцчшјљћ`),
		Translations: map[string][]string{
			"понедељак": {"monday"},
			"пре подне": {"am"},
			"септембар": {"september"},
			"децембар":  {"december"},
			"новембар":  {"november"},
			"по подне":  {"pm"},
			"четвртак":  {"thursday"},
			"октобар":   {"october"},
			"фебруар":   {"february"},
			"август":    {"august"},
			"година":    {"year"},
			"недеља":    {"week", "sunday"},
			"секунд":    {"second"},
			"субота":    {"saturday"},
			"уторак":    {"tuesday"},
			"јануар":    {"january"},
			"април":     {"april"},
			"месец":     {"month"},
			"минут":     {"minute"},
			"петак":     {"friday"},
			"среда":     {"wednesday"},
			"март":      {"march"},
			"gmt":       {"gmt"},
			"utc":       {"utc"},
			"авг":       {"august"},
			"апр":       {"april"},
			"год":       {"year"},
			"дан":       {"day"},
			"дец":       {"december"},
			"мар":       {"march"},
			"мај":       {"may"},
			"мес":       {"month"},
			"мин":       {"minute"},
			"нед":       {"week", "sunday"},
			"нов":       {"november"},
			"окт":       {"october"},
			"пет":       {"friday"},
			"пон":       {"monday"},
			"сат":       {"hour"},
			"сек":       {"second"},
			"сеп":       {"september"},
			"сре":       {"wednesday"},
			"суб":       {"saturday"},
			"уто":       {"tuesday"},
			"феб":       {"february"},
			"чет":       {"thursday"},
			"јан":       {"january"},
			"јул":       {"july"},
			"јун":       {"june"},
			"am":        {"am"},
			"pm":        {"pm"},
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
			"г":         {"year"},
			"д":         {"day"},
			"м":         {"month"},
			"н":         {"week"},
			"с":         {"second"},
			"ч":         {"hour"},
		},
		RelativeType: map[string]string{
			"следећег месеца": "in 1 month",
			"прошлог месеца":  "1 month ago",
			"следеће године":  "in 1 year",
			"следеће недеље":  "in 1 week",
			"прошле године":   "1 year ago",
			"прошле недеље":   "1 week ago",
			"овог месеца":     "0 month ago",
			"овог минута":     "0 minute ago",
			"ове године":      "0 year ago",
			"ове недеље":      "0 week ago",
			"овог сата":       "0 hour ago",
			"данас":           "0 day ago",
			"сутра":           "in 1 day",
			"сада":            "0 second ago",
			"јуче":            "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) секунде`), "$1 second ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) секунди`), "$1 second ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) секунди`), "in $1 second"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) секунду`), "in $1 second"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) година`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) године`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) месеца`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) месеци`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) минута`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) недеља`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) недеље`), "$1 week ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) година`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) годину`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) месеци`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) минута`), "in $1 minute"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) недеља`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) недељу`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) месец`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) минут`), "in $1 minute"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) дана`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) сата`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) сати`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) дана`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) сати`), "in $1 hour"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) год`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) мес`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) мин`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) нед`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) сек`), "$1 second ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) год`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) дан`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) мес`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) мин`), "in $1 minute"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) нед`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) сат`), "in $1 hour"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) сек`), "in $1 second"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) г`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) д`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) м`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) н`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) с`), "$1 second ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) ч`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) г`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) д`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) м`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) н`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) с`), "in $1 second"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) ч`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(пре \d+[.,]?\d* секунде|пре \d+[.,]?\d* секунди|за \d+[.,]?\d* секунди|за \d+[.,]?\d* секунду|пре \d+[.,]?\d* година|пре \d+[.,]?\d* године|пре \d+[.,]?\d* месеца|пре \d+[.,]?\d* месеци|пре \d+[.,]?\d* минута|пре \d+[.,]?\d* недеља|пре \d+[.,]?\d* недеље|за \d+[.,]?\d* година|за \d+[.,]?\d* годину|за \d+[.,]?\d* месеци|за \d+[.,]?\d* минута|за \d+[.,]?\d* недеља|за \d+[.,]?\d* недељу|за \d+[.,]?\d* месец|за \d+[.,]?\d* минут|пре \d+[.,]?\d* дана|пре \d+[.,]?\d* сата|пре \d+[.,]?\d* сати|за \d+[.,]?\d* дана|за \d+[.,]?\d* сати|пре \d+[.,]?\d* год|пре \d+[.,]?\d* мес|пре \d+[.,]?\d* мин|пре \d+[.,]?\d* нед|пре \d+[.,]?\d* сек|за \d+[.,]?\d* год|за \d+[.,]?\d* дан|за \d+[.,]?\d* мес|за \d+[.,]?\d* мин|за \d+[.,]?\d* нед|за \d+[.,]?\d* сат|за \d+[.,]?\d* сек|пре \d+[.,]?\d* г|пре \d+[.,]?\d* д|пре \d+[.,]?\d* м|пре \d+[.,]?\d* н|пре \d+[.,]?\d* с|пре \d+[.,]?\d* ч|за \d+[.,]?\d* г|за \d+[.,]?\d* д|за \d+[.,]?\d* м|за \d+[.,]?\d* н|за \d+[.,]?\d* с|за \d+[.,]?\d* ч)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(пре \d+[.,]?\d* секунде|пре \d+[.,]?\d* секунди|за \d+[.,]?\d* секунди|за \d+[.,]?\d* секунду|пре \d+[.,]?\d* година|пре \d+[.,]?\d* године|пре \d+[.,]?\d* месеца|пре \d+[.,]?\d* месеци|пре \d+[.,]?\d* минута|пре \d+[.,]?\d* недеља|пре \d+[.,]?\d* недеље|за \d+[.,]?\d* година|за \d+[.,]?\d* годину|за \d+[.,]?\d* месеци|за \d+[.,]?\d* минута|за \d+[.,]?\d* недеља|за \d+[.,]?\d* недељу|за \d+[.,]?\d* месец|за \d+[.,]?\d* минут|пре \d+[.,]?\d* дана|пре \d+[.,]?\d* сата|пре \d+[.,]?\d* сати|за \d+[.,]?\d* дана|за \d+[.,]?\d* сати|пре \d+[.,]?\d* год|пре \d+[.,]?\d* мес|пре \d+[.,]?\d* мин|пре \d+[.,]?\d* нед|пре \d+[.,]?\d* сек|за \d+[.,]?\d* год|за \d+[.,]?\d* дан|за \d+[.,]?\d* мес|за \d+[.,]?\d* мин|за \d+[.,]?\d* нед|за \d+[.,]?\d* сат|за \d+[.,]?\d* сек|пре \d+[.,]?\d* г|пре \d+[.,]?\d* д|пре \d+[.,]?\d* м|пре \d+[.,]?\d* н|пре \d+[.,]?\d* с|пре \d+[.,]?\d* ч|за \d+[.,]?\d* г|за \d+[.,]?\d* д|за \d+[.,]?\d* м|за \d+[.,]?\d* н|за \d+[.,]?\d* с|за \d+[.,]?\d* ч)$`),
		KnownWords:      []string{"следећег месеца", "прошлог месеца", "следеће године", "следеће недеље", "прошле године", "прошле недеље", "овог месеца", "овог минута", "ове године", "ове недеље", "овог сата", "понедељак", "пре подне", "септембар", "децембар", "новембар", "по подне", "четвртак", "октобар", "фебруар", "август", "година", "недеља", "секунд", "субота", "уторак", "јануар", "април", "данас", "месец", "минут", "петак", "среда", "сутра", "март", "сада", "јуче", "gmt", "utc", "авг", "апр", "год", "дан", "дец", "мар", "мај", "мес", "мин", "нед", "нов", "окт", "пет", "пон", "сат", "сек", "сеп", "сре", "суб", "уто", "феб", "чет", "јан", "јул", "јун", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "г", "д", "м", "н", "с", "ч"},
	})
}
