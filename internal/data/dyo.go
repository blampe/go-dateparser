// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var dyo_Locale = merge(nil, LocaleData{
	Name:      "dyo",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvwxyzŋ`),
	Translations: map[string]string{
		"settembar": "september",
		"disambar":  "december",
		"novembar":  "november",
		"aramisay":  "thursday",
		"febirie":   "february",
		"oktobar":   "october",
		"alarbay":   "wednesday",
		"aburil":    "april",
		"arjuma":    "friday",
		"sanvie":    "january",
		"suuyee":    "july",
		"minute":    "minute",
		"fuleeŋ":    "month",
		"sibiti":    "saturday",
		"second":    "second",
		"talata":    "tuesday",
		"lookuŋ":    "week",
		"funak":     "day",
		"teneŋ":     "monday",
		"dimas":     "sunday",
		"hour":      "hour",
		"sueŋ":      "june",
		"mars":      "march",
		"emit":      "year",
		"arj":       "friday",
		"gmt":       "gmt",
		"mee":       "may",
		"ten":       "monday",
		"sib":       "saturday",
		"dim":       "sunday",
		"ara":       "thursday",
		"tal":       "tuesday",
		"utc":       "utc",
		"ala":       "wednesday",
		"am":        "am",
		"ab":        "april",
		"ut":        "august",
		"de":        "december",
		"fe":        "february",
		"sa":        "january",
		"su":        "july",
		"ma":        "march",
		"me":        "may",
		"no":        "november",
		"ok":        "october",
		"pm":        "pm",
		"se":        "september",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"fucen":       "1 day ago",
		"kajom":       "in 1 day",
		"jaat":        "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|settembar|this hour|this week|this year|aramisay|disambar|novembar|alarbay|febirie|oktobar|aburil|arjuma|fuleeŋ|lookuŋ|minute|sanvie|second|sibiti|suuyee|talata|dimas|fucen|funak|kajom|teneŋ|emit|hour|jaat|mars|sueŋ|ala|ara|arj|dim|gmt|mee|now|sib|tal|ten|utc|\+|\.|\[|\]|\||ab|am|de|fe|ma|me|no|ok|pm|sa|se|su|ut| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
