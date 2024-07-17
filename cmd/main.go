package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// "github.com/karkitirtha10/simplebank/app"
	"github.com/karkitirtha10/simplebank/app"
	"github.com/karkitirtha10/simplebank/app/api/routes"
)

func main() {
	app := app.InitializeApp()
	router := gin.Default()
	
	routes.Register(app, router)
	err := router.Run(app.Config.AppPort) //0.0.0.0:8080
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
