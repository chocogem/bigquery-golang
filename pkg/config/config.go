package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST" `
	DBName         string `mapstructure:"DB_NAME" `
	DBUser         string `mapstructure:"DB_USER" `
	DBPort         string `mapstructure:"DB_PORT" `
	DBPassword     string `mapstructure:"DB_PASSWORD" `
	BQProjectId    string `mapstructure:"BQ_PROJECT_ID" `
	BQDatasetName  string `mapstructure:"BQ_DATASET" `
	LogrusLogLevel string `mapstructure:"LOGRUS_LOG_LEVEL" `
}

// List of environment variables to fetch for
var envs = []string{
	"BQ_PROJECT_ID","BQ_DATASET",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile("local.env")
	viper.ReadInConfig()

	// Bind from environment variable
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
