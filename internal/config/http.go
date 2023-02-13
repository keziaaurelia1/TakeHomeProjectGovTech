package config

import "github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/common/envreader"

type HTTPConfig struct {
	Address string `env:"HTTP_ADDRESS" default:""`
	Port    string `env:"HTTP_PORT" default:"8080" validate:"numeric"`
}

func ProvideHttpConfig() HTTPConfig {
	cfg := HTTPConfig{}
	envreader.BindEnv(&cfg)
	return cfg
}
