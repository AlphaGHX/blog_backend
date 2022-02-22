package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
	Local Local `mapstructure:"local"`
	Jwt   Jwt   `mapstructure:"jwt"`
	User  User  `mapstructure:"default-user"`
}
