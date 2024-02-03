package command

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/db"
	"github.com/karkitirtha10/simplebank/enums"
	datamodel "github.com/karkitirtha10/simplebank/models"
	"github.com/karkitirtha10/simplebank/repositories"
	"github.com/spf13/cobra"
)

var (
	generateOAuthClient = &cobra.Command{
		Use:   "generate-oauth-client",
		Short: "command to generate oauth client",
		ValidArgs: []string{
			string(enums.PERSONAL),
			string(enums.CLIENT),
		},
		Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {

			config := config.LoadConfig()
			oauthClientRepository := repositories.NewOAuthClientRepository(db.Connection(config.DbUrl))

			secret := uuid.New().String()
			ocType := enums.OAuthClientTypeShortNameEnum(args[0]).ToOAuthClientTypeEnum()
			if clientName == "" {
				clientName = config.AppName + " " + string(ocType) + " client"
			}

			ch := make(chan datamodel.InsertOAuthClientResult)
			go oauthClientRepository.Insert(
				ch,
				clientName,
				secret,
				ocType,
				false,
			)
			insertOAuthClientResult := <-ch

			if insertOAuthClientResult.Err != nil {
				panic(insertOAuthClientResult.Err.Error())
			}

			fmt.Println("client id: " + insertOAuthClientResult.OAuthClient.Id)
			fmt.Println("client secret: " + insertOAuthClientResult.OAuthClient.Secret)

		},
	}

	clientName string
)

func init() {
	rootCmd.AddCommand(generateOAuthClient)
	generateOAuthClient.Flags().StringVar(
		&clientName,
		"name",
		"",
		"name of the client. if not supplied, app name will be used as aprt of the client name",
	)

	//err := viper.BindPFlag("APP_NAME", rootCmd.PersistentFlags().Lookup("appname"))
	//if err != nil {
	//	panic(err)
	//}
}
