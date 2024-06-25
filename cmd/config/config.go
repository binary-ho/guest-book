package config

type ApplicationConfig struct {
	Db           DbConfig
	ServerConfig ServerConfig
}

func ApplicationConfiguration() *ApplicationConfig {
	return &ApplicationConfig{
		Db:           getDbConfig(),
		ServerConfig: getServerConfig(),
	}
}
