package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/repositories"
)

type AuthHandler struct {
	DB             *sqlx.DB
	UserRepository repositories.IUserRepository
}
