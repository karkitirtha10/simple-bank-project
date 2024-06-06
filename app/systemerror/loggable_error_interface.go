package systemerror

import "log/slog"

type LoggableError interface {
	Error()
	LogCustomError(logger *slog.Logger)
}
