// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ta_Locale = LocaleData{
	Name:      "ta",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) வாரத்திற்கு முன்பு(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) விநாடிகளுக்கு முன்(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ஆண்டுகளுக்கு முன்(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மாதங்களுக்கு முன்(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வாரங்களுக்கு முன்(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நாட்களுக்கு முன்(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மாதத்துக்கு முன்(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) விநாடிக்கு முன்(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)இந்த ஒரு மணிநேரத்தில்(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ஆண்டிற்கு முன்(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)இந்த ஒரு நிமிடத்தில்(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நாளுக்கு முன்(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மணிநேரம் முன்(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நிமிடங்களில்(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மணிநேரத்தில்(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நிமிடத்தில்(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) விநாடிகளில்(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ஆண்டுகளில்(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மாதங்களில்(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வாரங்களில்(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) விநாடியில்(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நாட்களில்(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மாதத்தில்(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வாரத்தில்(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) விநா முன்(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மணி முன்(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மாத முன்(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வார முன்(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ஆண்டில்(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நா முன்(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மா முன்(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வா முன்(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வி முன்(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ஆ முன்(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நாளில்(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ம முன்(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)அடுத்த ஆண்டு(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)அடுத்த மாதம்(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)அடுத்த வாரம்(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)கடந்த ஆண்டு(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)கடந்த மாதம்(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)கடந்த வாரம்(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நிமி(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) விநா(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)இந்த ஆண்டு(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)இந்த மாதம்(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)இந்த வாரம்(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)செப்டம்பர்(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மணி(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மாத(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வார(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நா(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) நி(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) மா(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வா(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) வி(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)அக்டோபர்(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)செவ்வாய்(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)டிசம்பர்(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)பிப்ரவரி(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)பிற்பகல்(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)முற்பகல்(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ஆ(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ம(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)இப்போது(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)திங்கள்(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)நவம்பர்(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)வியாழன்(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஆகஸ்ட்(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஏப்ரல்(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஞாயிறு(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)நேற்று(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)மார்ச்(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)விநாடி(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)வெள்ளி(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஆண்டு(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)இன்று(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஜனவரி(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)புதன்(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)மாதம்(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)வாரம்(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)செப்(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)செவ்(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஜூன்(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஜூலை(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஞாயி(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)திங்(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)நாளை(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)நாள்(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)பிப்(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)மார்(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)விநா(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)வியா(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)வெள்(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)அக்(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஏப்(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)சனி(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)டிச(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)புத(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)மணி(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)மாத(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஆக(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஜன(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)நவ(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)நா(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)மா(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)மே(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)வா(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)வி(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ஆ(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ம(\z|\W|_)`), "${1}hour${2}"},
	},
}

var ta_LK_Locale = LocaleData{
	Name:      "ta-LK",
	Parent:    &ta_Locale,
	DateOrder: "DMY",
}

var ta_MY_Locale = LocaleData{
	Name:      "ta-MY",
	Parent:    &ta_Locale,
	DateOrder: "DMY",
}

var ta_SG_Locale = LocaleData{
	Name:      "ta-SG",
	Parent:    &ta_Locale,
	DateOrder: "DMY",
}
