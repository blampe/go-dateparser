// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ksb_Locale = LocaleData{
	Name:      "ksb",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumaamosi(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumaatano(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumaatatu(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)alhamisi(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)febluali(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumaapii(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)nyiaghuo(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)septemba(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)desemba(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)januali(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumaane(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ng'waka(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)novemba(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekunde(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)agosti(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)aplili(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)evi eo(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ijumaa(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ng'ezi(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktoba(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)julai(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)keloi(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)machi(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)makeo(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)ghuo(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)juni(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)niki(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)siku(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ago(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)alh(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)des(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)iju(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jmn(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jmo(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jpi(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jtn(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jtt(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mac(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)mei(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)saa(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`), "${1}september${2}"},
	},
}
