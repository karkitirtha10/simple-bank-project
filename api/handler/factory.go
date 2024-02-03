package handler

import (
	"github.com/jmoiron/sqlx"
	accountHandler "github.com/karkitirtha10/simplebank/api/handler/account"
	authHandler "github.com/karkitirtha10/simplebank/api/handler/auth"
	usrhandler "github.com/karkitirtha10/simplebank/api/handler/user"

	//"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/repositories"
	//"github.com/karkitirtha10/simplebank/services"
)

// /////////////////new code///////////////////////////////
func NewAccountHandler(db *sqlx.DB) *accountHandler.AccountHandler {
	return &accountHandler.AccountHandler{
		DB:         db,
		Repository: repositories.NewAccountRepository(db),
	}
}

func NewUserHandler(db *sqlx.DB) *usrhandler.UserHandler {
	return &usrhandler.UserHandler{
		UserRepository: repositories.NewUserRepository(db),
	}
}

func NewAuthHandler(db *sqlx.DB) *authHandler.AuthHandler {
	return &authHandler.AuthHandler{
		DB:             db,
		UserRepository: repositories.NewUserRepository(db),
	}
}

///////////////////old code///////////////////////////////

/*
import (
	"github.com/jmoiron/sqlx"
	handler "github.com/karkitirtha10/simplebank/api/handler/account"
	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/repositories"
	"github.com/karkitirtha10/simplebank/services"
)

type Factory struct {
	DB                *sqlx.DB
	Config            config.Config
	ServiceFactory    services.Factory
	RepositoryFactory repositories.Factory
}

func newAccountHandler(db *sqlx.DB, repository repositories.IAccountRepository) *handler.AccountHandler {
	// return &handler.AccountHandler {
	// 	db: db,
	// 	repository: api.Factory.GetAccountService(db),
	// }
	return &handler.AccountHandler{
		DB:         db,
		Repository: repository,
	}
}
*/
