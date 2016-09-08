package english

import (
	"github.com/bluesuncorp/wash/helpers/trans"
	"github.com/go-playground/locales"
	"github.com/go-playground/log"
	"github.com/go-playground/universal-translator"
)

// Init initializes the english locale translations
func Init(uni *ut.UniversalTranslator) {

	locale := "en"

	en, found := uni.GetTranslator(locale)
	if !found {
		log.WithFields(log.F("locale", locale)).Fatal("Translation not found")
	}

	// cardinals
	trans.AddCardinal(en, "days", "{0} day", locales.PluralRuleOne, false)
	trans.AddCardinal(en, "days", "{0} days", locales.PluralRuleOther, false)

	// ordinals
	trans.AddOrdinal(en, "numsuffix", "{0}st", locales.PluralRuleOne, false)
	trans.AddOrdinal(en, "numsuffix", "{0}nd", locales.PluralRuleTwo, false)
	trans.AddOrdinal(en, "numsuffix", "{0}rd", locales.PluralRuleFew, false)
	trans.AddOrdinal(en, "numsuffix", "{0}th", locales.PluralRuleOther, false)

	// ranges
	trans.AddRange(en, "dayrange", "{0}-{1} days", locales.PluralRuleOther, false)

	// actual message translations
	trans.Add(en, "testonly", "Welcome {0}, you're the {1} to visit our site within the last {2}.", false)

	if err := en.VerifyTranslations(); err != nil {
		log.WithFields(log.F("error", err)).Fatal("Missing Translations!")
	}
}
