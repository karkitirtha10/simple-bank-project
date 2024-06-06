package systemerror

import (
	"log/slog"

	"github.com/karkitirtha10/simplebank/app/enums"
)

// SystemError has source of error. short tokenId , http status for response, error code, previous error etc
type SystemError struct {
	message      string
	shortMessage string
	errorCode    enums.SystemErrorCode
	httpStatus   int   //http.StatusOK
	previous     error // can be null
	Source       ErrorSource
	doNotLog     bool //default false
}

func (yo *SystemError) Error() string {
	return yo.message
}

func (yo *SystemError) LogCustomError(logger *slog.Logger) {

	logger.Error(
		yo.Error(),
		slog.String("error_code", string(yo.errorCode)),
		slog.Int("http_status", yo.httpStatus),
		slog.Group(
			"source",
			slog.String("file", yo.Source.File),
			slog.String("method", yo.Source.Method),
			slog.Int("line", yo.Source.Line),
		),
	)

}
