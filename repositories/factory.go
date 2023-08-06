package repositories

import "github.com/jmoiron/sqlx"

// /////////////////new code///////////////////////////////
func NewAccountRepository(db *sqlx.DB) IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

///////////////////old code///////////////////////////////
/*
import (
	"github.com/jmoiron/sqlx"
	// handler "github.com/karkitirtha10/simplebank/api/handler/account"
	// "github.com/karkitirtha10/simplebank/repositories"
	// "github.com/karkitirtha10/simplebank/config"
)

type Factory struct {
	DB *sqlx.DB
	// Config config.Config
}

func (f Factory) NewAccountRepository(db *sqlx.DB) IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (f Factory) NewFactory(db *sqlx.DB) *Factory {
	return &Factory{
		DB: db,
	}
}
*/
