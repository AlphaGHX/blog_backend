package config

type Jwt struct {
	SecretPath string `mapstructure:"secret-path"`
	ExpireTime int    `mapstructure:"expire-time"`
}
