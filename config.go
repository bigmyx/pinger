package main

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Pinger struct {
		Threads int `yaml:"threads", envconfig:"THREADS"`
	} `yaml:"pinger"`
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
