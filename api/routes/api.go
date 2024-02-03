package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/api/handler"
	"github.com/karkitirtha10/simplebank/config"
)

func Register(r *gin.Engine, db *sqlx.DB, config config.Config) {
	//register api routes here
	accountHandler := handler.NewAccountHandler(db)
	r.POST("/api/v1/accounts", accountHandler.Add)
	r.GET("/api/v1/accounts", accountHandler.List)

	//register api routes here
	userHandler := handler.NewUserHandler(db)
	r.POST("/api/v1/users", userHandler.Add)

	authHandler := handler.NewAuthHandler(db)
	r.POST("/api/v1/login", authHandler.Login)
	// r.GET("/api/v1/users", userHandler.List)
}
