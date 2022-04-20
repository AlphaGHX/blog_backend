package config

type Local struct {
	Bloghome      string `mapstructure:"bloghome"`
	MarkdownPath  string `mapstructure:"markdown-path"`
	TopimgPath    string `mapstructure:"topimg-path"`
	ListeningAddr string `mapstructure:"listening-addr"`
}
