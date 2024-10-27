package mysqlstore

type Config struct {
	DSN     string `yaml:"dsn"`
	MaxConn int    `yaml:"max_conn"`
	MaxOpen int    `yaml:"max_open"`
	Timeout int64  `yaml:"timeout"`
}

func DefaultCfg() *Config {
	return &Config{
		MaxConn: 100,
		MaxOpen: 10,
		Timeout: 2000,
	}
}
