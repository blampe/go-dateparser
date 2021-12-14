package main

import (
	"strings"
	"text/template"
)

var (
	fnMap = template.FuncMap{
		"localeName":   localeName,
		"parentLocale": parentLocale,
	}

	templates = map[string]*template.Template{
		"locale-map":   template.Must(template.New("").Funcs(fnMap).Parse(localeDataMapTemplate)),
		"locale-data":  template.Must(template.New("").Funcs(fnMap).Parse(localeDataTemplate)),
		"lang-order":   template.Must(template.New("").Parse(languageOrderTemplate)),
		"lang-map":     template.Must(template.New("").Parse(languageMapTemplate)),
		"lang-loc-map": template.Must(template.New("").Parse(languageLocalesMapTemplate)),
	}
)

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
	Parent                *LocaleData
	DateOrder             string
	NoWordSpacing         bool
	SentenceSplitterGroup int
	SkipWords             []string
	PertainWords          []string
	Simplifications       map[*regexp.Regexp]string
	Translations          []TranslationData
}

type TranslationData struct {
	Pattern     *regexp.Regexp
	Translation string
}

var LocaleDataMap = map[string]LocaleData {
	{{range $language, $data := . -}}
	"{{$language}}": {{localeName $language}},
	{{end}}
}
`

const localeDataTemplate = `
var {{localeName .Name}} = LocaleData {
	Name:                  "{{.Name}}",
	Parent:                {{parentLocale .}},
	DateOrder:             "{{.DateOrder}}",
	NoWordSpacing:         {{.NoWordSpacing}},
	SentenceSplitterGroup: {{.SentenceSplitterGroup}},
	SkipWords:    []string{ {{range $v := .SkipWords}}"{{$v}}", {{end}} },
	PertainWords: []string{ {{range $v := .PertainWords}}"{{$v}}", {{end}} },
	Simplifications: map[*regexp.Regexp]string{
		{{range $pattern, $replacement := .Simplifications -}}
		` + "regexp.MustCompile(`{{$pattern}}`)" + `: "{{$replacement}}",
		{{end}}
	},
	Translations: []TranslationData{
		{{range $trans := .Translations -}}
		{` + "regexp.MustCompile(`{{$trans.Pattern}}`)" + `, "{{$trans.Translation}}"},
		{{end}}
	},
}
`
