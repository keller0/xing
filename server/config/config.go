package config

type AllConf struct {
	App AppConfig
}

type AppConfig struct {
	LogLevel       string `yaml:"log_level"`
	LogPath        string `yaml:"log_path"`
	JwtKey         string `yaml:"jwt_key"`
	AllowOrigin    string `yaml:"allow_origins"`
	FileSaveFolder string `yaml:"file_save_folder"`
}
