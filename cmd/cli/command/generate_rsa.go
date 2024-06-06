package command

import (
	"fmt"
	"github.com/karkitirtha10/simplebank/app/services"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	generateRSA = &cobra.Command{
		Use:   "generate-rsa",
		Short: "command to generate private public key pair",
		Run: func(cmd *cobra.Command, args []string) {
			rsa := services.RSAGenerator{}
			rsa.Generate(
				2048,
				viper.GetString("PRIVATE_KEY_PATH"),
				viper.GetString("PUBLIC_KEY_PATH"),
			)
		},
	}

	privateKeyPath string

	publicKeyPath string
)

func init() {
	rootCmd.AddCommand(generateRSA)

	//private key path : if privateKeyPath flag is not supplied, use PRIVATE_KEY_PATH from env file
	rootCmd.PersistentFlags().StringVar(
		&privateKeyPath,
		"privateKeyPath",
		"",
		"private key of rsa key pair",
	)
	err := viper.BindPFlag(
		"PRIVATE_KEY_PATH",
		rootCmd.PersistentFlags().Lookup("privateKeyPath"),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//public Key Path : if publicKeyPath flag is not supplied, use PUBLIC_KEY_PATH from env file
	generateOAuthClient.PersistentFlags().StringVar(
		&publicKeyPath,
		"publicKeyPath",
		"xcz",
		"public key of rsa key pair",
	)
	err = viper.BindPFlag(
		"PUBLIC_KEY_PATH",
		rootCmd.PersistentFlags().Lookup("publicKeyPath"),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
