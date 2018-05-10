package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func Load() {
	file, err := ioutil.ReadFile("./config/env.yml")

	if err != nil {
		panic(err)
	}

	var config = make(map[string]string)
	yamlErr := yaml.Unmarshal(file, &config)

	if yamlErr != nil {
		panic(yamlErr)
	}

	for k, v := range config {
		os.Setenv(k, v)
	}
}
