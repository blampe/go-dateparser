// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ar_Locale = LocaleData{
	Name:      "ar",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "الساعة", "في", "مساءً", "هـ"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) أسبوع(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) ثانية(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) دقيقة(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) ساعة(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)قبل (\d+) أسبوع(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)قبل (\d+) ثانية(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)الأسبوع القادم(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)الأسبوع الماضي(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)الساعة الحالية(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) سنة(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) شهر(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)خلال (\d+) يوم(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)قبل (\d+) ساعة(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)السنة الحالية(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)السنة القادمة(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)السنة الماضية(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)قبل (\d+) سنة(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)قبل (\d+) شهر(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)قبل (\d+) يوم(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)الشهر القادم(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)الشهر الماضي(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)اليوم السابق(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)هذا الأسبوع(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)هذه الدقيقة(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ساعة واحدة(\z|\W|_)`), "${1}1 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)هذا الشهر(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)الأربعاء(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)الثلاثاء(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)الأسبوع(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)الإثنين(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)الاثنين(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)الثواني(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)الساعات(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)أكتوبر(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)الجمعة(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)الخميس(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ديسمبر(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ساعتين(\z|\W|_)`), "${1}2 hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)سبتمبر(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)صباحاً(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)صباحًا(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)فبراير(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)نوفمبر(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)أبريل(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)أسبوع(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)أغسطس(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)الأحد(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)الأمس(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)السبت(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)السنة(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)الشهر(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)اليوم(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ثانية(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)دقائق(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)دقيقة(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)ساعات(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)مساءً(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)يناير(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)يوليو(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)يومين(\z|\W|_)`), "${1}2 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)يونيو(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيام(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)الآن(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)خلال(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)ساعة(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)غدًا(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)مارس(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)مايو(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)أمس(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)سنة(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)شهر(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)عام(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)منذ(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)يوم(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ص(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)م(\z|\W|_)`), "${1}pm${2}"},
	},
}

var ar_AE_Locale = LocaleData{
	Name:      "ar-AE",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)السنة التالية(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)هذه السنة(\z|\W|_)`), "${1}0 year ago${2}"},
	},
}

var ar_BH_Locale = LocaleData{
	Name:      "ar-BH",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_DJ_Locale = LocaleData{
	Name:      "ar-DJ",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_DZ_Locale = LocaleData{
	Name:      "ar-DZ",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)جويلية(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)أفريل(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)جانفي(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)فيفري(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)جوان(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أوت(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماي(\z|\W|_)`), "${1}may${2}"},
	},
}

var ar_EG_Locale = LocaleData{
	Name:      "ar-EG",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_EH_Locale = LocaleData{
	Name:      "ar-EH",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_ER_Locale = LocaleData{
	Name:      "ar-ER",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_IL_Locale = LocaleData{
	Name:      "ar-IL",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_IQ_Locale = LocaleData{
	Name:      "ar-IQ",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)تشرين الثاني(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الثاني(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)تشرين الأول(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)تشرین الأول(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الأول(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)حزيران(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيلول(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)نيسان(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)آذار(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيار(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)تموز(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)شباط(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)آب(\z|\W|_)`), "${1}august${2}"},
	},
}

var ar_JO_Locale = LocaleData{
	Name:      "ar-JO",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)تشرين الثاني(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الثاني(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)تشرين الأول(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الأول(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)حزيران(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيلول(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)نيسان(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)آذار(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيار(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)تموز(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)شباط(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)آب(\z|\W|_)`), "${1}august${2}"},
	},
}

var ar_KM_Locale = LocaleData{
	Name:      "ar-KM",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_KW_Locale = LocaleData{
	Name:      "ar-KW",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_LB_Locale = LocaleData{
	Name:      "ar-LB",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)تشرين الثاني(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الثاني(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)تشرين الأول(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الأول(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)حزيران(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيلول(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)نيسان(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)آذار(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيار(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)تموز(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)شباط(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)آب(\z|\W|_)`), "${1}august${2}"},
	},
}

var ar_LY_Locale = LocaleData{
	Name:      "ar-LY",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_MA_Locale = LocaleData{
	Name:      "ar-MA",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)يوليوز(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)دجنبر(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)شتنبر(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)نونبر(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)غشت(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماي(\z|\W|_)`), "${1}may${2}"},
	},
}

var ar_MR_Locale = LocaleData{
	Name:      "ar-MR",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)إبريل(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)دجمبر(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)شتمبر(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)أغشت(\z|\W|_)`), "${1}august${2}"},
	},
}

var ar_OM_Locale = LocaleData{
	Name:      "ar-OM",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_PS_Locale = LocaleData{
	Name:      "ar-PS",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)تشرين الثاني(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الثاني(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)تشرين الأول(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الأول(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)حزيران(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيلول(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)نيسان(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)آذار(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيار(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)تموز(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)شباط(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)آب(\z|\W|_)`), "${1}august${2}"},
	},
}

var ar_QA_Locale = LocaleData{
	Name:      "ar-QA",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_SA_Locale = LocaleData{
	Name:      "ar-SA",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_SD_Locale = LocaleData{
	Name:      "ar-SD",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_SO_Locale = LocaleData{
	Name:      "ar-SO",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_SS_Locale = LocaleData{
	Name:      "ar-SS",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_SY_Locale = LocaleData{
	Name:      "ar-SY",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)تشرين الثاني(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الثاني(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)تشرين الأول(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)كانون الأول(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)حزيران(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيلول(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)نيسان(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)آذار(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)أيار(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)تموز(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)شباط(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)آب(\z|\W|_)`), "${1}august${2}"},
	},
}

var ar_TD_Locale = LocaleData{
	Name:      "ar-TD",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}

var ar_TN_Locale = LocaleData{
	Name:      "ar-TN",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)جويلية(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)أفريل(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)جانفي(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)فيفري(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)جوان(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)أوت(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماي(\z|\W|_)`), "${1}may${2}"},
	},
}

var ar_YE_Locale = LocaleData{
	Name:      "ar-YE",
	Parent:    &ar_Locale,
	DateOrder: "DMY",
}
