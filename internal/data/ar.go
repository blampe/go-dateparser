// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ar_Locale = merge(nil, LocaleData{
	Name:      "ar",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "الساعة", "في", "مساء", "هـ"},
	Translations: map[string]string{
		"الثلاثاء": "tuesday",
		"الاربعاء": "wednesday",
		"الساعات":  "hour",
		"الدقايق":  "minute",
		"الاثنين":  "monday",
		"الثواني":  "second",
		"الاسبوع":  "week",
		"الساعة":   "",
		"ديسمبر":   "december",
		"فبراير":   "february",
		"الجمعة":   "friday",
		"نوفمبر":   "november",
		"اكتوبر":   "october",
		"سبتمبر":   "september",
		"الخميس":   "thursday",
		"صباحا":    "am",
		"ابريل":    "april",
		"اغسطس":    "august",
		"ساعات":    "hour",
		"يناير":    "january",
		"يوليو":    "july",
		"يونيو":    "june",
		"دقايق":    "minute",
		"دقيقة":    "minute",
		"الشهر":    "month",
		"السبت":    "saturday",
		"ثانية":    "second",
		"الاحد":    "sunday",
		"اسبوع":    "week",
		"السنة":    "year",
		"ايام":     "day",
		"ساعة":     "hour",
		"خلال":     "in",
		"مارس":     "march",
		"مايو":     "may",
		"مساء":     "pm",
		"منذ":      "ago",
		"يوم":      "day",
		"gmt":      "gmt",
		"شهر":      "month",
		"utc":      "utc",
		"سنة":      "year",
		"عام":      "year",
		"في":       "",
		"هـ":       "",
		"am":       "am",
		"pm":       "pm",
		"'":        "",
		",":        "",
		";":        "",
		"@":        "",
		"[":        "",
		"]":        "",
		"|":        "",
		" ":        " ",
		"+":        "+",
		"-":        "-",
		".":        ".",
		"/":        "/",
		":":        ":",
		"ص":        "am",
		"م":        "pm",
		"z":        "z",
	},
	RelativeType: map[string]string{
		"الساعة الحالية": "0 hour ago",
		"الاسبوع الماضي": "1 week ago",
		"الاسبوع القادم": "in 1 week",
		"السنة الحالية":  "0 year ago",
		"السنة الماضية":  "1 year ago",
		"السنة القادمة":  "in 1 year",
		"اليوم السابق":   "1 day ago",
		"الشهر الماضي":   "1 month ago",
		"الشهر القادم":   "in 1 month",
		"هذه الدقيقة":    "0 minute ago",
		"هذا الاسبوع":    "0 week ago",
		"ساعة واحدة":     "1 hour ago",
		"هذا الشهر":      "0 month ago",
		"ساعتين":         "2 hour",
		"اليوم":          "0 day ago",
		"الامس":          "1 day ago",
		"يومين":          "2 day",
		"الان":           "0 second ago",
		"امس":            "1 day ago",
		"غدا":            "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)خلال (\d+) دقيقة`), "in $1 minute"},
		{regexp.MustCompile(`(?i)خلال (\d+) ثانية`), "in $1 second"},
		{regexp.MustCompile(`(?i)خلال (\d+) اسبوع`), "in $1 week"},
		{regexp.MustCompile(`(?i)قبل (\d+) دقيقة`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)قبل (\d+) ثانية`), "$1 second ago"},
		{regexp.MustCompile(`(?i)قبل (\d+) اسبوع`), "$1 week ago"},
		{regexp.MustCompile(`(?i)خلال (\d+) ساعة`), "in $1 hour"},
		{regexp.MustCompile(`(?i)قبل (\d+) ساعة`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)خلال (\d+) يوم`), "in $1 day"},
		{regexp.MustCompile(`(?i)خلال (\d+) شهر`), "in $1 month"},
		{regexp.MustCompile(`(?i)خلال (\d+) سنة`), "in $1 year"},
		{regexp.MustCompile(`(?i)قبل (\d+) يوم`), "$1 day ago"},
		{regexp.MustCompile(`(?i)قبل (\d+) شهر`), "$1 month ago"},
		{regexp.MustCompile(`(?i)قبل (\d+) سنة`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(خلال \d+ اسبوع|خلال \d+ ثانية|خلال \d+ دقيقة|خلال \d+ ساعة|قبل \d+ اسبوع|قبل \d+ ثانية|قبل \d+ دقيقة|خلال \d+ سنة|خلال \d+ شهر|خلال \d+ يوم|قبل \d+ ساعة|قبل \d+ سنة|قبل \d+ شهر|قبل \d+ يوم)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(خلال \d+ اسبوع|خلال \d+ ثانية|خلال \d+ دقيقة|خلال \d+ ساعة|قبل \d+ اسبوع|قبل \d+ ثانية|قبل \d+ دقيقة|خلال \d+ سنة|خلال \d+ شهر|خلال \d+ يوم|قبل \d+ ساعة|قبل \d+ سنة|قبل \d+ شهر|قبل \d+ يوم)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(الاسبوع القادم|الاسبوع الماضي|الساعة الحالية|السنة الحالية|السنة القادمة|السنة الماضية|الشهر القادم|الشهر الماضي|اليوم السابق|هذا الاسبوع|هذه الدقيقة|ساعة واحدة|هذا الشهر|الاربعاء|الثلاثاء|الاثنين|الاسبوع|الثواني|الدقايق|الساعات|اكتوبر|الجمعة|الخميس|الساعة|ديسمبر|ساعتين|سبتمبر|فبراير|نوفمبر|ابريل|اسبوع|اغسطس|الاحد|الامس|السبت|السنة|الشهر|اليوم|ثانية|دقايق|دقيقة|ساعات|صباحا|يناير|يوليو|يومين|يونيو|الان|ايام|خلال|ساعة|مارس|مايو|مساء|gmt|utc|امس|سنة|شهر|عام|غدا|منذ|يوم|\+|\.|\[|\]|\||am|pm|في|هـ| |'|,|-|/|:|;|@|z|ص|م)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_AE_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-AE",
	DateOrder: "DMY",
	RelativeType: map[string]string{
		"السنة التالية": "in 1 year",
		"هذه السنة":     "0 year ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(السنة التالية|هذه السنة)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_BH_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-BH",
	DateOrder: "DMY",
})

var ar_DJ_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-DJ",
	DateOrder: "DMY",
})

