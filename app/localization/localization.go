package localization

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Localization scoped services
type Localization struct {
	Internationalisation *i18n.Bundle
	//lang string
	localizer *i18n.Localizer
}

func (yo *Localization) SetLang(lang string) {
	yo.localizer = i18n.NewLocalizer(yo.Internationalisation, lang)
}

// get and directly translate using localizer
func (yo *Localization) GetLang() *i18n.Localizer {
	return yo.localizer
}

// translate using configed default language
func (yo *Localization) Translate(messageId string) string {
	return yo.localizer.MustLocalize(
		&i18n.LocalizeConfig{MessageID: messageId},
	)
}

func (yo *Localization) TranslateTemplate(
	messageId string,
	templateData map[string]interface{},
) string {
	return yo.localizer.MustLocalize(
		&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: templateData,
		},
	)
}

// translatePlural translate using configed default language
func (yo *Localization) TranslatePlural(
	messageId string,
	templateData map[string]interface{},
	pluralCount int,
) string {
	return yo.localizer.MustLocalize(
		&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: templateData,
			PluralCount:  pluralCount,
		},
	)

}

func (yo *Localization) TranslateLang(lang string, messageId string) string {
	newLocalizer := i18n.NewLocalizer(yo.Internationalisation, lang)
	return newLocalizer.MustLocalize(
		&i18n.LocalizeConfig{MessageID: messageId},
	)
}

func (yo *Localization) TranslatePluralLang(
	lang string, //el, en, ne. filename without extension in lang direcrory
	messageId string,
	templateData map[string]interface{},
	pluralCount int,
) string {
	newLocalizer := i18n.NewLocalizer(yo.Internationalisation, lang)
	return newLocalizer.MustLocalize(
		&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: templateData,
			PluralCount:  pluralCount,
		},
	)

}

type ILocalization interface {
	SetLang(lang string)

	GetLang() *i18n.Localizer

	Translate(messageId string) string

	TranslateTemplate(
		messageId string,
		templateData map[string]interface{},
	) string

	TranslatePlural(
		messageId string,
		templateData map[string]interface{},
		pluralCount int,
	) string

	TranslateLang(lang string, messageId string) string

	TranslatePluralLang(
		lang string, //el, en, np. filename without extension in lang direcrory
		messageId string,
		templateData map[string]interface{},
		pluralCount int,
	) string
}

func InitializeInternationalisation() *i18n.Bundle {
	//language.Nepali.String() == ne
	//language.English.String() == en
	bundle := i18n.NewBundle(language.English) //default: english . //todo: language.English should come from method arg. and supplied the value from config. config.SingleConfig().Locale
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("lang/en.json")
	bundle.MustLoadMessageFile("lang/ne.json")
	return bundle
}

func NewLocalization(lang string) ILocalization {
	internationalisationBundle := InitializeInternationalisation()
	return &Localization{
		Internationalisation: internationalisationBundle,
		localizer:            i18n.NewLocalizer(internationalisationBundle, lang),
	}
}
