package config

import (
	"github.com/BurntSushi/toml"
)

func Init(filepath string) (Config, error) {
	c := defaultConfig()
	if _, err := toml.DecodeFile(filepath, &c); err != nil {
		panic(err)
	}

	err := c.Smtp.validate()

	return c, err
}
