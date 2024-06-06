package routes

import (
	"github.com/karkitirtha10/simplebank/app"
	"github.com/karkitirtha10/simplebank/app/api/middleware"
)

func Register(app app.Application) {
	//register api routes here
	r := app.Router
	r.POST("/api/v1/login", app.AuthController.Login)

	authRoute := r.Group(
		"auth",
		middleware.AuthMiddleware(
			app.JWTService,
			app.OAuthAccessTokenRepository,
			app.ErrorHandler,
			app.Config,
		),
	)

	authRoute.POST("/api/v1/users", app.UserController.Add)

	/*
		accountHandler := controller.NewAccountHandler(app.DB)
		r.POST("/api/v1/accounts", accountHandler.Add)
		r.GET("/api/v1/accounts", accountHandler.List)


	*/
	// r.GET("/api/v1/users", userHandler.List)
}
