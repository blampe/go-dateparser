// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ii_Locale = LocaleData{
	Name:      "ii",
	DateOrder: "YMD",
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
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꀋꅔꉈ(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꃆꏂꑍ(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꊂꃘ(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꊂꇖ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꊂꉬ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꊂꋍ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꊂꌕ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꊂꑍ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꊰꊪꆪ(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꊰꑋꆪ(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꑭꆏꑍ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꀃꑍ(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꁯꋒ(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꃘꆪ(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꄮꈉ(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꃘ(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꇖ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꉬ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꋍ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꌕ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆏꑍ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꇖꆪ(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꈬꆪ(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꉆꆪ(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꉬꆪ(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꊰꆪ(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꋍꆪ(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꌕꆪ(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꎸꄑ(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꏃꆪ(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꑍꆪ(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꑭꆏ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꆪ(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꇙ(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꈎ(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ꑍ(\z|\W|_)`), "${1}day${2}"},
	},
}
