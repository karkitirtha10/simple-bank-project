package command

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/config"
	"github.com/spf13/cobra"
)

// usage 1: ./dist/simple-bank-cli generate-oauth-client personal
// usage 2: ./dist/simple-bank-cli generate-oauth-client client --name=khalti-payment-gateway
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
			// oauthClientRepository := repositories.SingleOAuthClientRepository()
			oauthClientRepository := application.OAuthClientRepository
			//millisecond time ordered unique id
			clientId := uuid.Must(uuid.NewV7()).String()

			//unordered unique id (whole id is generated randomly)
			secret := uuid.New().String()
			ocType := enums.OAuthClientTypeShortNameEnum(args[0]).ToOAuthClientTypeEnum()
			if clientName == "" {
				clientName = config.AppName + " " + string(ocType) + " client"
			}

			//thread
			ch := make(chan datamodel.InsertOAuthClientResult)
			go oauthClientRepository.Insert(
				ch,
				clientId,
				clientName,
				secret,
				ocType,
				false,
			) //thread
			insertOAuthClientResult := <-ch

			if insertOAuthClientResult.Err != nil {
				panic(insertOAuthClientResult.Err.Error())
			}

			fmt.Println("successfully generated " + string(ocType) + " oauth client")
			fmt.Println("CLIENT ID: " + insertOAuthClientResult.OAuthClient.Id)
			fmt.Println("CLIENT SECRET: " + insertOAuthClientResult.OAuthClient.Secret)
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
}
