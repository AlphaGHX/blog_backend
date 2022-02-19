package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
	Local Local `mapstructure:"local"`
}
