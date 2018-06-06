package config

import (
	"flag"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	env  string
	port string
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

	flag.StringVar(&env, "env", os.Getenv("env"), "set up env variable")
	flag.StringVar(&port, "port", os.Getenv("port"), "set up server port")

	flag.Parse()

	if env != "" {
		os.Setenv("env", env)
	}

	if port != "" {
		os.Setenv("port", port)
	}
}
