package main

import (
	"github.com/asaskevich/govalidator"
	appConfig "github.com/xdimedrolx/moly/internal/app/config"
	"github.com/xdimedrolx/moly/pkg/config"
	"github.com/xdimedrolx/moly/pkg/platform/log"
)

type Config struct {
	Log      log.Config       `valid:"required"`
	App      appConfig.Config `valid:"required"`
	HttpPort int              `valid:"required,port"`
}

func NewDefaultConfig() Config {
	return Config{
		Log:      log.NewDefaultConfig(),
		HttpPort: 3000,
	}
}

func (c Config) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

var _ config.Config = (*Config)(nil)
