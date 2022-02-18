package config

type Mysql struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Host         string `mapstructure:"host"`
	Charset      string `mapstructure:"charset"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
}
