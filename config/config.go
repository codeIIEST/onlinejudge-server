package config

import (
	"github.com/spf13/viper"
)

//Config store configuration variables
type Config struct {
	Port      string
	DbName    string
	DbUser    string
	DbPass    string
	DbPort    string
	SecretKey string
}

//Read reads the environment variables and returns config struct
func Read() (*Config, error) {
	viper.AutomaticEnv()

	viper.BindEnv("Port", "PORT")
	viper.BindEnv("DbName", "DB_NAME")
	viper.BindEnv("DbUser", "DB_USER")
	viper.BindEnv("DbPass", "DB_PASS")
	viper.BindEnv("DbPort", "DB_PORT")
	viper.BindEnv("SecretKey", "SECRET_KEY")

	viper.SetDefault("Port", "8080")
	viper.SetDefault("DbPort", "5432")

	var conf Config
	err := viper.Unmarshal(&conf)
	return &conf, err
}
