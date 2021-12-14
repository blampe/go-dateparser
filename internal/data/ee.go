// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ee_Locale = LocaleData{
	Name:      "ee",
	DateOrder: "MDY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)gaƒoƒo (\d+) si wo va yi(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa (\d+) si wo va yi(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)le aɖabaƒoƒo (\d+) wo me(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)sekend (\d+) si wo va yi(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)le ƒe (\d+) si gbɔna me(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)le ƒe (\d+) si va yi me(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ŋkeke (\d+) si wo va yi(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɣleti (\d+) si wo va yi(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)gaƒoƒo (\d+) si va yi(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa (\d+) si va yi(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)le aɖabaƒoƒo (\d+) me(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)le gaƒoƒo (\d+) wo me(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)le kɔsiɖa (\d+) wo me(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)le sekend (\d+) wo me(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)sekend (\d+) si va yi(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)le ŋkeke (\d+) wo me(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)le ɣleti (\d+) wo me(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)ŋkeke (\d+) si va yi(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe (\d+) si va yi me(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe (\d+) si wo va yi(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɣleti (\d+) si va yi(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)le gaƒoƒo (\d+) me(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)le kɔsiɖa (\d+) me(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)le sekend (\d+) me(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)le ŋkeke (\d+) me(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)le ɣleti (\d+) me(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe (\d+) si va yi(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa si gbɔ na(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa si va yi(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ɣleti si gbɔ na(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)le ƒe (\d+) me(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɣleti si va yi(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)etsɔ si gbɔna(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)etsɔ si va yi(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe si gbɔ na(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)adeɛmekpɔxe(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa ɖeka(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe si va yi(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)deasiamime(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa sia(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ɣleti sia(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)anyɔnyɔ(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)memleɖa(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)siamlɔm(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)afɔfĩe(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzodze(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)gaƒoƒo(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)kɔsiɖa(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekend(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)tedoxe(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)yawoɖa(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe sia(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)blaɖa(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzome(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzove(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzoɖa(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋkeke(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ɣetrɔ(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ɣleti(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)dama(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)egbe(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)fifi(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)fiɖa(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)kele(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)kuɖa(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)masa(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ade(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)afɔ(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)any(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)bla(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dam(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)dea(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzd(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzm(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzo(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dzv(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)fiɖ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)kel(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)kuɖ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)kɔs(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mas(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mem(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sia(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ted(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)yaw(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ŋdi(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)ƒe(\z|\W|_)`), "${1}year${2}"},
	},
}

var ee_TG_Locale = LocaleData{
	Name:      "ee-TG",
	Parent:    &ee_Locale,
	DateOrder: "MDY",
}
