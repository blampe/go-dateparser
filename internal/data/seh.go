// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var seh_Locale = LocaleData{
	Name:      "seh",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
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
		{regexp.MustCompile(`(\A|\W|_)chishanu(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)decembro(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)fevreiro(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)manguana(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)novembro(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)setembro(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)augusto(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)chipiri(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)chiposi(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)chitatu(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dimingu(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)janeiro(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)segundo(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)chinai(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ntsiku(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)otubro(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)sabudu(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)abril(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)chaka(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)julho(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)junho(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)marco(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwezi(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)hora(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)lero(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)maio(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)week(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)zuro(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)abr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)aug(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)dim(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)fev(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mai(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)nai(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)otu(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)pir(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)pos(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sab(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)set(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)sha(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tat(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
	},
}
