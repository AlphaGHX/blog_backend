package config

type User struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
