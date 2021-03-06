package config

import (
	"github.com/spf13/viper"
	"log"
)

type configuration struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
	Redis    redis    `yaml:"redis"`
	Jwt      jwt      `yaml:"jwt"`
}

func Config() (cfg *configuration) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	cfg = new(configuration)

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return
}
