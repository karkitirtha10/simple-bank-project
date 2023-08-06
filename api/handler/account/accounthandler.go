package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/repositories"
)

type AccountHandler struct {
	DB         *sqlx.DB
	Repository repositories.IAccountRepository
}
