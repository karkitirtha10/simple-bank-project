// package main

// import (
// 	"flag"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/golang-migrate/migrate/v4/database/postgres"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/karkitirtha10/simplebank/api/routes"
// 	"github.com/karkitirtha10/simplebank/db"

// 	"github.com/karkitirtha10/simplebank/cmd/commands"
// 	"github.com/karkitirtha10/simplebank/config"
// 	// "github.com/karkitirtha10/simplebank/db"
// )

// // entrypoint to the application
// func main() {
// 	//load configuration values
// 	c, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalln("failed to load environment variables")
// 	}

// 	ConsoleCommands(c)

// 	// flag.Parsed() can be used to check whether flag.Parse() is called already
// 	flag.Parse()

// 	//if no command is supplied, run the web server
// 	//when command is supplied, no of non flag arguments  is >= 2. eg: main.go migration:up
// 	if flag.NArg() < 2 {
// 		WebServer(c)
// 	}

// }

// // run server here
// func WebServer(config config.Config) {

// 	router := gin.Default() //logger and recovery middleware included by default
// 	db := sqlx.MustConnect("postgres", config.DbUrl)

// 	// var user model.User
// 	// err := db.QueryRowx("SELECT * FROM users LIMIT 1").StructScan(&user)
// 	// fmt.Println(err)
// 	// fmt.Println(user)
// 	// return
// 	//migrate -path ./db/migrations -database "postgresql://pgsuperuser:Admin@1@localhost:5432/simplebank?sslmode=disable" up

// 	routes.Register(router, db, config)
// 	router.Run(config.AppPort) //0.0.0.0:8080
// }

// // register  custom commands here
// func ConsoleCommands(c config.Config) {
// 	var dbMigrate db.DBMigration
// 	commands.Register(commands.NewCreateMigrationCommand(c.MigrationUrl, dbMigrate))
// 	commands.Register(commands.NewMigrateUpCommand(dbMigrate, c))
// 	commands.Register(commands.NewMigrateDownCommand(dbMigrate, c))
// 	commands.Register(commands.NewGenerateRSACommand())
// }
