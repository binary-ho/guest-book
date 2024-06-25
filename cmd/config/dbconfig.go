package config

import (
	"fmt"
	"guestbook/common/util"
)

type DbConfig struct {
	Db struct {
		Mysql struct {
			User     string `yaml:"user"`
			DbName   string `yaml:"dbname"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Sslmode  string `yaml:"sslmode"`
			Driver   string `yaml:"driver"`
		}
	}
}

const (
	dbPropertyPath        = "common\\property\\db-property.yml"
	dataSourceFormat      = "%s:%s@/%s?parseTime=true"
	databaseAddressFormat = "%s:%s"
)

func getDbConfig() DbConfig {
	config := DbConfig{}
	err := util.FetchYAML(dbPropertyPath, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func (config DbConfig) DataSource() string {
	user := config.Db.Mysql.User
	password := config.Db.Mysql.Password
	dbname := config.Db.Mysql.DbName
	return fmt.Sprintf(dataSourceFormat, user, password, dbname)
}

func (config DbConfig) DatabaseAddress() string {
	host := config.Db.Mysql.Host
	port := config.Db.Mysql.Port
	return fmt.Sprintf(databaseAddressFormat, host, port)
}

func (config DbConfig) Driver() string {
	return config.Db.Mysql.Driver
}
