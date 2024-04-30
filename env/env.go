package env

import "github.com/spf13/viper"

type Config struct {
	HttpPort        string
	LearningService *LearningService
}

type LearningService struct {
	Host string
	Port string
}

var Conf *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("LEARNING_SERVICE_HOST")
	viper.BindEnv("LEARNING_SERVICE_PORT")

	Conf = &Config{
		HttpPort: viper.GetString("HTTP_PORT"),
		LearningService: &LearningService{
			Host: viper.GetString("LEARNING_SERVICE_HOST"),
			Port: viper.GetString("LEARNING_SERVICE_PORT"),
		},
	}
}
