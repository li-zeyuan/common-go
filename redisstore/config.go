package redisstore

import "time"

const DefaultUnlockTimeDuration = time.Second * 10

type Config struct {
	DSN                       string        `yaml:"dsn"`
	DefaultUnlockTimeDuration time.Duration `yaml:"default_unlock_time_duration"`
}

func NewDefaultConf() *Config {
	return &Config{
		DefaultUnlockTimeDuration: DefaultUnlockTimeDuration,
	}
}
