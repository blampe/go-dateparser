// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var hu_Locale = LocaleData{
	Name:      "hu",
	DateOrder: "YMD.",
	SkipWords: []string{"'", ",", "-", "-a", "-ai", "-akor", "-e", "-ei", "-ekor", "-es", "-i", "-ig", "-je", "-jei", "-ji", "-kor", "-tól", "-től", "-áig", "-án", "-ától", "-éig", "-én", "-étől", "-ös", ".", "/", ";", "@", "[", "]", "|"},
	Simplifications: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)egy(\z|\W|_)`): "${1}1${2}",
	},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) másodperccel ezelőtt(\z|\W|_)`), "${1}$2 second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hónappal ezelőtt(\z|\W|_)`), "${1}$2 month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) másodperc múlva(\z|\W|_)`), "${1}in $2 second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) héttel ezelőtt(\z|\W|_)`), "${1}$2 week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) nappal ezelőtt(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) órával ezelőtt(\z|\W|_)`), "${1}$2 hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) évvel ezelőtt(\z|\W|_)`), "${1}$2 year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hónap múlva(\z|\W|_)`), "${1}in $2 month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) perc múlva(\z|\W|_)`), "${1}in $2 minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hét múlva(\z|\W|_)`), "${1}in $2 week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) nap múlva(\z|\W|_)`), "${1}in $2 day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) óra múlva(\z|\W|_)`), "${1}in $2 hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)ebben a percben(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ebben az órában(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)következő hónap(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) év múlva(\z|\W|_)`), "${1}in $2 year${3}"},
		{regexp.MustCompile(`(\A|\W|_)következő hét(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)következő év(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)másodperccel(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)másodperctől(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) napja(\z|\W|_)`), "${1}$2 day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)előző hónap(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)másodpercek(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)másodpercig(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)tegnapelőtt(\z|\W|_)`), "${1}2 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ez a hónap(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)másodperce(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)szeptember(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)augusztus(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)csütörtök(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)előző hét(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)másodperc(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)december(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)előző év(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ez a hét(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ez az év(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)hónappal(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)november(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)vasárnap(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ezelőtt(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)február(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)hónapja(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)hónapok(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)március(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)október(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)perccel(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)perctől(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)szombat(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)április(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)holnap(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)héttel(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)január(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)július(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)június(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)nappal(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)percek(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)percig(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)péntek(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)szerda(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tegnap(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)órától(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)órával(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)hetek(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)hétfő(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)hónap(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)május(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)múlva(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)napja(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)napok(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)perce(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)szept(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)évvel(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)óráig(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)órája(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)febr(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)hete(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)kedd(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)most(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)márc(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)viii(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)évek(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)órák(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)aug(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)hét(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)iii(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)júl(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jún(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)máj(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)már(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)nap(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)sze(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)szo(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)vas(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)vii(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)xii(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ápr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)éve(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)óra(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)cs(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)de(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)du(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)hó(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ii(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)iv(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ix(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ma(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mp(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)vi(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)xi(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)év(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)i(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)k(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)p(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)v(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)x(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ó(\z|\W|_)`), "${1}hour${2}"},
	},
}
