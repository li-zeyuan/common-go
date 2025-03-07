package minioclient

import "time"

type Config struct {
	Endpoint           string        `yaml:"endpoint"`
	Bucket             string        `yaml:"bucket"`
	Policy             string        `yaml:"policy"`
	AccessKeyID        string        `yaml:"access_key_id"`
	SecretAccessKey    string        `yaml:"secret_access_key"`
	PresignedPutExpiry time.Duration `yaml:"presigned_put_expiry"`
	PresignedGetExpiry time.Duration `yaml:"presigned_get_expiry"`
}

func NewDefaultConf() *Config {
	return &Config{
		Policy:             BucketPublicReadPolicy,
		PresignedPutExpiry: time.Hour,
		PresignedGetExpiry: time.Hour * 24,
	}
}
