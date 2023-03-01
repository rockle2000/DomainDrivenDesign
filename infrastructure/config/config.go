package config

import "github.com/spf13/viper"

type AppConfig struct {
	Port     int
	DBConfig DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
	TimeZone string
}

func LoadConfig(path string) (*AppConfig, error) {
	var config *AppConfig
	viper.AddConfigPath(path)

	viper.SetConfigName("config-base")
	//viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	//name of config file
	viper.SetConfigName("config")
	// extension of config file (eg: .env)
	//viper.SetConfigType("yml")
	if err := viper.MergeInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func (cf *AppConfig) GetDbConfig() DBConfig {
	return cf.DBConfig
}
