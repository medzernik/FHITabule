package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Nastavenie struct {
		Poschodie int    `yaml:"poschodie"`
		APIkluc   string `yaml:"apikluc"`
	}
}

var Cfg Config

func Initialization() {

	readFile(&Cfg)
	fmt.Printf("%+v", Cfg)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
