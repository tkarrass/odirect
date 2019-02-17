package config // import "udico.de/odirect/config"

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	debug bool
	Cache struct {
		File string
		KeySeed  int
	}
}

var C config

func init() {
	// Provide some neat default values
	viper.SetDefault("Cache.File", "odirect.cache")
	viper.SetDefault("Cache.Key", 31337)
}

// Reads the configuration file and maps it to the configuration struct, respecting
// the parameters' default values.
func ReadConfig() {
	viper.SetConfigName("odirect")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Warn("cannot read config file")
		return
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		log.WithError(err).Warn("cannot parse config")
	}
}
