package main

import (
	"fmt"
	"sort"
	"strings"
	"text/template"
	"unicode/utf8"
)

var (
	fnMap = template.FuncMap{
		"regex":        regex,
		"charset":      charset,
		"localeName":   localeName,
		"parentLocale": parentLocale,
		"sortMap":      sortMap,
	}

	templates = map[string]*template.Template{
		"locale-map":   template.Must(template.New("").Funcs(fnMap).Parse(localeDataMapTemplate)),
		"locale-data":  template.Must(template.New("").Funcs(fnMap).Parse(localeDataTemplate)),
		"lang-order":   template.Must(template.New("").Parse(languageOrderTemplate)),
		"lang-map":     template.Must(template.New("").Parse(languageMapTemplate)),
		"lang-loc-map": template.Must(template.New("").Parse(languageLocalesMapTemplate)),
	}
)

func regex(pattern string) string {
	if pattern == "" {
		return "nil"
	}

	return "regexp.MustCompile(`(?i)" + pattern + "`)"
}

func charset(chars []rune) string {
	if len(chars) == 0 {
		return "nil"
	}

	var strRunes []string
	for _, r := range chars {
		strRune := fmt.Sprintf("%q", r)
		strRune = strings.Trim(strRune, "'")
		strRunes = append(strRunes, strRune)
	}

	return "[]rune(`" + strings.Join(strRunes, "") + "`)"
}

func localeName(language string) string {
	language = strings.ReplaceAll(language, "-", "_")
	language = strings.ReplaceAll(language, " ", "")
	return language + "_Locale"
}

func parentLocale(data LocaleData) string {
	if data.Parent == "" {
		return "nil"
	}

	return "&" + localeName(data.Parent)
}

func sortMap(m map[string]string) []MapEntry {
	var entries []MapEntry
	for k, v := range m {
		entries = append(entries, MapEntry{Key: k, Value: v})
	}

	sort.Slice(entries, func(a, b int) bool {
		eA, eB := entries[a], entries[b]
		lenKeyA := utf8.RuneCountInString(eA.Key)
		lenKeyB := utf8.RuneCountInString(eB.Key)

		if lenKeyA != lenKeyB {
			return lenKeyA > lenKeyB
		}

		if eA.Value != eB.Value {
			return eA.Value < eB.Value
		}

		return eA.Key < eB.Key
	})

	return entries
}

const languageOrderTemplate = `
// Code is generated by script; DO NOT EDIT.

package data

var LanguageOrder = map[string]int{
	{{range $i, $language := . -}}
	"{{$language}}": {{$i}},
	{{end}}
}
`

const languageMapTemplate = `
// Code is generated by script; DO NOT EDIT.

package data

var LanguageMap = map[string][]string{
	{{range $language, $sublanguage := . -}}
	"{{$language}}": {
		{{- range $sub := $sublanguage}}
			"{{$sub}}",
		{{- end}}
	},
	{{end}}
}
`

const languageLocalesMapTemplate = `
// Code is generated by script; DO NOT EDIT.

package data

var LanguageLocalesMap = map[string][]string{
	{{range $language, $locales := . -}}
	"{{$language}}": {
		{{- range $locale := $locales}}
			"{{$locale}}",
		{{- end}}
	},
	{{end}}
}
`

const localeDataMapTemplate = `
// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

type LocaleData struct {
	Name                  string
	DateOrder             string
	NoWordSpacing         bool
	SentenceSplitterGroup int
	Charset               []rune
	Abbreviations         []string
	Simplifications       []ReplacementData
	Translations          map[string]string
	RelativeType          map[string]string
	RelativeTypeRegexes   []ReplacementData
	RxCombined            *regexp.Regexp
	RxExactCombined       *regexp.Regexp
	RxKnownWords          *regexp.Regexp
}

type ReplacementData struct {
	Rx          *regexp.Regexp
	Replacement string
}

func merge(parent *LocaleData, child LocaleData) LocaleData {
	if parent == nil {
		return child
	}

	// Merge list
	child.Charset = append(child.Charset, parent.Charset...)
	child.Abbreviations = append(child.Abbreviations, parent.Abbreviations...)
	child.Simplifications = append(child.Simplifications, parent.Simplifications...)
	child.RelativeTypeRegexes = append(child.RelativeTypeRegexes, parent.RelativeTypeRegexes...)

	// Prepare maps
	if len(child.Translations) == 0 {
		child.Translations = map[string]string{}
	}

	if len(child.RelativeType) == 0 {
		child.RelativeType = map[string]string{}
	}

	// Merge maps
	for word, translation := range parent.Translations {
		child.Translations[word] = translation
	}

	for pattern, translation := range parent.RelativeType {
		child.RelativeType[pattern] = translation
	}

	// Replace regexes
	if child.RxCombined == nil {
		child.RxCombined = parent.RxCombined
	}

	if child.RxExactCombined == nil {
		child.RxExactCombined = parent.RxExactCombined
	}

	if child.RxKnownWords == nil {
		child.RxKnownWords = parent.RxKnownWords
	}

	return child
}

var LocaleDataMap = map[string]*LocaleData {
	{{range $language, $data := . -}}
	"{{$language}}": &{{localeName $language}},
	{{end}}
}
`

const localeDataTemplate = `
var {{localeName .Name}} = merge({{parentLocale .}}, LocaleData {
	Name:                  "{{.Name}}",
	DateOrder:             "{{.DateOrder}}",
	NoWordSpacing:         {{.NoWordSpacing}},
	Charset:               {{charset .Charset}},
	Abbreviations:    []string{ {{range $v := .Abbreviations}}"{{$v}}", {{end}} },
	SentenceSplitterGroup: {{.SentenceSplitterGroup}},
	Simplifications: []ReplacementData{
		{{range $data := .Simplifications -}}
		{ {{regex $data.Pattern}}, "{{$data.Replacement}}" },
		{{end}}
	},
	Translations: map[string]string{
		{{range $e := (sortMap .Translations) -}}
		"{{$e.Key}}": "{{$e.Value}}",
		{{end}}
	},
	RelativeType: map[string]string{
		{{range $e := (sortMap .RelativeType) -}}
		"{{$e.Key}}": "{{$e.Value}}",
		{{end}}
	},
	RelativeTypeRegexes: []ReplacementData{
		{{range $e := (sortMap .RelativeTypeRegexes) -}}
		{ {{regex $e.Key}}, "{{$e.Value}}" },
		{{end}}
	},
	RxCombined: {{regex .CombinedRegexPattern}},
	RxExactCombined: {{regex .ExactCombinedRegexPattern}},
	RxKnownWords: {{regex .KnownWordsPattern}},
})
`
