// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var hi_Locale = LocaleData{
	Name:      "hi",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "के", "को", "बजे", "सन्", "से"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) सप्ताह पहले(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) सप्ताह में(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) सेकंड पहले(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) घंटे पहले(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) वर्ष पहले(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) सेकंड में(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) घंटे में(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) दिन पहले(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) माह पहले(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) मिनट में(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) वर्ष में(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) घं पहले(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) दिन में(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) माह में(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) से पहले(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) घं में(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) मि में(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) से में(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)पिछला सप्ताह(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)अगला सप्ताह(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)पिछला वर्ष(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)अगला वर्ष(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)इस सप्ताह(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)पिछला माह(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)पूर्वाह्न(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)अगला माह(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)शुक्रवार(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)अक्टूबर(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)अक्तूबर(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)अपराह्न(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)इस वर्ष(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)गुरुवार(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)दिसम्बर(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)मंगलवार(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)यह घंटा(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)यह मिनट(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)सितम्बर(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)अक्तू॰(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)अप्रैल(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)इस माह(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)दिसंबर(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)नवम्बर(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)फ़रवरी(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)बुधवार(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)रविवार(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)वर्षों(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)शनिवार(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)सप्ताह(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)सितंबर(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)सोमवार(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)अगस्त(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)जनवरी(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)जुलाई(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)नवंबर(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)परसों(\z|\W|_)`), "${1}2 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)पूर्व(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)महीना(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)महीने(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)मार्च(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)शुक्र(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)सेकंड(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)गुरु(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)घंटा(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)घंटे(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)जुल॰(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)दिवस(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)दिस॰(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)पहले(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)फ़र॰(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)मंगल(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)वर्ष(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)सित॰(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)अग॰(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)जन॰(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)जून(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)दिन(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)नव॰(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)बाद(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)बुध(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)मास(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)माह(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)में(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)रवि(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)शनि(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)साल(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)सोम(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)अब(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)आज(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)कल(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)घं(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)मई(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)से(\z|\W|_)`), "${1}second${2}"},
	},
}
