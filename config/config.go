package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	env  string
	port string
)

const (
	configPath = "./config/env.yml"
)

func Load() {
	file, err := ioutil.ReadFile(configPath)

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

// MustGet return env and return error if not found key
func MustGet(key string) (result string) {
	var env = os.Getenv(key)
	if env == "" {
		msg := fmt.Sprintf(
			"can not find %s in `ENV`, please checkout your env file.",
			key,
		)

		panic(msg)
	}

	return env
}
