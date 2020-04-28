package fincloud

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path string
}

type ConfigData struct {
	CertificateList []Certificate `yaml:"certs"`
}

type Certificate struct {
	SubaccountName string `yaml:"nsa"`
	Key            string `yaml:"key"`
	CreateYmdt     string `yaml:"created"`
}

func (c Config) Parse() (*ConfigData, error) {
	data, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return nil, fmt.Errorf("Error reading certificate file: %v", err)
	}

	configData := ConfigData{}

	err = yaml.Unmarshal([]byte(data), &configData)
	if err != nil {
		return nil, err
	}

	return &configData, nil
}

func (c Config) Write(configData *ConfigData) error {
	f, err := os.Create(c.Path)
	if err != nil {
		return err
	}

	n, err := yaml.Marshal(configData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = f.WriteString(string(n))
	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
