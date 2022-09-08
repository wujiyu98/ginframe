package config

import (
	"bytes"
	_ "embed"

	"github.com/spf13/viper"
)

//go:embed config.toml
var cfg string

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBufferString(cfg))

}
