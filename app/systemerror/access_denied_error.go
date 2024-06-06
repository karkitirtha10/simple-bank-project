package systemerror

import (
	"net/http"
)

// donot log error
type AccessDeniedError struct {
	message string //use for future
	Source  ErrorSource
}

func (yo *AccessDeniedError) Error() string {
	return http.StatusText(http.StatusForbidden) // or resource
}
