package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	IsDebug bool `yaml:"IsDebug"`
	PORT    int  `yaml:"PORT"`
}

var CONFIG ConfigStruct

func InitConfig() {
	content, _ := ioutil.ReadFile("config/config.yaml")

	CONFIG = ConfigStruct{}

	err := yaml.Unmarshal(content, &CONFIG)

	if CONFIG.IsDebug {
		fmt.Printf("config loading --- %v, %v \n", err, CONFIG)
	}
}
