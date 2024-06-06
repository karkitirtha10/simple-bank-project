package systemerror

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/karkitirtha10/simplebank/app/enums"
)

// always log
// NotFoundError has source of error. short tokenId , http status for response, error code, previous error etc
type NotFoundError struct {
	Resource  string
	Attribute string
	Previous  error       // can be null
	Message string
	Source    ErrorSource //todo
}

func (yo *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found for %s", yo.Resource, yo.Attribute)
}

func (yo *NotFoundError) LogCustomError(logger *slog.Logger) {

	logger.Error(
		yo.Error(),
		slog.String("error_code", string(enums.NOT_FOUND_ERROR)),
		slog.Int("http_status", http.StatusNotFound),
		slog.Group(
			"source",
			slog.String("file", yo.Source.File),
			slog.String("method", yo.Source.Method),
			slog.Int("line", yo.Source.Line),
		),
	)
}

/*
func NewNotFound(
	messageParts string, //eg user || admin = 0, email = name@company.com
	previous error, // can be null
) *SystemError {
	messageSlice := strings.Split(messageParts, "||")

	return &SystemError{
		message:      fmt.Sprintf("%s not found for %s", messageSlice[0], messageSlice[1]),
		shortMessage: fmt.Sprintf("%s not found", messageSlice[0]),
		previous:     previous,
		Source:       NewSource(2),
	}
}*/
