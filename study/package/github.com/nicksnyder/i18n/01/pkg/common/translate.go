package common

import (
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Translate(lang, id string, data map[string]interface{}) string {
	localizer := i18n.NewLocalizer(Bundle, lang, "zh-CN")
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if err != nil {
		log.Fatalf("Failed to localize message: %v", err)
	}
	return msg
}
