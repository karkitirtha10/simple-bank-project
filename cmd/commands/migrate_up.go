package commands

import (
	"flag"
	"os"

	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/db"
)

type MigrateUp struct {
	DBMMigration db.DBMigration
	Conf         config.Config
}

func (MigrateUp) Name() string {
	return string("migration:up")
}

func (mUp MigrateUp) Handle() {
	flagset := flag.NewFlagSet(mUp.Name(), flag.ExitOnError)
	mStep := flagset.Int("migrate_step", 0, "no of migration")

	flagset.Parse(os.Args[2:])

	// fmt.Println("inside migrate up: ", os.Args[2:])
	mUp.DBMMigration.MigrateUp(*mStep, mUp.Conf)
}

func NewMigrateUpCommand(
	dBMMigration db.DBMigration,
	conf config.Config,
) *MigrateUp {
	return &MigrateUp{
		DBMMigration: dBMMigration,
		Conf:         conf,
	}
}
