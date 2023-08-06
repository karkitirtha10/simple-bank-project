package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/karkitirtha10/simplebank/db"
)

type CreateMigration struct {
	MigrationUrl string
	DBMMigration db.DBMigration
}

func (CreateMigration) Name() string {
	return string("migration:create")
}

func (cMigtn CreateMigration) Handle() {
	flagset := flag.NewFlagSet(cMigtn.Name(), flag.ExitOnError)
	migrationName := flagset.String("migration_name", "", "name of migration file")

	flagset.Parse(os.Args[2:])
	err := cMigtn.DBMMigration.CreateMigration(cMigtn.MigrationUrl, *migrationName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created migration file")
}

func NewCreateMigrationCommand(
	migrationUrl string,
	dbMMigration db.DBMigration,
) *CreateMigration {
	return &CreateMigration{
		MigrationUrl: migrationUrl,
		DBMMigration: dbMMigration,
	}
}
