// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var gu_Locale = LocaleData{
	Name:      "gu",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) અઠવાડિયા પહેલાં(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) મહિના પહેલાં(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) સેકંડ પહેલાં(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) અઠવાડિયામાં(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) કલાક પહેલાં(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) દિવસ પહેલાં(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) વર્ષ પહેલાં(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) વર્ષ પહેલા(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) અઠ પહેલાં(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) મહિનામાં(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) મિનિટમાં(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) સેકંડમાં(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) કલાકમાં(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) દિવસમાં(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) વર્ષમાં(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)આવતા અઠવાડિયે(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) અઠ માં(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)ગયા અઠવાડિયે(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)આ અઠવાડિયે(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)આવતા મહિને(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)આવતા વર્ષે(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)અઠવાડિયું(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ગયા મહિને(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ગયા વર્ષે(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)જાન્યુઆરી(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ફેબ્રુઆરી(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)સપ્ટેમ્બર(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)આવતીકાલે(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ડિસેમ્બર(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)શુક્રવાર(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)આ મહિને(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)આ મિનિટ(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)આ વર્ષે(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ઑક્ટોબર(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ગુરુવાર(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)નવેમ્બર(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)મંગળવાર(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)આ કલાક(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)એપ્રિલ(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ગઈકાલે(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)જાન્યુ(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ફેબ્રુ(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)બુધવાર(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)રવિવાર(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)શનિવાર(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)સેકન્ડ(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)સોમવાર(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ઑક્ટો(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ઑગસ્ટ(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)જુલાઈ(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)મહિનો(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)માર્ચ(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)શુક્ર(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)સપ્ટે(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)હમણાં(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)કલાક(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)ગુરુ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ડિસે(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)દિવસ(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)મંગળ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)વર્ષ(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)આજે(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)જૂન(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)નવે(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)બુધ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)રવિ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)શનિ(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)સોમ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)અઠ(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)મે(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)સે(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ક(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)મ(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)વ(\z|\W|_)`), "${1}year${2}"},
	},
}
