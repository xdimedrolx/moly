package log

import "github.com/asaskevich/govalidator"

type Config struct {
	// Format specifies the output log format.
	// Accepted values are: json, logfmt
	Format string `valid:"required,in(json|logfmt)"`

	// Level is the minimum log level that should appear on the output.
	Level string `valid:"required,in(debug|info|error)"`

	// NoColor makes sure that no log output gets colorized.
	NoColor bool `valid:"required"`
}

func NewDefaultConfig() Config {
	return Config{Format: "json", Level: "info", NoColor: true}
}

func (c Config) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
