package config

type Local struct {
	BloghomeDir     string `mapstructure:"bloghome-dir"`
	UserhomeDir     string `mapstructure:"userhome-dir"`
	MarkdownFile    string `mapstructure:"markdown-file"`
	TopimgFile      string `mapstructure:"topimg-file"`
	AdminAvatarFile string `mapstructure:"admin-avatar-file"`
	DataDir         string `mapstructure:"data-dir"`
	ListeningAddr   string `mapstructure:"listening-addr"`
}
