package config

import (
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Domain string `toml:"domain"`
	Mongo  *Mongo
}

type Mongo struct {
	ConnectionString string `toml:"connectionString"`
	Database         string `toml:"database"`
	Collection       string `toml:"collection"`
}

var conf *Config
var configError error
var decodeConfig sync.Once

// GetConfig Returns reference to Application configuration
func GetConfig() *Config {
	decodeConfig.Do(initConfig)
	return conf
}

// GetMongoConfig Returns reference to Mongo configuration
func GetMongoConfig() *Mongo {
	return GetConfig().Mongo
}

// initConfig Loads `config.toml` into config.Config struct
func initConfig() {
	conf = &Config{}
	if _, err := toml.DecodeFile("config.toml", conf); err != nil {
		configError = err
		log.Fatal("Failed to load configuration file", err)
	}
}
