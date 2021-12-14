// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kk_Locale = LocaleData{
	Name:      "kk",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) секундтан кейін(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) минуттан кейін(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сағаттан кейін(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) аптадан кейін(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) жылдан кейін(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) күннен кейін(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) секунд бұрын(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) айдан кейін(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сағат бұрын(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) апта бұрын(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) жыл бұрын(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) күн бұрын(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) мин кейін(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сағ бұрын(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сағ кейін(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сек бұрын(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сек кейін(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ай бұрын(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ап бұрын(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ап кейін(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ж бұрын(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ж кейін(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)былтырғы жыл(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)келесі апта(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)биылғы жыл(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)келесі жыл(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)өткен апта(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)желтоқсан(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)келесі ай(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)осы минут(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)осы сағат(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)бейсенбі(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)дүйсенбі(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)жексенбі(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)осы апта(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)сейсенбі(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сәрсенбі(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)қыркүйек(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)өткен ай(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)маусым(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)наурыз(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)осы ай(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунд(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)қараша(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)қаңтар(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ақпан(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)бүгін(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ертең(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)мамыр(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)сағат(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)сенбі(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сәуір(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)тамыз(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)шілде(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)қазан(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)қазір(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)апта(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)жұма(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)кеше(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ақп(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)жел(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)жыл(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)күн(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)мам(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)мау(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)нау(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)сағ(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)сәу(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)там(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)шіл(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)қаз(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)қар(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)қаң(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)қыр(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ай(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ап(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)бс(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)дс(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)жм(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)жс(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сб(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ср(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сс(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ж(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)с(\z|\W|_)`), "${1}second${2}"},
	},
}
