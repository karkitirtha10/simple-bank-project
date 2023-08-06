package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// "github.com/spf13/viper"
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/api/routes"
	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/db"
	"github.com/karkitirtha10/simplebank/services"
)

func maincopy() {

	//commandline flags
	migrationName := flag.String("migration_name", "", "name of migration file")
	mf := flag.String("migrate_flag", "", "migrate flag for migration relate command using go/migrate")
	mStep := flag.Int("migrate_step", 0, "no of migration")
	flag.Parse()

	//load configuration values
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed to load environment variables")
	}

	/*
		viper.SetConfigFile("./pkg/common/envs/.env")
		viper.ReadInConfig()

		//.(string) is a type assertions .
		//type assertion is a way to extract a value of a specific type from an interface type, here viper.Get("DB_URL") returns a interface{}
		//It allows you to check whether an interface value holds a specific underlying type and retrieve that underlying value if the assertion is successful.
		//throws error if the value is not string
		// port := viper.Get("PORT").(string)
		dbUrl := viper.Get("DB_URL").(string)
		migrationsDir := viper.Get("MIGRATION_URL").(string)
	*/

	var dbMigrate db.DBMigration
	fmt.Println(*mf)

	//webserver
	if *mf == "" {
		WebServer(c)
	}

	//migration commands
	/* todo: handle excess step error of file not found */
	if *mf == "up" {
		dbMigrate.MigrateUp(*mStep, c)
	} else if *mf == "down" {
		dbMigrate.MigrateDown(*mStep, c)
	} else if *mf == "rsa" {
		rsa := services.RSAGeneartor{}
		rsa.Generate(2048)
	} else if *mf == "create" {
		err := dbMigrate.CreateMigration(c.MigrationUrl, *migrationName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created migration file")
	}

}

func WebServercopy(config config.Config) {

	router := gin.Default() //logger and recovery middleware included by default
	db := sqlx.MustConnect("postgres", config.DbUrl)

	routes.Register(router, db, config)
	router.Run(config.AppPort) //0.0.0.0:8080
}

/*
func WebServer(config config.Config) {

	router := gin.Default() //logger and recovery middleware included by default
	db := sqlx.MustConnect("postgres", config.DbUrl)

	// accountsApi := accounts.Api{
	// 	Factory: factory.Factory{},
	// }
	//register api handlers
	accountsApi := account.NewApi(factory.Factory{})
	//register api handlers
	accountsApi.Handle(router, db, config)

	router.Run(config.AppPort) //0.0.0.0:8080
}
*/
