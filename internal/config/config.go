package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(configPath string) (Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

type Config struct {
	DB     Database
	Server Server
	Key    JwtSecret
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Server struct {
	Port int
}

type JwtSecret struct {
	JwtSecret string
}
