package command

import (
	"fmt"

	"github.com/karkitirtha10/simplebank/db/seeders"
	"github.com/spf13/cobra"
)

var (
	seederName string

	//    ./dist/simple-bank-cli db-seed --seeder=
	//    ./dist/simple-bank-cli db-seed --seeder=UserSeeder
	//    here UserSeeder is the struct name
	seedDatabase = &cobra.Command{
		Use:   "db-seed",
		Short: "command to seed database",
		Run: func(cmd *cobra.Command, args []string) {
			s := seeders.NewDBSeeder(application, seederName)

			err := s.Seed()
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	// privateKeyPath string

	// publicKeyPath string
)

func init() {
	seedDatabase.Flags().StringVar(&seederName, "seeder", "", "Specify the seeder name to run")
	// Mark the seeder flag as required
	// seedDatabase.MarkFlagRequired("seeder")
	rootCmd.AddCommand(seedDatabase)
}
