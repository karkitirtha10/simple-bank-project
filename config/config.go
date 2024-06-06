package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	//App
	AppName string `mapstructure:"APP_NAME"`
	AppUrl  string `mapstructure:"APP_URL"`
	AppPort string `mapstructure:"APP_PORT"`
	//Database
	DbUrl        string `mapstructure:"DB_URL"`
	MigrationUrl string `mapstructure:"MIGRATION_URL"`
	//Auth
	PrivateKeyPath                  string `mapstructure:"PRIVATE_KEY_PATH"`
	PublicKeyPath                   string `mapstructure:"PUBLIC_KEY_PATH"`
	ClientCredentialsTokenExpiresIn int64  `mapstructure:"CLIENT_CREDENTIALS_TOKEN_EXPIRES"`
	OAuthPersonalAccessClientId     string `mapstructure:"OAUTH_PERSONAL_ACCESS_CLIENT_ID"`
	OAuthPersonalAccessClientSecret string `mapstructure:"OAUTH_PERSONAL_ACCESS_CLIENT_SECRET"`
	PersonalAccessTokenExpiresIn    int64  `mapstructure:"PERSONAL_ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenExpiresIn           int64  `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
	//locale
	Locale string `mapstructure:"Locale"`
}

var cnf *Config

func init() {
	cnf = LoadConfig()
}

func SingleConfig() *Config {
	return cnf
}

// LoadConfig loads and  gets the config.
// `panics` if it cannot read the env file or cannot unmarshall
func LoadConfig() (c *Config) {
	// viper.AddConfigPath("./pkg/common/config/envs") // scan envs folder

	viper.SetConfigFile(".env") //reads from /pkg/common/envs/.env exclusively
	viper.AddConfigPath(".")    //reads from /pkg/common/envs/.env exclusively
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
