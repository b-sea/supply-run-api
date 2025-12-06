package cli

const cfgEnvPrefix = "SUPPLYRUN"

// Config is the supply run configuration.
type Config struct {
	Server struct {
		Port         int `config:"port"`
		ReadTimeout  int `config:"readTimeout"`
		WriteTimeout int `config:"writeTimeout"`
	} `config:"server"`

	MariaDB struct {
		Host     string `config:"host"`
		Username string `config:"username"`
		Password string `config:"password"`
	} `config:"mariadb"`

	Logger struct {
		Level string `config:"level"`
	} `config:"logger"`
}

func defaultConfig() Config {
	return Config{
		Server: struct {
			Port         int `config:"port"`
			ReadTimeout  int `config:"readTimeout"`
			WriteTimeout int `config:"writeTimeout"`
		}{
			Port:         5000, //nolint: mnd
			ReadTimeout:  5,    //nolint: mnd
			WriteTimeout: 5,    //nolint: mnd
		},
		MariaDB: struct {
			Host     string `config:"host"`
			Username string `config:"username"`
			Password string `config:"password"`
		}{},
		Logger: struct {
			Level string `config:"level"`
		}{
			Level: "info",
		},
	}
}
