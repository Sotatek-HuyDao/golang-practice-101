package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Version string `envconfig:"version" required:"true"`
}

func MustParseConfig(prefix ...string) *Config {
	prf := ""
	if len(prefix) > 0 {
		prf = prefix[0]
	}

	cfg, err := ParseConfig(prf)
	if err != nil {
		panic(err)
	}
	return cfg
}

func ParseConfig(prefix string) (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
