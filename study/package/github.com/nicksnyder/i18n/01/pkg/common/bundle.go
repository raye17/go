package common

import (
	"embed"
	"io/fs"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

//go:embed language/*.yaml
var localeFs embed.FS
var Bundle *i18n.Bundle

func InitBundle() {
	Bundle = i18n.NewBundle(language.Chinese)
	Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	entries, err := fs.ReadDir(localeFs, "language")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}
	for _, entry := range entries {
		data, err := localeFs.ReadFile("language/" + entry.Name())
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		_, err = Bundle.ParseMessageFileBytes(data, entry.Name())
		if err != nil {
			log.Fatalf("Failed to parse message file: %v", err)
		}
	}
}
