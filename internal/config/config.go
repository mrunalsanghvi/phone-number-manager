package config

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Mongo struct {
		URI string `mapstructure:"uri"`
		DB  string `mapstructure:"db"`
	} `mapstructure:"mongo"`
	Logging struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logging"`
}