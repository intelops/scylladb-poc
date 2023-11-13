package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ScyllaHosts         string `mapstructure:"SCYLLA_HOSTS"`
	ScyllaKeyspace      string `mapstructure:"SCYLLA_KEYSPACE"`
	ScyllaMigrationsDir string `mapstructure:"SCYLLA_MIGRATIONS_DIR"`
}

func LoadConfig() (config Config ,err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}
	err = viper.Unmarshal(&config)
	return
}
