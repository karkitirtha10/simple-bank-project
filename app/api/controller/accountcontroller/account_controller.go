package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/repositories"
)

type AccountHandler struct {
	DB         *sqlx.DB
	Repository repositories.IAccountRepository
}

// func NewAccountHandler(
// 	db *sqlx.DB,
// 	accountRepository repositories.IAccountRepository,
// ) {

// }
