package command

import (
	"github.com/karkitirtha10/simplebank/app"
	"github.com/karkitirtha10/simplebank/config"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "simple-bank-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	}
	application app.Application
)

func init() {
	application = app.InitializeApp()
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	config.LoadConfig()

	//this below code works but is repaced by LoadConfig
	/*
		//use this if below code does not work
		// viper.AddConfigPath("/")
		// viper.SetConfigType("env")
		// viper.SetConfigName(".env")

			viper.SetConfigFile(".env") //config file to read
		viper.AddConfigPath("/")    //folder where .env file (metioned above) exists
		viper.SetConfigType("env")  //type of config file. other example yaml, json

		viper.AutomaticEnv()
		err := viper.ReadInConfig()

		if err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	*/
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
