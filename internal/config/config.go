package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	ApplicationConfig `yaml:"application"`
	DBConfiguration   `yaml:"db_configuration"`
}

type DBConfiguration struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

func ProvideConfig() *Config {
	conf := Config{}
	file, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(file), &conf)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	return &conf
}
