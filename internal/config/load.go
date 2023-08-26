package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

const (
	delimiter = "."

	tagName = "koanf"

	upTemplate     = "================ Loaded Configuration ================"
	bottomTemplate = "======================================================"
)

func Load(print bool) *Config {
	k := koanf.New(delimiter)

	// load default configuration from default function
	if err := k.Load(structs.Provider(defaultConfig(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load config from environment variables
	if err := LoadEnv(k); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	config := Config{}
	var tag = koanf.UnmarshalConf{Tag: tagName}
	if err := k.UnmarshalWithConf("", &config, tag); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	if print {
		// pretty print loaded configuration using provided template
		log.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config
}

const (
	envPrefix    = "GHOLAM_CLI"
	envSeperator = "__"
)

func LoadEnv(k *koanf.Koanf) error {
	var prefix = envPrefix + envSeperator

	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, prefix))
		return strings.ReplaceAll(base, envSeperator, delimiter)
	}

	// load environment variables
	if err := k.Load(env.Provider(prefix, delimiter, callback), nil); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}
