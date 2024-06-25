package config

import (
	"fmt"
	"guestbook/common/util"
)

type ServerConfig struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	}
}

const (
	serverPropertyPath  = "common\\property\\local-server-property.yml"
	serverAddressFormat = "%s:%d"
)

func getServerConfig() ServerConfig {
	config := ServerConfig{}
	err := util.FetchYAML(serverPropertyPath, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func (config ServerConfig) ServerAddress() string {
	host := config.Server.Host
	port := config.Server.Port
	return fmt.Sprintf(serverAddressFormat, host, port)
}
