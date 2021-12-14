// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var pt_Locale = LocaleData{
	Name:         "pt",
	DateOrder:    "DMY",
	SkipWords:    []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "cerca", "de", "e", "|", "às"},
	PertainWords: []string{"de"},
	Simplifications: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)alguns segundos(\z|\W|_)`): "${1}44 segundos${2}",
		regexp.MustCompile(`(\A|\W|_)um(\z|\W|_)`):              "${1}1${2}",
		regexp.MustCompile(`(\A|\W|_)uma(\z|\W|_)`):             "${1}1${2}",
	},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)em (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) segundos(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) segundo(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) semanas(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) semana(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) horas(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) meses(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)próxima semana(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)semana passada(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) mins(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) segs(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) anos(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) dias(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) hora(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)segunda-feira(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) seg(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) ano(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) dia(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) mês(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) seg(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) sem(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)quarta-feira(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)quinta-feira(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ano passado(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)esta semana(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)este minuto(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mês passado(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)próximo ano(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)próximo mês(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)sexta-feira(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)terça-feira(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)em (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) h(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)anteontem(\z|\W|_)`), "${1}2 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)esta hora(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)fevereiro(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)septembro(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)dezembro(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)este ano(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)este mês(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)novembro(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)segundos(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)setembro(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)domingo(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)janeiro(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)minutos(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)outubro(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)segunda(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)segundo(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)semanas(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)agosto(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)amanhã(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)quarta(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)quinta(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)semana(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)sábado(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)abril(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)agora(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)atrás(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)horas(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)julho(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)junho(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)março(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)meses(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ontem(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sexta(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)terça(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)anos(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)dias(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)hoje(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)hora(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)maio(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)abr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ago(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ano(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)dez(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)dia(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)dom(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)fev(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mai(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)mês(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)out(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)qua(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)qui(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sab(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)seg(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sem(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)set(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)sex(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sáb(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ter(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)em(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)há(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`), "${1}second${2}"},
	},
}

var pt_AO_Locale = LocaleData{
	Name:      "pt-AO",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_CH_Locale = LocaleData{
	Name:      "pt-CH",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_CV_Locale = LocaleData{
	Name:      "pt-CV",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_GQ_Locale = LocaleData{
	Name:      "pt-GQ",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_GW_Locale = LocaleData{
	Name:      "pt-GW",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_LU_Locale = LocaleData{
	Name:      "pt-LU",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_MO_Locale = LocaleData{
	Name:      "pt-MO",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_MZ_Locale = LocaleData{
	Name:      "pt-MZ",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_PT_Locale = LocaleData{
	Name:      "pt-PT",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_ST_Locale = LocaleData{
	Name:      "pt-ST",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}

var pt_TL_Locale = LocaleData{
	Name:      "pt-TL",
	Parent:    &pt_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundos(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minutos(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) segundo(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semanas(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) minuto(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) semana(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) horas(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) meses(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) anos(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dias(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) hora(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) ano(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) dia(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) min(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) mês(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) sem(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) h(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)dentro de (\d+) s(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)há (\d+) s(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)da manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)da tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)manhã(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
	},
}
