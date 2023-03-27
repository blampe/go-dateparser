// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var vi_Locale LocaleData

func init() {
	vi_Locale = merge(nil, LocaleData{
		Name:                  "vi",
		DateOrder:             "DMY",
		Charset:               []rune(`bceghilnorstuyzàáâêìíôúăđưạảầậểồổộớờủứ`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(?:ngay|nam)\s(\d+[.,]?\d*)(\z|[^\pL\pM\d]|_)`), "${1}${2}${3}"},
		},
		Translations: map[string][]string{
			"thang muoi hai": {"december"},
			"thang muoi mot": {"november"},
			"giay đong ho":   {"second"},
			"nguyen ban":     {"minute"},
			"thang chin":     {"september"},
			"thang muoi":     {"october"},
			"thang bay":      {"july"},
			"thang hai":      {"february"},
			"thang mot":      {"january"},
			"thang nam":      {"may"},
			"thang sau":      {"june"},
			"thang tam":      {"august"},
			"truoc đay":      {"ago"},
			"ban ngay":       {"day"},
			"cach đay":       {"ago"},
			"chu nhat":       {"sunday"},
			"hang nhi":       {"second"},
			"thang 10":       {"october"},
			"thang 11":       {"november"},
			"thang 12":       {"december"},
			"thang ba":       {"march"},
			"thang tu":       {"april"},
			"thang 1":        {"january"},
			"thang 2":        {"february"},
			"thang 3":        {"march"},
			"thang 4":        {"april"},
			"thang 5":        {"may"},
			"thang 6":        {"june"},
			"thang 7":        {"july"},
			"thang 8":        {"august"},
			"thang 9":        {"september"},
			"thu bay":        {"saturday"},
			"thu hai":        {"monday"},
			"thu nam":        {"thursday"},
			"thu sau":        {"friday"},
			"tuan le":        {"week"},
			"thg 10":         {"october"},
			"thg 11":         {"november"},
			"thg 12":         {"december"},
			"thu ba":         {"tuesday"},
			"thu tu":         {"wednesday"},
			"thang":          {"month"},
			"thg 1":          {"january"},
			"thg 2":          {"february"},
			"thg 3":          {"march"},
			"thg 4":          {"april"},
			"thg 5":          {"may"},
			"thg 6":          {"june"},
			"thg 7":          {"july"},
			"thg 8":          {"august"},
			"thg 9":          {"september"},
			"thu 1":          {"sunday"},
			"thu 2":          {"monday"},
			"thu 3":          {"tuesday"},
			"thu 4":          {"wednesday"},
			"thu 5":          {"thursday"},
			"thu 6":          {"friday"},
			"thu 7":          {"saturday"},
			"trong":          {"in"},
			"truoc":          {"ago"},
			"buoi":           {"day"},
			"chut":           {"minute"},
			"giay":           {"second"},
			"ngay":           {"day"},
			"phut":           {"minute"},
			"th 2":           {"monday"},
			"th 3":           {"tuesday"},
			"th 4":           {"wednesday"},
			"th 5":           {"thursday"},
			"th 6":           {"friday"},
			"th 7":           {"saturday"},
			"tuan":           {"week"},
			"gio":            {"hour"},
			"gmt":            {"gmt"},
			"lat":            {"minute"},
			"luc":            {""},
			"nam":            {"year"},
			"thg":            {"month"},
			"utc":            {"utc"},
			"am":             {"am"},
			"ch":             {"pm"},
			"cn":             {"sunday"},
			"pm":             {"pm"},
			"sa":             {"am"},
			" ":              {" "},
			"'":              {""},
			"+":              {"+"},
			",":              {""},
			"-":              {"-"},
			".":              {"."},
			"/":              {"/"},
			":":              {":"},
			";":              {""},
			"@":              {""},
			"[":              {""},
			"]":              {""},
			"z":              {"z"},
			"|":              {""},
		},
		RelativeType: map[string]string{
			"thang truoc": "1 month ago",
			"tuan truoc":  "1 week ago",
			"nam ngoai":   "1 year ago",
			"thang nay":   "0 month ago",
			"thang sau":   "in 1 month",
			"ngay mai":    "in 1 day",
			"phut nay":    "0 minute ago",
			"tuan nay":    "0 week ago",
			"tuan sau":    "in 1 week",
			"bay gio":     "0 second ago",
			"gio nay":     "0 hour ago",
			"hom nay":     "0 day ago",
			"hom qua":     "1 day ago",
			"nam nay":     "0 year ago",
			"nam sau":     "in 1 year",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) thang nua`), "in $1 month"},
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) giay nua`), "in $1 second"},
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) ngay nua`), "in $1 day"},
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) phut nua`), "in $1 minute"},
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) tuan nua`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) thang truoc`), "$1 month ago"},
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) gio nua`), "in $1 hour"},
			{regexp.MustCompile(`(?i)sau (\d+[.,]?\d*) nam nua`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) giay truoc`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ngay truoc`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) phut truoc`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) tuan truoc`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) gio truoc`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) nam truoc`), "$1 year ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(sau \d+[.,]?\d* thang nua|sau \d+[.,]?\d* giay nua|sau \d+[.,]?\d* ngay nua|sau \d+[.,]?\d* phut nua|sau \d+[.,]?\d* tuan nua|\d+[.,]?\d* thang truoc|sau \d+[.,]?\d* gio nua|sau \d+[.,]?\d* nam nua|\d+[.,]?\d* giay truoc|\d+[.,]?\d* ngay truoc|\d+[.,]?\d* phut truoc|\d+[.,]?\d* tuan truoc|\d+[.,]?\d* gio truoc|\d+[.,]?\d* nam truoc)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(sau \d+[.,]?\d* thang nua|sau \d+[.,]?\d* giay nua|sau \d+[.,]?\d* ngay nua|sau \d+[.,]?\d* phut nua|sau \d+[.,]?\d* tuan nua|\d+[.,]?\d* thang truoc|sau \d+[.,]?\d* gio nua|sau \d+[.,]?\d* nam nua|\d+[.,]?\d* giay truoc|\d+[.,]?\d* ngay truoc|\d+[.,]?\d* phut truoc|\d+[.,]?\d* tuan truoc|\d+[.,]?\d* gio truoc|\d+[.,]?\d* nam truoc)$`),
		KnownWords:      []string{"thang muoi hai", "thang muoi mot", "giay đong ho", "thang truoc", "nguyen ban", "thang chin", "thang muoi", "tuan truoc", "nam ngoai", "thang bay", "thang hai", "thang mot", "thang nam", "thang nay", "thang sau", "thang tam", "truoc đay", "ban ngay", "cach đay", "chu nhat", "hang nhi", "ngay mai", "phut nay", "thang 10", "thang 11", "thang 12", "thang ba", "thang tu", "tuan nay", "tuan sau", "bay gio", "gio nay", "hom nay", "hom qua", "nam nay", "nam sau", "thang 1", "thang 2", "thang 3", "thang 4", "thang 5", "thang 6", "thang 7", "thang 8", "thang 9", "thu bay", "thu hai", "thu nam", "thu sau", "tuan le", "thg 10", "thg 11", "thg 12", "thu ba", "thu tu", "thang", "thg 1", "thg 2", "thg 3", "thg 4", "thg 5", "thg 6", "thg 7", "thg 8", "thg 9", "thu 1", "thu 2", "thu 3", "thu 4", "thu 5", "thu 6", "thu 7", "trong", "truoc", "buoi", "chut", "giay", "ngay", "phut", "th 2", "th 3", "th 4", "th 5", "th 6", "th 7", "tuan", "gio", "gmt", "lat", "luc", "nam", "thg", "utc", "am", "ch", "cn", "pm", "sa", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
