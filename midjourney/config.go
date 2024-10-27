package midjourney

type Config struct {
	Address       string `yaml:"address"`
	Key           string `yaml:"key"`
	Pattern       string `yaml:"pattern"`
	UseImageProxy bool   `yaml:"use_image_proxy"`
}

func NewDefault() *Config {
	return &Config{
		Address:       "https://api.openai-hk.com/",
		Key:           "",
		Pattern:       "fast", // fast; relax
		UseImageProxy: true,
	}
}
