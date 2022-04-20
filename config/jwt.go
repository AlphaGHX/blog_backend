package config

type Jwt struct {
	PrivateKeyPath string `mapstructure:"private-key"`
	PublicKeyPath  string `mapstructure:"public-key"`
	ExpireTime     int    `mapstructure:"expire-time"`
}
