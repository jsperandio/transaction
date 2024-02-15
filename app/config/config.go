package config

import (
	"log/slog"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

const ConfEnvironment = "CONF"

var instance *koanf.Koanf

func Load() {
	instance = koanf.New(".")

	ce := os.Getenv(ConfEnvironment)
	if ce != "" {
		files := strings.Split(ce, ",")

		for _, f := range files {
			parser := yaml.Parser()
			if err := instance.Load(file.Provider(f), parser); err != nil {
				slog.Error("Failed to load config file", err)
				panic(err)
			}
		}

	}
}

func UnmarshalWithPath(path string, o interface{}) error {
	return instance.UnmarshalWithConf(path, &o, koanf.UnmarshalConf{Tag: "config"})
}
