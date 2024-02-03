package command

import (
	"github.com/karkitirtha10/simplebank/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateRSA = &cobra.Command{
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
