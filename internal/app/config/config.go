package config

type Config struct {
	Test string `valid:"alphanum,required"`
}
