package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/37acoder/pingcollector/ttype"
)

func LoadConfigFromJsonFile(file string) (ttype.Config, error) {
	readFile, err := os.ReadFile(file)
	fmt.Println(string(readFile))
	conf := ttype.Config{}
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(readFile, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
