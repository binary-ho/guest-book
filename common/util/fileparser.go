package util

import (
	"gopkg.in/yaml.v3"
	"os"
)

func fetchYAML(fileName string, config interface{}) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, config)
}
