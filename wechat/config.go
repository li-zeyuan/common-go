package wechat

type Config struct {
	AppId    string    `yaml:"app_id"`
	Secret   string    `yaml:"secret"`
	Pay      PayConfig `yaml:"pay"`
	RobotUrl string    `yaml:"robot_url"`
}

type PayConfig struct {
	Enable                     bool   `yaml:"enable"`
	PrivateKeyPath             string `yaml:"private_key_path"`
	MchID                      string `yaml:"mch_id"`
	MchCertificateSerialNumber string `yaml:"mch_certificate_serial_number"`
	MchAPIv3Key                string `yaml:"mch_api_v3_key"`
	NotifyUrl                  string `yaml:"notify_url"`
}
