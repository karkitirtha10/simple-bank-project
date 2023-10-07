package db

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/karkitirtha10/simplebank/config"

	"github.com/jmoiron/sqlx"
)

func NewDB(config config.Config) *sqlx.DB {
	return sqlx.MustConnect("postgres", config.DbUrl)
}
