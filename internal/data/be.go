// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var be_Locale = LocaleData{
	Name:      "be",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "каля", "у", "і", "ў"},
	Simplifications: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)(\d+)\s*гадзін\s(\d+)\s*хвілін(\z|\W|_)`): "${1}$2:$3${4}",
		regexp.MustCompile(`(\A|\W|_)^гадзіна(\z|\W|_)`):                       "${1}1 гадзіна${2}",
		regexp.MustCompile(`(\A|\W|_)^секунду(\z|\W|_)`):                       "${1}1 секунду${2}",
		regexp.MustCompile(`(\A|\W|_)^хвіліну(\z|\W|_)`):                       "${1}1 хвіліну${2}",
		regexp.MustCompile(`(\A|\W|_)гадзіну(\z|\W|_)`):                        "${1}1 гадзіну${2}",
		regexp.MustCompile(`(\A|\W|_)некалькі секунд(\z|\W|_)`):                "${1}44 секунды${2}",
		regexp.MustCompile(`(\A|\W|_)некалькі хвілін(\z|\W|_)`):                "${1}2 хвіліны${2}",
	},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) гадзіну таму(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) гадзіны таму(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) секунду таму(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) секунды таму(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) тыдзень таму(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)на наступным тыдні(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) гадзіну(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) гадзіны(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) секунду(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) секунды(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) тыдзень(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) хвіліну(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) хвіліны(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)у наступным месяцы(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) месяца таму(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) месяца(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)у наступным годзе(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) дзень таму(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) месяц таму(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) тыдня таму(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)на мінулым тыдні(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) дзень(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) месяц(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) тыдня(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)у мінулым месяцы(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) гадз таму(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) года таму(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) гадз(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) года(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)у мінулым годзе(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) год таму(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) дня таму(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) мес таму(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) тыд таму(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)на гэтым тыдні(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) год(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) дня(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) мес(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) тыд(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)у гэту гадзіну(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)у гэту хвіліну(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)у гэтым месяцы(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) хв(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)у гэтым годзе(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) г таму(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) д таму(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) с таму(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) г(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) д(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)праз (\d+) с(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)кастрычніка(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)кастрычнік(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)на працягу(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)панядзелак(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)таму назад(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)красавіка(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)лістапада(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)пазаўчора(\z|\W|_)`), "${1}2 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)верасень(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)красавік(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)лістапад(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)сакавіка(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)студзень(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)студзеня(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)аўторак(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)верасня(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)гадзіна(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)гадзіну(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)гадзіны(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)жнівень(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)жнівеня(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)месяцаў(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)нядзеля(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пятніца(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сакавік(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунда(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунду(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунды(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)снежань(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)тыдзень(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)хвіліну(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)хвіліны(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)чэрвень(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)чэрвеня(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)гадзін(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)жніўня(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)заўтра(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)лютага(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ліпень(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ліпеня(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)месяца(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)месяцы(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунд(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)серада(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)снежня(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)субота(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)траўня(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)тыдняў(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)хвілін(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)чацвер(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)гадоў(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)дзень(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)месяц(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)назад(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)сення(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)сёння(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)тыдня(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)тыдні(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)учора(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)цяпер(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ўчора(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)гадз(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)гады(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)года(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)дзен(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)дзён(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)люты(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)таму(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)хвіл(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)аўт(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)вер(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)врс(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)год(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)дні(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)жнв(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)жні(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)кас(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)кра(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)крс(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)кст(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)лют(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ліп(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ліс(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)май(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)мая(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)мес(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)няд(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пнд(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пят(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сак(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)сек(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)сне(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)снж(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)стд(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)сту(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)суб(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)тра(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)тыд(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)чцв(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)чэр(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)аў(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)нд(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пн(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пт(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сб(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ср(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)чв(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)чц(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)г(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)д(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)с(\z|\W|_)`), "${1}second${2}"},
	},
}
