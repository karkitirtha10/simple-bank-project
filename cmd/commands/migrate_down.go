package commands

import (
	"flag"
	"os"

	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/db"
)

type MigrateDown struct {
	DBMMigration db.DBMigration
	Conf         config.Config
}

func (MigrateDown) Name() string {
	return string("migration:down")
}

func (mDown MigrateDown) Handle() {
	flagset := flag.NewFlagSet(mDown.Name(), flag.ExitOnError)
	mStep := flagset.Int("migrate_step", 0, "no of migration")

	flagset.Parse(os.Args[2:])
	mDown.DBMMigration.MigrateDown(*mStep, mDown.Conf)
}

func NewMigrateDownCommand(
	dBMMigration db.DBMigration,
	conf config.Config,
) *MigrateDown {
	return &MigrateDown{
		DBMMigration: dBMMigration,
		Conf:         conf,
	}
}
