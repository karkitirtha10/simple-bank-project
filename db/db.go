package db

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/config"
)

// var dbConnection *sqlx.DB

// func init() {
// 	dbConnection = Connection(config.SingleConfig().DbUrl)
// }

// func SingleDB() *sqlx.DB {
// 	return dbConnection
// }

func Connection(dataSource string) *sqlx.DB {
	return sqlx.MustConnect("postgres", dataSource)
}

func NewConnectionFromConfig() *sqlx.DB {
	c := config.LoadConfig()
	//if err != nil {
	//	log.Fatalln("failed to load environment variables")
	//}
	return sqlx.MustConnect("postgres", c.DbUrl)
}

func TenantConnection(db *sqlx.DB) *sqlx.DB {
	/*
	* todo:-
	* get tenant datasource parts like tenant database name, credentials and other options.
	*
	* datasource parts can be also stored in:
	* 1.  main database.
	* 2.  form of yml, dictionary or in env.
	*
	* recommended approach : database
	*    dynamic : tenant can be added dynamically and tenant database can be created from *    ui or command
	*    vs staticly update the data structure (yaml, dictionary) manually for each tenant
	 */
	tenantDataSource := ""
	return Connection(tenantDataSource)
}
