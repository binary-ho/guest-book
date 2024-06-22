package util

import (
	"fmt"
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
	FileName        = "common\\property\\db-property.yml"
	DataSource      = "%s:%s@/%s?parseTime=true"
	DatabaseAddress = "%s:%s"
)

func GetDbConfig() (DbConfig, error) {
	config := DbConfig{}
	err := fetchYAML(FileName, &config)
	return config, err
}

func (config DbConfig) DataSource() string {
	user := config.Db.Mysql.User
	password := config.Db.Mysql.Password
	dbname := config.Db.Mysql.DbName
	return fmt.Sprintf(DataSource, user, password, dbname)
}

func (config DbConfig) DatabaseAddress() string {
	host := config.Db.Mysql.Host
	port := config.Db.Mysql.Port
	return fmt.Sprintf(DatabaseAddress, host, port)
}

func (config DbConfig) Driver() string {
	return config.Db.Mysql.Driver
}
