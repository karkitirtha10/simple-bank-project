package datamodel

import (
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type UserResult struct {
	User dbmodel.User
	Err  error
}
