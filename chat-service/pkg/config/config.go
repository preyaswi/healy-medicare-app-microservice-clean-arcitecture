package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl  string `mapstructure:"DB_URL"`
	DBname string `mapstructure:"DB_NAME"`

	Port      string `mapstructure:"PORT"`
	KafkaPort string `mapstructure:"KAFKA_PORT"`
	KafkaTpic string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"DB_URL", "DB_NAME", "PORT", "KAFKA_PORT", "KAFKA_TOPIC",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
