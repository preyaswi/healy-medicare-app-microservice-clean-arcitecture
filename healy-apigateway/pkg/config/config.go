package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	GoogleClientId string `mapstructure:"YOUR_CLIENT_ID"`
	GoogleSecretId string `mapstructure:"YOUR_CLIENT_SECRET"`
	RedirectURL    string `mapstructure:"REDIRECT_URL"`
	PatientSvc     string `mapstructure:"PATIENT_SVC"`
	DoctorSvc      string `mapstructure:"DOCTOR_SVC"`
	AdminSvc       string `mapstructure:"ADMIN_SVC"`
	ChatSvc        string `mapstructure:"CHAT_SVC"`
	KafkaPort      string `mapstructure:"KAFKA_PORT"`
}

var envs = []string{"PORT", "YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET", "REDIRECT_URL", "PATIENT_SVC", "DOCTOR_SVC", "ADMIN_SVC", "CHAT_SVC", "KAFKA_PORT"}

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
