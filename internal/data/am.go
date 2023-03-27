// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var am_Locale LocaleData

func init() {
	am_Locale = merge(nil, LocaleData{
		Name:      "am",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzሁህለላልሐሑመሙሚማሜምሣረሩሪሬርሰሳሴስቀቂቃቅበቡባብቬቴትቶችነናንኖኞአኤእኦከክወዋውዓዕዚዛየዩያይደዲዳድጁጃገጥፈፌፕ`),
		Translations: map[string][]string{
			"ሴፕቴምበር": {"september"},
			"ኖቬምበር":  {"november"},
			"ኦክቶበር":  {"october"},
			"ዲሴምበር":  {"december"},
			"ጃንዩወሪ":  {"january"},
			"ፌብሩወሪ":  {"february"},
			"ማክሰኞ":   {"tuesday"},
			"ሰከንድ":   {"second"},
			"ሳምንት":   {"week"},
			"ኤፕሪል":   {"april"},
			"ኦገስት":   {"august"},
			"ከሰዓት":   {"pm"},
			"gmt":    {"gmt"},
			"utc":    {"utc"},
			"ሐሙስ":    {"thursday"},
			"ማርች":    {"march"},
			"ማክሰ":    {"tuesday"},
			"ረቡዕ":    {"wednesday"},
			"ሰዓት":    {"hour"},
			"ሴፕቴ":    {"september"},
			"ቅዳሜ":    {"saturday"},
			"ኖቬም":    {"november"},
			"ኤፕሪ":    {"april"},
			"እሑድ":    {"sunday"},
			"ኦክቶ":    {"october"},
			"ኦገስ":    {"august"},
			"ዓመት":    {"year"},
			"ዓርብ":    {"friday"},
			"ደቂቃ":    {"minute"},
			"ዲሴም":    {"december"},
			"ጁላይ":    {"july"},
			"ጃንዩ":    {"january"},
			"ጥዋት":    {"am"},
			"ፌብሩ":    {"february"},
			"am":     {"am"},
			"pm":     {"pm"},
			"ሜይ":     {"may"},
			"ሰኞ":     {"monday"},
			"ቀን":     {"day"},
			"ወር":     {"month"},
			"ጁን":     {"june"},
			" ":      {" "},
			"'":      {""},
			"+":      {"+"},
			",":      {""},
			"-":      {"-"},
			".":      {"."},
			"/":      {"/"},
			":":      {":"},
			";":      {""},
			"@":      {""},
			"[":      {""},
			"]":      {""},
			"z":      {"z"},
			"|":      {""},
		},
		RelativeType: map[string]string{
			"የሚቀጥለው ሳምንት": "in 1 week",
			"የሚቀጥለው ዓመት":  "in 1 year",
			"ባለፈው ሳምንት":   "1 week ago",
			"የሚቀጥለው ወር":   "in 1 month",
			"ያለፈው ሳምንት":   "1 week ago",
			"በዚህ ሣምንት":    "0 week ago",
			"በዚህ ሳምንት":    "0 week ago",
			"ያለፈው ዓመት":    "1 year ago",
			"በዚህ ዓመት":     "0 year ago",
			"ያለፈው ወር":     "1 month ago",
			"በዚህ ወር":      "0 month ago",
			"ይህ ሰዓት":      "0 hour ago",
			"ይህ ደቂቃ":      "0 minute ago",
			"ትላንትና":       "1 day ago",
			"ትናንት":        "1 day ago",
			"አሁን":         "0 second ago",
			"ነገ":          "in 1 day",
			"ዛሬ":          "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ሰከንዶች ውስጥ`), "in $1 second"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ሳምንታት ውስጥ`), "in $1 week"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ደቂቃዎች ውስጥ`), "in $1 minute"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ሰከንዶች በፊት`), "$1 second ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ሳምንታት በፊት`), "$1 week ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ደቂቃዎች በፊት`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ሰከንድ ውስጥ`), "in $1 second"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ሰዓቶች ውስጥ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ሳምንት ውስጥ`), "in $1 week"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ዓመታት ውስጥ`), "in $1 year"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ሰከንድ በፊት`), "$1 second ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ሰዓቶች በፊት`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ሳምንት በፊት`), "$1 week ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ዓመታት በፊት`), "$1 year ago"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ሰዓት ውስጥ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ቀናት ውስጥ`), "in $1 day"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ቀኖች ውስጥ`), "in $1 day"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ወራት ውስጥ`), "in $1 month"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ደቂቃ ውስጥ`), "in $1 minute"},
			{regexp.MustCompile(`(?i)ከ (\d+[.,]?\d*) ቀን በፊት`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ሰዓት በፊት`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ቀናት በፊት`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ቀኖች በፊት`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ወራት በፊት`), "$1 month ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ዓመት በፊት`), "$1 year ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ደቂቃ በፊት`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ቀን ውስጥ`), "in $1 day"},
			{regexp.MustCompile(`(?i)በ(\d+[.,]?\d*) ወር ውስጥ`), "in $1 month"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ቀን በፊት`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ከ(\d+[.,]?\d*) ወር በፊት`), "$1 month ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(በ\d+[.,]?\d* ሰከንዶች ውስጥ|በ\d+[.,]?\d* ሳምንታት ውስጥ|በ\d+[.,]?\d* ደቂቃዎች ውስጥ|ከ\d+[.,]?\d* ሰከንዶች በፊት|ከ\d+[.,]?\d* ሳምንታት በፊት|ከ\d+[.,]?\d* ደቂቃዎች በፊት|በ\d+[.,]?\d* ሰከንድ ውስጥ|በ\d+[.,]?\d* ሰዓቶች ውስጥ|በ\d+[.,]?\d* ሳምንት ውስጥ|በ\d+[.,]?\d* ዓመታት ውስጥ|ከ\d+[.,]?\d* ሰከንድ በፊት|ከ\d+[.,]?\d* ሰዓቶች በፊት|ከ\d+[.,]?\d* ሳምንት በፊት|ከ\d+[.,]?\d* ዓመታት በፊት|በ\d+[.,]?\d* ሰዓት ውስጥ|በ\d+[.,]?\d* ቀናት ውስጥ|በ\d+[.,]?\d* ቀኖች ውስጥ|በ\d+[.,]?\d* ወራት ውስጥ|በ\d+[.,]?\d* ደቂቃ ውስጥ|ከ \d+[.,]?\d* ቀን በፊት|ከ\d+[.,]?\d* ሰዓት በፊት|ከ\d+[.,]?\d* ቀናት በፊት|ከ\d+[.,]?\d* ቀኖች በፊት|ከ\d+[.,]?\d* ወራት በፊት|ከ\d+[.,]?\d* ዓመት በፊት|ከ\d+[.,]?\d* ደቂቃ በፊት|በ\d+[.,]?\d* ቀን ውስጥ|በ\d+[.,]?\d* ወር ውስጥ|ከ\d+[.,]?\d* ቀን በፊት|ከ\d+[.,]?\d* ወር በፊት)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(በ\d+[.,]?\d* ሰከንዶች ውስጥ|በ\d+[.,]?\d* ሳምንታት ውስጥ|በ\d+[.,]?\d* ደቂቃዎች ውስጥ|ከ\d+[.,]?\d* ሰከንዶች በፊት|ከ\d+[.,]?\d* ሳምንታት በፊት|ከ\d+[.,]?\d* ደቂቃዎች በፊት|በ\d+[.,]?\d* ሰከንድ ውስጥ|በ\d+[.,]?\d* ሰዓቶች ውስጥ|በ\d+[.,]?\d* ሳምንት ውስጥ|በ\d+[.,]?\d* ዓመታት ውስጥ|ከ\d+[.,]?\d* ሰከንድ በፊት|ከ\d+[.,]?\d* ሰዓቶች በፊት|ከ\d+[.,]?\d* ሳምንት በፊት|ከ\d+[.,]?\d* ዓመታት በፊት|በ\d+[.,]?\d* ሰዓት ውስጥ|በ\d+[.,]?\d* ቀናት ውስጥ|በ\d+[.,]?\d* ቀኖች ውስጥ|በ\d+[.,]?\d* ወራት ውስጥ|በ\d+[.,]?\d* ደቂቃ ውስጥ|ከ \d+[.,]?\d* ቀን በፊት|ከ\d+[.,]?\d* ሰዓት በፊት|ከ\d+[.,]?\d* ቀናት በፊት|ከ\d+[.,]?\d* ቀኖች በፊት|ከ\d+[.,]?\d* ወራት በፊት|ከ\d+[.,]?\d* ዓመት በፊት|ከ\d+[.,]?\d* ደቂቃ በፊት|በ\d+[.,]?\d* ቀን ውስጥ|በ\d+[.,]?\d* ወር ውስጥ|ከ\d+[.,]?\d* ቀን በፊት|ከ\d+[.,]?\d* ወር በፊት)$`),
		KnownWords:      []string{"የሚቀጥለው ሳምንት", "የሚቀጥለው ዓመት", "ባለፈው ሳምንት", "የሚቀጥለው ወር", "ያለፈው ሳምንት", "በዚህ ሣምንት", "በዚህ ሳምንት", "ያለፈው ዓመት", "በዚህ ዓመት", "ያለፈው ወር", "ሴፕቴምበር", "በዚህ ወር", "ይህ ሰዓት", "ይህ ደቂቃ", "ትላንትና", "ኖቬምበር", "ኦክቶበር", "ዲሴምበር", "ጃንዩወሪ", "ፌብሩወሪ", "ማክሰኞ", "ሰከንድ", "ሳምንት", "ትናንት", "ኤፕሪል", "ኦገስት", "ከሰዓት", "gmt", "utc", "ሐሙስ", "ማርች", "ማክሰ", "ረቡዕ", "ሰዓት", "ሴፕቴ", "ቅዳሜ", "ኖቬም", "አሁን", "ኤፕሪ", "እሑድ", "ኦክቶ", "ኦገስ", "ዓመት", "ዓርብ", "ደቂቃ", "ዲሴም", "ጁላይ", "ጃንዩ", "ጥዋት", "ፌብሩ", "am", "pm", "ሜይ", "ሰኞ", "ቀን", "ነገ", "ወር", "ዛሬ", "ጁን", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
