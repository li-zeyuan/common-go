package txcloud

type Config struct {
	AppID               string   `yaml:"app_id"`
	Bucket              string   `yaml:"bucket"`
	Region              string   `yaml:"region"`
	SecretID            string   `yaml:"secret_id"`
	SecretKey           string   `yaml:"secret_key"`
	ObjectLimitSizeByte int      `yaml:"object_limit_size_byte"`
	AllowContentType    []string `yaml:"allow_content_type"`
}

func DefaultConfig() *Config {
	return &Config{
		AppID:               "",
		Bucket:              "",
		Region:              "",
		SecretID:            "",
		SecretKey:           "",
		ObjectLimitSizeByte: 1024 * 1024 * 10,
		AllowContentType:    []string{"image/jpeg", "image/png"},
	}
}
