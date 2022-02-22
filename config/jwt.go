package config

type Jwt struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime int `mapstructure:"expire-time"`
}
