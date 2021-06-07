package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//AppConfig Application configuration
type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Cache struct {
		Driver   string `yaml:"driver"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DBNumber int    `yaml:"dbnumber"`
	}
	Endpoint struct {
		Auth            string `yaml:"auth"`
		Commodities     string `yaml:"commodities"`
		ConvertCurrency string `yaml:"convertcurrency"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 5001
	defaultConfig.Database.Driver = "sqlite"
	defaultConfig.Database.Name = "efishery-test"
	defaultConfig.Database.Address = ""
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = ""
	defaultConfig.Database.Password = ""

	defaultConfig.Cache.Driver = "redis"
	defaultConfig.Cache.Address = "localhost"
	defaultConfig.Cache.Port = 6379
	defaultConfig.Cache.DBNumber = 1
	defaultConfig.Cache.Password = ""

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("error to load config file, will use default value ", err)
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}
