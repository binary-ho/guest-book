package config

type ApplicationConfig struct {
	Db           dbConfig
	ServerConfig serverConfig
}

func ApplicationConfiguration() *ApplicationConfig {
	return &ApplicationConfig{
		Db:           getDbConfig(),
		ServerConfig: getServerConfig(),
	}
}
