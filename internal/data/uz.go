// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var uz_Locale = LocaleData{
	Name:      "uz",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) daqiqadan keyin(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) soniyadan keyin(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) haftadan keyin(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) soatdan keyin(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) kundan keyin(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) soniya oldin(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) yildan keyin(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hafta oldin(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) oydan keyin(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) soat oldin(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) kun oldin(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) yil oldin(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) oy oldin(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)keyingi hafta(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)o‘tgan hafta(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)shu daqiqada(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)keyingi yil(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)chorshanba(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)keyingi oy(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)o'tgan yil(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)o‘tgan yil(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)shu soatda(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)o‘tgan oy(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)payshanba(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)shu hafta(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)yakshanba(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dushanba(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)seshanba(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sentabr(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)shu yil(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)avgust(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu yil(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)dekabr(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ertaga(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)fevral(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)noyabr(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktabr(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)shanba(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)shu oy(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)soniya(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)yanvar(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)aprel(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)bugun(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)hafta(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)hozir(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)kecha(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)chor(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dush(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)juma(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mart(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)sesh(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)shan(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)soat(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)avg(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)dek(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)fev(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyl(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyn(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)jum(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)kun(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)may(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)noy(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)pay(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sen(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)son(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)yak(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)yan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)yil(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)oy(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)tk(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)to(\z|\W|_)`), "${1}am${2}"},
	},
}
