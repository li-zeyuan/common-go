package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration interface {
	Verify() error
}

func LoadConfigFile(cfgPath string, cfg Configuration) error {
	buf, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Printf("Read config file(%s) failed: %v.\n", cfgPath, err)
		return err
	}

	if err = yaml.Unmarshal(buf, cfg); err != nil {
		log.Printf("Unmarshal config file(%s) failed: %v.\n", cfgPath, err)
		return err
	}

	if err = cfg.Verify(); err != nil {
		log.Printf("verify config failed: %v.\n", err)
		return err
	}

	log.Printf("	INFO	unmarshal config: %+v", cfg)
	return nil
}

func DecodeConfigFile(path string, config any) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Read config file(%s) failed: %v.\n", path, err)
		return err
	}

	if err = yaml.Unmarshal(buf, config); err != nil {
		log.Printf("Unmarshal config file(%s) failed: %v.\n", path, err)
		return err
	}

	log.Println("	INFO	unmarshal config: ", config)
	return nil
}
