package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var instance Config

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	MovieFilePath string `mapstructure:"MOVIE_FILE_PATH"`
	OmdbClientUrl string `mapstructure:"OMDB_CLIENT_URL"`
	OmdbApiKey    string `mapstructure:"OMDB_API_KEY"`
}

func GetConfig() Config {
	var once sync.Once
	once.Do(func() {
		config, err := LoadConfig("./")
		if err != nil {
			log.Fatalf("could not load env variables. %s", err.Error())
		}
		instance = config
	})

	return instance
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}
