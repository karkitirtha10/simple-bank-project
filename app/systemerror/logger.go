package systemerror

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	// "github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/enums"
	// "github.com/karkitirtha10/simplebank/db"
)

type ErrorLoggerInterface interface {
	Log(err error)
}

type ErrorLogger struct {
	// *sqlx.DB
	PostgressDBLogWriter io.Writer
	// *config.Config
}

// var donotLogErrors []Type =

func (yo ErrorLogger) Log(err error) {

	switch err.(type) {
	case *NotFoundError:
		notFound := err.(*NotFoundError)
		yo.LogError(notFound.LogCustomError)
	case *SystemError:
		systemError := err.(*SystemError)
		if !systemError.doNotLog {
			yo.LogError(systemError.LogCustomError)
		}
	default:
		yo.LogUnconvertedError(err)
	}
}

func (yo ErrorLogger) LogError(logFunc func(*slog.Logger)) {
	fileName := "storage/app_log/app_log_" + time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return
	}

	defer logFile.Close()

	fmt.Println("1")
	// dbLogWriter := NewPostgressDBLogWriter(db.SingleDB())
	// multiWriter := io.MultiWriter(logFile, dbLogWriter)
	multiWriter := io.MultiWriter(logFile, yo.PostgressDBLogWriter)

	logger := slog.New(slog.NewJSONHandler(multiWriter, nil))
	logFunc(logger)
}

func (yo ErrorLogger) LogUnconvertedError(error error) {
	fileName := "storage/app_log/app_log_" + time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return
	}

	defer logFile.Close()

	// dbLogWriter := NewPostgressDBLogWriter(yo.DB)
	multiWriter := io.MultiWriter(logFile, yo.PostgressDBLogWriter)

	logger := slog.New(slog.NewJSONHandler(multiWriter, nil))

	source := NewSource(1)
	logger.Error(
		error.Error(),
		slog.String("error_code", string(enums.GENERAL_SERVER_ERROR)),
		slog.Int("http_status", http.StatusInternalServerError),
		slog.Group(
			"source",
			slog.String("file", source.File),
			slog.String("method", source.Method),
			slog.Int("line", source.Line),
		),
	)

}

/*
func LogError(notFoundErr *NotFoundError) error {
	fileName := "app_log" + time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer logFile.Close()

	dbLogWriter := NewPostgressDBLogWriter(db.SingleDB())
	multiWriter := io.MultiWriter(logFile, dbLogWriter)

	logger := slog.New(slog.NewJSONHandler(multiWriter, nil))

	logger.Error(
		notFoundErr.Error(),
		slog.String("error_code", string(enums.NOT_FOUND_ERROR)),
		slog.Int("http_status", http.StatusNotFound),
		slog.Group(
			"source",
			slog.String("file", notFoundErr.Source.File),
			slog.String("method", notFoundErr.Source.Method),
			slog.Int("line", notFoundErr.Source.Line),
		),
	)
	return nil
}
*/
