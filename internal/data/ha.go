// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ha_Locale = LocaleData{
	Name:      "ha",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)faburairu(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)afirilu(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)alhamis(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)disamba(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)janairu(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumma'a(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)litinin(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)nuwamba(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)satumba(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)shekara(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)agusta(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)asabar(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)daƙiƙa(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)lahadi(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)laraba(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktoba(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)talata(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)kwana(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)maris(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)gobe(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)jiya(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mako(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)mayu(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)wata(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)yuli(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)yuni(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)afi(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)agu(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)alh(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)asa(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)awa(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)dis(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)fab(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jum(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lah(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lar(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lit(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)may(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)nuw(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)sat(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)tal(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)yau(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)yul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)yun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
	},
}

var ha_GH_Locale = LocaleData{
	Name:      "ha-GH",
	Parent:    &ha_Locale,
	DateOrder: "DMY",
}

var ha_NE_Locale = LocaleData{
	Name:      "ha-NE",
	Parent:    &ha_Locale,
	DateOrder: "DMY",
}
