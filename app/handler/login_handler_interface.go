package handler

import (
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/inputmodel"
)

type LoginHandlerInterface interface {
	Handle(loginInput inputmodel.LoginInput) (*datamodel.LoginHandlerResult, error)
}
