// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ksf_Locale = LocaleData{
	Name:      "ksf",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ ntɛk di bɔk(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ ntɛk di bɛɛ(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ táabɛɛ(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ táafɔk(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ táanin(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ táaraa(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí a ntɔntɔ(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ ntɛk(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ táan(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ bɛɛ(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ nin(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí akǝ ráá(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mɛkrɛdí(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)cámɛɛn(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)cɛɛnko(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ridúrǝ(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)rinkɔɔ(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sárúwá(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)gɛɛnǝ(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)júmbá(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jǝǝdí(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lǝndí(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)maadí(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)samdí(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sɔndǝ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwíí(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋwós(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)bǝk(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)háu(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)júm(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jǝǝ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lǝn(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)maa(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mɛk(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sam(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sɔn(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ10(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ11(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ12(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ1(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ2(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ3(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ4(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ5(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ6(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ7(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ8(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋ9(\z|\W|_)`), "${1}september${2}"},
	},
}
