package systemerror

import (
	"log/slog"

	"github.com/karkitirtha10/simplebank/app/enums"
)

// always log
type BasicError struct {
	Message   string
	ErrorCode enums.SystemErrorCode
	Source    ErrorSource //todo
}

func (yo *BasicError) Error() string {
	return yo.Message
}

func (yo *BasicError) LogCustomError(logger *slog.Logger) {
	logger.Error(
		yo.Error(),
		slog.String("error_code", string(enums.NOT_FOUND_ERROR)),
		slog.Int("http_status", int(0)),
		slog.Group(
			"source",
			slog.String("file", yo.Source.File),
			slog.String("method", yo.Source.Method),
			slog.Int("line", yo.Source.Line),
		),
	)
}

func NewBasicError(
	message string,
	errorCode enums.SystemErrorCode,
) *BasicError {
	return &BasicError{
		Message:   message,
		ErrorCode: errorCode,
		Source:    NewSource(2),
	}
}