var ar_DZ_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-DZ",
	DateOrder: "DMY",
	Translations: map[string]string{
		"جويلية": "july",
		"افريل":  "april",
		"فيفري":  "february",
		"جانفي":  "january",
		"جوان":   "june",
		"اوت":    "august",
		"gmt":    "gmt",
		"ماي":    "may",
		"utc":    "utc",
		"am":     "am",
		"pm":     "pm",
		" ":      " ",
		"+":      "+",
		"-":      "-",
		".":      ".",
		"/":      "/",
		":":      ":",
		"z":      "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(جويلية|افريل|جانفي|فيفري|جوان|gmt|utc|اوت|ماي|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_EG_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-EG",
	DateOrder: "DMY",
})

var ar_EH_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-EH",
	DateOrder: "DMY",
})

var ar_ER_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-ER",
	DateOrder: "DMY",
})

var ar_IL_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-IL",
	DateOrder: "DMY",
})

var ar_IQ_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-IQ",
	DateOrder: "DMY",
	Translations: map[string]string{
		"كانون الثاني": "january",
		"تشرين الثاني": "november",
		"كانون الاول":  "december",
		"تشرين الاول":  "october",
		"تشرین الاول":  "october",
		"حزيران":       "june",
		"نيسان":        "april",
		"ايلول":        "september",
		"شباط":         "february",
		"تموز":         "july",
		"اذار":         "march",
		"ايار":         "may",
		"gmt":          "gmt",
		"utc":          "utc",
		"am":           "am",
		"اب":           "august",
		"pm":           "pm",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(تشرين الثاني|كانون الثاني|تشرين الاول|تشرین الاول|كانون الاول|حزيران|ايلول|نيسان|اذار|ايار|تموز|شباط|gmt|utc|\+|\.|am|pm|اب| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_JO_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-JO",
	DateOrder: "DMY",
	Translations: map[string]string{
		"كانون الثاني": "january",
		"تشرين الثاني": "november",
		"كانون الاول":  "december",
		"تشرين الاول":  "october",
		"حزيران":       "june",
		"نيسان":        "april",
		"ايلول":        "september",
		"شباط":         "february",
		"تموز":         "july",
		"اذار":         "march",
		"ايار":         "may",
		"gmt":          "gmt",
		"utc":          "utc",
		"am":           "am",
		"اب":           "august",
		"pm":           "pm",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(تشرين الثاني|كانون الثاني|تشرين الاول|كانون الاول|حزيران|ايلول|نيسان|اذار|ايار|تموز|شباط|gmt|utc|\+|\.|am|pm|اب| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_KM_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-KM",
	DateOrder: "DMY",
})

var ar_KW_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-KW",
	DateOrder: "DMY",
})

