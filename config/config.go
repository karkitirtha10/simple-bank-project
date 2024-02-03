package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	AppPort        string `mapstructure:"APP_PORT"`
	DbUrl          string `mapstructure:"DB_URL"`
	MigrationUrl   string `mapstructure:"MIGRATION_URL"`
	PrivateKeyPath string `mapstructure:"PRIVATE_KEY_PATH"`
	PublicKeyPath  string `mapstructure:"PUBLIC_KEY_PATH"`
}

// LoadConfig loads and  gets the config.
// `panics` if it cannot read the env file or cannot unmarshall
func LoadConfig() (c Config) {
	// viper.AddConfigPath("./pkg/common/config/envs") // scan envs folder
	viper.SetConfigFile(".env") //reads from /pkg/common/envs/.env exclusivly
	viper.AddConfigPath("/")    //reads from /pkg/common/envs/.env exclusivly

	// viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		// log.Fatal(err)
		panic(err.Error())
	}

	//load values from env to config struct
	err = viper.Unmarshal(&c)

	if err != nil {
		// log.Fatal(err)
		panic(err.Error())
	}
	return
}
