package utils

import (
	"fmt"
	"os"
)

func GetURLByID(name string, id interface{}) string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s/%s/%s", hostname, name, id)
	return url
}
