// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sah_Locale = LocaleData{
	Name:      "sah",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) сөкүүндэ ынараа өттүгэр(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) нэдиэлэ анараа өттүгэр(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) чаас ынараа өттүгэр(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) күн ынараа өттүгэр(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сыл ынараа өттүгэр(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сөк анараа өттүгэр(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ый ынараа өттүгэр(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сөкүүндэннэн(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) мүнүүтэннэн(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) нэдиэлэннэн(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) чааһынан(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)ааспыт нэдиэлэ(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) күнүнэн(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) сылынан(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)атырдьых ыйын(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)кэлэр нэдиэлэ(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ыйынан(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)атырдьых ыйа(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)балаҕан ыйын(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)баскыһыанньа(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)бэнидиэнньик(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)оптуорунньук(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)аныгыскы ый(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)балаҕан ыйа(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)кулун тутар(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)бу нэдиэлэ(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)муус устар(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ааспыт ый(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)тохсунньу(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)алтынньы(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ахсынньы(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)былырыын(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)бэс ыйын(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)бээтиҥсэ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сэтинньи(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)сөкүүндэ(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ыам ыйын(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)билигин(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)бэс ыйа(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)бэҕэһээ(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)нэдиэлэ(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)олунньу(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)от ыйын(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)субуота(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)чэппиэр(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ыам ыйа(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)от ыйа(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)сарсын(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)сэрэдэ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)бу ый(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)быйыл(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)бүгүн(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)эһиил(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)олун(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)тохс(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)чаас(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)алт(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)атр(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ахс(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)блҕ(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)бэс(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)клн(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)күн(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)мсу(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)отй(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)сыл(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)сэт(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ыам(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)бн(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)бс(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)бэ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)оп(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сб(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сэ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)чп(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ый(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)эи(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)эк(\z|\W|_)`), "${1}pm${2}"},
	},
}
