package systemerror

import (
	"net/http"
)

// donot log error
// todo : attempt may be logged. if need to log as error then log it using separate logic insted of using central errorlogger
type UnAuthenticatedError struct {
	tokenId string //use for future
	Source  ErrorSource
}

func (yo *UnAuthenticatedError) Error() string {
	return http.StatusText(http.StatusUnauthorized)
}
