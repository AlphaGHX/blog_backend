package config

type Config struct {
	Local Local `mapstructure:"local"`
	Jwt   Jwt   `mapstructure:"jwt"`
	User  User  `mapstructure:"default-user"`
}
