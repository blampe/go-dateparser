// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var tl_Locale = merge(nil, LocaleData{
	Name:      "tl",
	DateOrder: "",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "ganap", "na", "noon", "noong", "sa", "|"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)isang araw(\z|[^\pL\pM\d]|_)`), "${1}2 araw${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)kahapon(\z|[^\pL\pM\d]|_)`), "${1}1 araw nakaraan${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ngayon(\z|[^\pL\pM\d]|_)`), "${1}0 segundo nakalipas${2}"},
	},
	Translations: map[string]string{
		"miyerkules": "wednesday",
		"nakalipas":  "ago",
		"disyembre":  "december",
		"nobyembre":  "november",
		"setyembre":  "september",
		"nakaraan":   "ago",
		"biyernes":   "friday",
		"pebrero":    "february",
		"oktubre":    "october",
		"segundo":    "second",
		"huwebes":    "thursday",
		"agosto":     "august",
		"minuto":     "minute",
		"sabado":     "saturday",
		"linggo":     "sunday",
		"martes":     "tuesday",
		"ganap":      "",
		"noong":      "",
		"abril":      "april",
		"enero":      "january",
		"hulyo":      "july",
		"hunyo":      "june",
		"marso":      "march",
		"lunes":      "monday",
		"buwan":      "month",
		"noon":       "",
		"araw":       "day",
		"oras":       "hour",
		"mayo":       "may",
		"taon":       "year",
		"abr":        "april",
		"ago":        "august",
		"dis":        "december",
		"peb":        "february",
		"biy":        "friday",
		"gmt":        "gmt",
		"ene":        "january",
		"hul":        "july",
		"hun":        "june",
		"mar":        "march",
		"may":        "may",
		"lun":        "monday",
		"nob":        "november",
		"okt":        "october",
		"sab":        "saturday",
		"set":        "september",
		"lin":        "sunday",
		"huw":        "thursday",
		"utc":        "utc",
		"miy":        "wednesday",
		"na":         "",
		"am":         "am",
		"sa":         "in",
		"pm":         "pm",
		"'":          "",
		",":          "",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"|":          "",
		" ":          " ",
		"+":          "+",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		"z":          "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(miyerkules|disyembre|nakalipas|nobyembre|setyembre|biyernes|nakaraan|huwebes|oktubre|pebrero|segundo|agosto|linggo|martes|minuto|sabado|abril|buwan|enero|ganap|hulyo|hunyo|lunes|marso|noong|araw|mayo|noon|oras|taon|abr|ago|biy|dis|ene|gmt|hul|hun|huw|lin|lun|mar|may|miy|nob|okt|peb|sab|set|utc|\+|\.|\[|\]|\||am|na|pm|sa| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
