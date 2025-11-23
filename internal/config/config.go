package config

import (
	"time"
)

// Logging config
type LogConf struct {
	Level string `mapstructure:"level" validate:"loglevel" example:"info"`
	Type  string `mapstructure:"type" validate:"logtype" example:"json"`
}

// Server config
type SerConf struct {
	AddressHttp  string        `mapstructure:"address" validate:"hostname_port" example:"0.0.0.0:3225"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout" validate:"min=0s,max=20m" example:"5s"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout" validate:"min=0s,max=20m" example:"5s"`
}

// Application config
type AppConf struct {
	LogConf LogConf `mapstructure:"log"`
	SerConf SerConf `mapstructure:"http"`
}

// Default config values
func GetMapDefaultsSettings() map[string]interface{} {
	return map[string]interface{}{
		"log.level":         "info",
		"log.type":          "json",
		"http.address":      "0.0.0.0:3225",
		"http.readTimeout":  "0s",
		"http.writeTimeout": "0s",
	}
}
