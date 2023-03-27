// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ru_Locale LocaleData

func init() {
	ru_Locale = merge(nil, LocaleData{
		Name:                  "ru",
		DateOrder:             "DMY",
		Charset:               []rune(`cgtuzабвгдезийклмнопрстуфцчшщыьэюя`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)од(?:на|ну|ни|нои|ин)(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)дв(?:а|е|ое|ух)(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пар[ауы](\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)три(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)четыре(\z|[^\pL\pM\d]|_)`), "${1}4${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пять(\z|[^\pL\pM\d]|_)`), "${1}5${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)шесть(\z|[^\pL\pM\d]|_)`), "${1}6${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)семь(\z|[^\pL\pM\d]|_)`), "${1}7${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)восемь(\z|[^\pL\pM\d]|_)`), "${1}8${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)девять(\z|[^\pL\pM\d]|_)`), "${1}9${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)десять(\z|[^\pL\pM\d]|_)`), "${1}10${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)одиннадцать(\z|[^\pL\pM\d]|_)`), "${1}11${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)двенадцать(\z|[^\pL\pM\d]|_)`), "${1}12${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пятнадцать(\z|[^\pL\pM\d]|_)`), "${1}15${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)двадцать(\z|[^\pL\pM\d]|_)`), "${1}20${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)тридцать(\z|[^\pL\pM\d]|_)`), "${1}30${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)сорок(\z|[^\pL\pM\d]|_)`), "${1}40${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)пятьдесят(\z|[^\pL\pM\d]|_)`), "${1}50${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)несколько секунд(\z|[^\pL\pM\d]|_)`), "${1}44 секунды${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полчаса(\z|[^\pL\pM\d]|_)`), "${1}30 минут${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полгода(\z|[^\pL\pM\d]|_)`), "${1}6 месяцев${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полтора часа(\z|[^\pL\pM\d]|_)`), "${1}90 минут${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)полтора года(\z|[^\pL\pM\d]|_)`), "${1}18 месяцев${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)секунд[уы](\z|[^\pL\pM\d]|_)`), "${1}${2}1 секунду${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)минут[уы](\z|[^\pL\pM\d]|_)`), "${1}${2}1 минуту${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)часа?(\z|[^\pL\pM\d]|_)`), "${1}${2}1 час${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)(?:день|дня)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 день${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)сут(?:ки|ок)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 сутки${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)недел[юи](\z|[^\pL\pM\d]|_)`), "${1}${2}1 неделю${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)месяца?(\z|[^\pL\pM\d]|_)`), "${1}${2}1 месяц${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)((?:через|спустя|в течение)\s+)года?(\z|[^\pL\pM\d]|_)`), "${1}${2}1 год${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)секунду(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 секунду${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)минуту(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 минуту${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)час(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 час${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)день(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 день${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)сутки(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 сутки${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)неделю(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 неделю${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)месяц(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 месяц${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)([^\d]\s+|^)год(\s+назад)(\z|[^\pL\pM\d]|_)`), "${1}${2}1 год${3}${4}"},
		},
		Translations: map[string][]string{
			"воскресение": {"sunday"},
			"воскресенье": {"sunday"},
			"понедельник": {"monday"},
			"в течение":   {"in"},
			"примерно":    {""},
			"сентябрь":    {"september"},
			"сентября":    {"september"},
			"августа":     {"august"},
			"вторник":     {"tuesday"},
			"декабрь":     {"december"},
			"декабря":     {"december"},
			"месяцев":     {"month"},
			"октябрь":     {"october"},
			"октября":     {"october"},
			"пятница":     {"friday"},
			"пятницу":     {"friday"},
			"секунда":     {"second"},
			"секунду":     {"second"},
			"секунды":     {"second"},
			"суббота":     {"saturday"},
			"субботу":     {"saturday"},
			"февраль":     {"february"},
			"февраля":     {"february"},
			"четверг":     {"thursday"},
			"август":      {"august"},
			"апрель":      {"april"},
			"апреля":      {"april"},
			"месяца":      {"month"},
			"минута":      {"minute"},
			"минуту":      {"minute"},
			"минуты":      {"minute"},
			"недели":      {"week"},
			"недель":      {"week"},
			"неделю":      {"week"},
			"неделя":      {"week"},
			"ноябрь":      {"november"},
			"ноября":      {"november"},
			"секунд":      {"second"},
			"спустя":      {"in"},
			"январь":      {"january"},
			"января":      {"january"},
			"марта":       {"march"},
			"месяц":       {"month"},
			"минут":       {"minute"},
			"назад":       {"ago"},
			"около":       {""},
			"среда":       {"wednesday"},
			"среду":       {"wednesday"},
			"сутки":       {"day"},
			"суток":       {"day"},
			"часов":       {"hour"},
			"через":       {"in"},
			"года":        {"year"},
			"году":        {"year"},
			"день":        {"day"},
			"днеи":        {"day"},
			"июль":        {"july"},
			"июля":        {"july"},
			"июнь":        {"june"},
			"июня":        {"june"},
			"март":        {"march"},
			"нояб":        {"november"},
			"сент":        {"september"},
			"февр":        {"february"},
			"часа":        {"hour"},
			"gmt":         {"gmt"},
			"utc":         {"utc"},
			"авг":         {"august"},
			"апр":         {"april"},
			"вск":         {"sunday"},
			"втр":         {"tuesday"},
			"год":         {"year"},
			"дек":         {"december"},
			"дня":         {"day"},
			"июл":         {"july"},
			"июн":         {"june"},
			"лет":         {"year"},
			"маи":         {"may"},
			"мар":         {"march"},
			"мая":         {"may"},
			"мес":         {"month"},
			"мин":         {"minute"},
			"нед":         {"week"},
			"ноя":         {"november"},
			"окт":         {"october"},
			"пнд":         {"monday"},
			"птн":         {"friday"},
			"сбт":         {"saturday"},
			"сек":         {"second"},
			"сен":         {"september"},
			"срд":         {"wednesday"},
			"фев":         {"february"},
			"час":         {"hour"},
			"чтв":         {"thursday"},
			"янв":         {"january"},
			"am":          {"am"},
			"pm":          {"pm"},
			"во":          {""},
			"вс":          {"sunday"},
			"вт":          {"tuesday"},
			"дн":          {"day"},
			"дп":          {"am"},
			"пн":          {"monday"},
			"пп":          {"pm"},
			"пт":          {"friday"},
			"сб":          {"saturday"},
			"ср":          {"wednesday"},
			"чт":          {"thursday"},
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
			"в":           {""},
			"г":           {"year"},
			"и":           {""},
			"с":           {"second"},
			"ч":           {"hour"},
		},
		RelativeType: map[string]string{
			"на следующеи неделе": "in 1 week",
			"в следующем месяце":  "in 1 month",
			"на прошлои неделе":   "1 week ago",
			"в прошлом месяце":    "1 month ago",
			"в следующем году":    "in 1 year",
			"послепослезавтра":    "in 3 day",
			"в прошлом году":      "1 year ago",
			"на этои неделе":      "0 week ago",
			"в этом месяце":       "0 month ago",
			"в эту минуту":        "0 minute ago",
			"в этом году":         "0 year ago",
			"в этом часе":         "0 hour ago",
			"послезавтра":         "in 2 day",
			"позавчера":           "2 day ago",
			"сегодня":             "0 day ago",
			"завтра":              "in 1 day",
			"сеичас":              "0 second ago",
			"вчера":               "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секунду назад`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секунды назад`), "$1 second ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) секунду`), "in $1 second"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) секунды`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) месяца назад`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) минуту назад`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) минуты назад`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) недели назад`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) неделю назад`), "$1 week ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) месяца`), "in $1 month"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) минуту`), "in $1 minute"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) минуты`), "in $1 minute"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) недели`), "in $1 week"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) неделю`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) месяц назад`), "$1 month ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) месяц`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) года назад`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) день назад`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) часа назад`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) года`), "in $1 year"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) день`), "in $1 day"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) часа`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) год назад`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) дня назад`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мес назад`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мин назад`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) нед назад`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сек назад`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) час назад`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) год`), "in $1 year"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) дня`), "in $1 day"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) мес`), "in $1 month"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) мин`), "in $1 minute"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) нед`), "in $1 week"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) сек`), "in $1 second"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) час`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) дн назад`), "$1 day ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) дн`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) г назад`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) д назад`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ч назад`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) г`), "in $1 year"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) д`), "in $1 day"},
			{regexp.MustCompile(`(?i)через (\d+[.,]?\d*) ч`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* секунду назад|\d+[.,]?\d* секунды назад|через \d+[.,]?\d* секунду|через \d+[.,]?\d* секунды|\d+[.,]?\d* месяца назад|\d+[.,]?\d* минуту назад|\d+[.,]?\d* минуты назад|\d+[.,]?\d* недели назад|\d+[.,]?\d* неделю назад|через \d+[.,]?\d* месяца|через \d+[.,]?\d* минуту|через \d+[.,]?\d* минуты|через \d+[.,]?\d* недели|через \d+[.,]?\d* неделю|\d+[.,]?\d* месяц назад|через \d+[.,]?\d* месяц|\d+[.,]?\d* года назад|\d+[.,]?\d* день назад|\d+[.,]?\d* часа назад|через \d+[.,]?\d* года|через \d+[.,]?\d* день|через \d+[.,]?\d* часа|\d+[.,]?\d* год назад|\d+[.,]?\d* дня назад|\d+[.,]?\d* мес назад|\d+[.,]?\d* мин назад|\d+[.,]?\d* нед назад|\d+[.,]?\d* сек назад|\d+[.,]?\d* час назад|через \d+[.,]?\d* год|через \d+[.,]?\d* дня|через \d+[.,]?\d* мес|через \d+[.,]?\d* мин|через \d+[.,]?\d* нед|через \d+[.,]?\d* сек|через \d+[.,]?\d* час|\d+[.,]?\d* дн назад|через \d+[.,]?\d* дн|\d+[.,]?\d* г назад|\d+[.,]?\d* д назад|\d+[.,]?\d* ч назад|через \d+[.,]?\d* г|через \d+[.,]?\d* д|через \d+[.,]?\d* ч)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* секунду назад|\d+[.,]?\d* секунды назад|через \d+[.,]?\d* секунду|через \d+[.,]?\d* секунды|\d+[.,]?\d* месяца назад|\d+[.,]?\d* минуту назад|\d+[.,]?\d* минуты назад|\d+[.,]?\d* недели назад|\d+[.,]?\d* неделю назад|через \d+[.,]?\d* месяца|через \d+[.,]?\d* минуту|через \d+[.,]?\d* минуты|через \d+[.,]?\d* недели|через \d+[.,]?\d* неделю|\d+[.,]?\d* месяц назад|через \d+[.,]?\d* месяц|\d+[.,]?\d* года назад|\d+[.,]?\d* день назад|\d+[.,]?\d* часа назад|через \d+[.,]?\d* года|через \d+[.,]?\d* день|через \d+[.,]?\d* часа|\d+[.,]?\d* год назад|\d+[.,]?\d* дня назад|\d+[.,]?\d* мес назад|\d+[.,]?\d* мин назад|\d+[.,]?\d* нед назад|\d+[.,]?\d* сек назад|\d+[.,]?\d* час назад|через \d+[.,]?\d* год|через \d+[.,]?\d* дня|через \d+[.,]?\d* мес|через \d+[.,]?\d* мин|через \d+[.,]?\d* нед|через \d+[.,]?\d* сек|через \d+[.,]?\d* час|\d+[.,]?\d* дн назад|через \d+[.,]?\d* дн|\d+[.,]?\d* г назад|\d+[.,]?\d* д назад|\d+[.,]?\d* ч назад|через \d+[.,]?\d* г|через \d+[.,]?\d* д|через \d+[.,]?\d* ч)$`),
		KnownWords:      []string{"на следующеи неделе", "в следующем месяце", "на прошлои неделе", "в прошлом месяце", "в следующем году", "послепослезавтра", "в прошлом году", "на этои неделе", "в этом месяце", "в эту минуту", "в этом году", "в этом часе", "воскресение", "воскресенье", "понедельник", "послезавтра", "в течение", "позавчера", "примерно", "сентябрь", "сентября", "августа", "вторник", "декабрь", "декабря", "месяцев", "октябрь", "октября", "пятница", "пятницу", "сегодня", "секунда", "секунду", "секунды", "суббота", "субботу", "февраль", "февраля", "четверг", "август", "апрель", "апреля", "завтра", "месяца", "минута", "минуту", "минуты", "недели", "недель", "неделю", "неделя", "ноябрь", "ноября", "сеичас", "секунд", "спустя", "январь", "января", "вчера", "марта", "месяц", "минут", "назад", "около", "среда", "среду", "сутки", "суток", "часов", "через", "года", "году", "день", "днеи", "июль", "июля", "июнь", "июня", "март", "нояб", "сент", "февр", "часа", "gmt", "utc", "авг", "апр", "вск", "втр", "год", "дек", "дня", "июл", "июн", "лет", "маи", "мар", "мая", "мес", "мин", "нед", "ноя", "окт", "пнд", "птн", "сбт", "сек", "сен", "срд", "фев", "час", "чтв", "янв", "am", "pm", "во", "вс", "вт", "дн", "дп", "пн", "пп", "пт", "сб", "ср", "чт", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "в", "г", "и", "с", "ч"},
	})
}

var ru_BY_Locale LocaleData

func init() {
	ru_BY_Locale = merge(&ru_Locale, LocaleData{
		Name:      "ru-BY",
		DateOrder: "DMY",
	})
}

var ru_KG_Locale LocaleData

func init() {
	ru_KG_Locale = merge(&ru_Locale, LocaleData{
		Name:      "ru-KG",
		DateOrder: "DMY",
	})
}

var ru_KZ_Locale LocaleData

func init() {
	ru_KZ_Locale = merge(&ru_Locale, LocaleData{
		Name:      "ru-KZ",
		DateOrder: "DMY",
	})
}

var ru_MD_Locale LocaleData

func init() {
	ru_MD_Locale = merge(&ru_Locale, LocaleData{
		Name:      "ru-MD",
		DateOrder: "DMY",
	})
}

var ru_UA_Locale LocaleData

func init() {
	ru_UA_Locale = merge(&ru_Locale, LocaleData{
		Name:      "ru-UA",
		DateOrder: "DMY",
	})
}
