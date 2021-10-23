package config

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
	"github.com/sherifabdlnaby/configuro"
	"github.com/xdimedrolx/moly/pkg/config"
)

func Load(path string, cfg config.Config) error {
	c, err := configuro.NewConfig(
		configuro.WithLoadFromEnvVars("APP"),
		configuro.WithLoadFromConfigFile(path, true),
	)

	if err != nil {
		return err
	}

	return c.Load(cfg)
}

func LoadAndValidate(path string, cfg config.Config) {
	err := Load(path, cfg)
	emperror.Panic(errors.Wrap(err, "failed to load configuration"))

	err = cfg.Validate()
	emperror.Panic(errors.Wrap(err, "failed to validate configuration"))
}
