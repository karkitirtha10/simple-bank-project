package handler

import (
	"github.com/karkitirtha10/simplebank/repositories"
)

type UserHandler struct {
	UserRepository repositories.IUserRepository
}