var ar_LB_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-LB",
	DateOrder: "DMY",
	Translations: map[string]string{
		"كانون الثاني": "january",
		"تشرين الثاني": "november",
		"كانون الاول":  "december",
		"تشرين الاول":  "october",
		"حزيران":       "june",
		"نيسان":        "april",
		"ايلول":        "september",
		"شباط":         "february",
		"تموز":         "july",
		"اذار":         "march",
		"ايار":         "may",
		"gmt":          "gmt",
		"utc":          "utc",
		"am":           "am",
		"اب":           "august",
		"pm":           "pm",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(تشرين الثاني|كانون الثاني|تشرين الاول|كانون الاول|حزيران|ايلول|نيسان|اذار|ايار|تموز|شباط|gmt|utc|\+|\.|am|pm|اب| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_LY_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-LY",
	DateOrder: "DMY",
})

var ar_MA_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-MA",
	DateOrder: "DMY",
	Translations: map[string]string{
		"يوليوز": "july",
		"دجنبر":  "december",
		"نونبر":  "november",
		"شتنبر":  "september",
		"غشت":    "august",
		"gmt":    "gmt",
		"ماي":    "may",
		"utc":    "utc",
		"am":     "am",
		"pm":     "pm",
		" ":      " ",
		"+":      "+",
		"-":      "-",
		".":      ".",
		"/":      "/",
		":":      ":",
		"z":      "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(يوليوز|دجنبر|شتنبر|نونبر|gmt|utc|غشت|ماي|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_MR_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-MR",
	DateOrder: "DMY",
	Translations: map[string]string{
		"دجمبر": "december",
		"شتمبر": "september",
		"اغشت":  "august",
		"gmt":   "gmt",
		"utc":   "utc",
		"am":    "am",
		"pm":    "pm",
		" ":     " ",
		"+":     "+",
		"-":     "-",
		".":     ".",
		"/":     "/",
		":":     ":",
		"z":     "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(دجمبر|شتمبر|اغشت|gmt|utc|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_OM_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-OM",
	DateOrder: "DMY",
})

var ar_PS_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-PS",
	DateOrder: "DMY",
	Translations: map[string]string{
		"كانون الثاني": "january",
		"تشرين الثاني": "november",
		"كانون الاول":  "december",
		"تشرين الاول":  "october",
		"حزيران":       "june",
		"نيسان":        "april",
		"ايلول":        "september",
		"شباط":         "february",
		"تموز":         "july",
		"اذار":         "march",
		"ايار":         "may",
		"gmt":          "gmt",
		"utc":          "utc",
		"am":           "am",
		"اب":           "august",
		"pm":           "pm",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(تشرين الثاني|كانون الثاني|تشرين الاول|كانون الاول|حزيران|ايلول|نيسان|اذار|ايار|تموز|شباط|gmt|utc|\+|\.|am|pm|اب| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_QA_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-QA",
	DateOrder: "DMY",
})

var ar_SA_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-SA",
	DateOrder: "DMY",
})

var ar_SD_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-SD",
	DateOrder: "DMY",
})

var ar_SO_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-SO",
	DateOrder: "DMY",
})

var ar_SS_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-SS",
	DateOrder: "DMY",
})

var ar_SY_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-SY",
	DateOrder: "DMY",
	Translations: map[string]string{
		"كانون الثاني": "january",
		"تشرين الثاني": "november",
		"كانون الاول":  "december",
		"تشرين الاول":  "october",
		"حزيران":       "june",
		"نيسان":        "april",
		"ايلول":        "september",
		"شباط":         "february",
		"تموز":         "july",
		"اذار":         "march",
		"ايار":         "may",
		"gmt":          "gmt",
		"utc":          "utc",
		"am":           "am",
		"اب":           "august",
		"pm":           "pm",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(تشرين الثاني|كانون الثاني|تشرين الاول|كانون الاول|حزيران|ايلول|نيسان|اذار|ايار|تموز|شباط|gmt|utc|\+|\.|am|pm|اب| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_TD_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-TD",
	DateOrder: "DMY",
})

var ar_TN_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-TN",
	DateOrder: "DMY",
	Translations: map[string]string{
		"جويلية": "july",
		"افريل":  "april",
		"فيفري":  "february",
		"جانفي":  "january",
		"جوان":   "june",
		"اوت":    "august",
		"gmt":    "gmt",
		"ماي":    "may",
		"utc":    "utc",
		"am":     "am",
		"pm":     "pm",
		" ":      " ",
		"+":      "+",
		"-":      "-",
		".":      ".",
		"/":      "/",
		":":      ":",
		"z":      "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(جويلية|افريل|جانفي|فيفري|جوان|gmt|utc|اوت|ماي|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ar_YE_Locale = merge(&ar_Locale, LocaleData{
	Name:      "ar-YE",
	DateOrder: "DMY",
})
