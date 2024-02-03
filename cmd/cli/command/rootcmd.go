package command

import (
	"fmt"

	"github.com/karkitirtha10/simplebank/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//successfullyh tested cobra .check make file for command
//		make test-cobra

var (
	// Used for flags.
	// cfgFile     string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "simple-bank-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	}

	subCmd = &cobra.Command{
		Use:   "cobratest",
		Short: "My subcommand",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("1 inside pre run. privateKeyPath is: ", viper.GetString("PRIVATE_KEY_PATH"))
			fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("2 inside Run. privateKeyPath is: ", viper.GetString("PRIVATE_KEY_PATH"))
			fmt.Println(viper.GetString("PRIVATE_KEY_PATH"))

			for key := range viper.GetViper().AllSettings() {
				fmt.Println(string(key))

			}
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
		},
	}

	privateKeyPath string
	publicKeyPath  string
)

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVarP(&privateKeyPath, "privateKeyPath", "pvtk", "private key", "private key")
	rootCmd.PersistentFlags().StringVar(&privateKeyPath, "privateKeyPath", "", "private key of rsa key pair")

	err := viper.BindPFlag("PRIVATE_KEY_PATH", rootCmd.PersistentFlags().Lookup("privateKeyPath"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// rootCmd.PersistentFlags().StringVarP(&publicKeyPath, "publicKeyPath", "pubk", "public key", "public key")
	rootCmd.PersistentFlags().StringVar(&publicKeyPath, "publicKeyPath", "", "public key of rsa key pair")
	err = viper.BindPFlag("PUBLIC_KEY_PATH", rootCmd.PersistentFlags().Lookup("publicKeyPath"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//it doesnot work . the value will not be read in init()
	//works  in subcommand.
	// fmt.Println("0 inside init => privateKeyPath is: ", viper.GetString("PRIVATE_KEY_PATH"))
	rootCmd.AddCommand(subCmd)
	rootCmd.AddCommand(generateRSA)

	// fmt.Println("1 inside init => privateKeyPath is: ", viper.GetString("PRIVATE_KEY_PATH"))

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

func initConfig() {
	/*
		if cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			// Search config in home directory with name ".cobra" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigType("yaml")
			viper.SetConfigName(".cobra")
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	*/

	//use this if below code does not work
	// viper.AddConfigPath("/")
	// viper.SetConfigType("env")
	// viper.SetConfigName(".env")

	_ = config.LoadConfig()

	//this below code works but is repaced by LoadConfig
	/*
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
