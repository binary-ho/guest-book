package main

import (
	"guestbook/common/util"
)

type Config struct {
	Db util.DbConfig
}

func ApplicationConfig() (Config, error) {
	dbConfig, err := util.GetDbConfig()
	return Config{
		Db: dbConfig,
	}, err
}
