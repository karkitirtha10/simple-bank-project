package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MigrationUrl string `mapstructure:"MIGRATION_URL"`
	DbUrl        string `mapstructure:"DB_URL"`
	AppPort      string `mapstructure:"APP_PORT"`
}

func LoadConfig() (c Config, err error) {
	// viper.AddConfigPath("./pkg/common/config/envs") // scan envs folder
	viper.SetConfigFile(".env") //reads from /pkg/common/envs/.env exclusivly
	viper.AddConfigPath("/")    //reads from /pkg/common/envs/.env exclusivly

	// viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		// log.Fatal(err)
		return
	}

	//load values from env to config struct
	err = viper.Unmarshal(&c)

	return
}
