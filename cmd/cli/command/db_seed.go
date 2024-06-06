package command

import (
	"github.com/karkitirtha10/simplebank/db/seeders"
	"github.com/spf13/cobra"
)

var (
	seedDatabase = &cobra.Command{
		Use:   "db-seed",
		Short: "command to seed database",
		Run: func(cmd *cobra.Command, args []string) {
			s := seeders.NewDBSeeder(application)
			s.Seed()
		},
	}

	// privateKeyPath string

	// publicKeyPath string
)

func init() {
	rootCmd.AddCommand(seedDatabase)
}
