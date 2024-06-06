package systemerror

import (
	"io"
	"net/http"
	"runtime"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/enums"
)


// in multitenant system db is created in every request. db will be scoped service. use SingleErrorHandler instead of SingleErrorHandler


func NewErrorHandler(
	errorLogger ErrorLoggerInterface,
	errorRenderer ErrorRendererInterface,
) ErrorHandlerInterface {
	return &ErrorHandler{
		ErrorLogger:   errorLogger,
		ErrorRenderer: errorRenderer,
	}
}

func NewDBLogWriter(db *sqlx.DB) io.Writer {
	return PostgressDBLogWriter{db}
}

func NewErrorLogger(dbWriter io.Writer) ErrorLoggerInterface {
	return &ErrorLogger{
		PostgressDBLogWriter: dbWriter,
	}
}

func NewErrorRenderer() ErrorRendererInterface {
	return &ErrorRenderer{}
}

//////////////////////////////////

func NewAccessDeniedError() *AccessDeniedError {
	return &AccessDeniedError{Source: NewSource(2)}
}

func NewNotFound(
	resource string, //eg: user
	attribute string, //eg: || admin = 0, email = name@company.com
	previous error, // can be null
) *NotFoundError {
	return &NotFoundError{
		Resource:  resource,
		Attribute: attribute,
		Previous:  previous,
		Source:    NewSource(2),
	}
}

func NewNotFoundWithMessage(
	resource string, //eg: user
	attribute string, //eg: || admin = 0, email = name@company.com
	previous error, // can be null
) *NotFoundError {
	return &NotFoundError{
		Resource:  resource,
		Attribute: attribute,
		Previous:  previous,
		Source:    NewSource(2),
	}
}

func NewUnAuthenticated() *UnAuthenticatedError {
	return &UnAuthenticatedError{Source: NewSource(2)}
}

//start system errors

func FromPrevious(
	shortMessage string,
	errorCode enums.SystemErrorCode,
	httpStatus int, //http.StatusOK
	previous error, // can be null
) *SystemError {
	var message string
	if previous != nil {
		message = previous.Error()
	}
	return &SystemError{
		message:      message,
		shortMessage: shortMessage,
		errorCode:    errorCode,
		httpStatus:   httpStatus,
		previous:     previous,
		Source:       NewSource(2),
	}
}

func NewError(
	message string,
	shortMessage string,
	errorCode enums.SystemErrorCode,
	httpStatus int, //http.StatusOK
	previous error, // can be null
) *SystemError {
	return &SystemError{
		message:      message,
		shortMessage: shortMessage,
		errorCode:    errorCode,
		httpStatus:   httpStatus,
		previous:     previous,
		Source:       NewSource(2),
	}
}

func NewSystemError(
	message string,
	shortMessage string,
	errorCode enums.SystemErrorCode,
	httpStatus int, //http.StatusOK
) *SystemError {
	return &SystemError{
		message:      message,
		shortMessage: shortMessage,
		errorCode:    errorCode,
		httpStatus:   httpStatus,
		previous:     nil,
		Source:       NewSource(2),
	}
}

func NewClientError(
	message string,
	doNotLog bool,
	previous error, // can be null
) *SystemError {
	return &SystemError{
		message:      message,
		shortMessage: "oops! something went wrong",
		errorCode:    enums.GENERAL_CLIENT_ERROR,
		httpStatus:   http.StatusBadRequest,
		previous:     previous,
		doNotLog:     doNotLog,
		Source:       NewSource(2),
	}
}

func NewServerError(
	message string,
	previous error, // can be null
) *SystemError {
	return &SystemError{
		message:      message,
		shortMessage: "oops! something went wrong",
		errorCode:    enums.GENERAL_SERVER_ERROR,
		httpStatus:   http.StatusInternalServerError,
		previous:     previous,
		Source:       NewSource(2),
	}
}

func FromHttpStatus(
	message string,
	httpStatus int,
	doNotLog bool,
	previous error, // can be null
) *SystemError {
	return &SystemError{
		message:      message,
		shortMessage: http.StatusText(httpStatus),
		errorCode:    enums.GENERAL_SERVER_ERROR,
		httpStatus:   httpStatus,
		previous:     previous,
		doNotLog:     doNotLog,
		Source:       NewSource(2),
	}
}

func NewSource(skip int) ErrorSource {
	pc, file, line, ok := runtime.Caller(skip)
	var method string
	if ok {
		method = runtime.FuncForPC(pc).Name()
	}

	return ErrorSource{
		File:   file,
		Line:   line,
		Method: method,
	}
}

// func NewValidationError(original validator.ValidationErrors) *ValidationError {
// 	return &ValidationError{original: original, source: NewSource(2)}
// }

// func NewErrorHandler(dbWriter io.Writer, localization services.ILocalization) ErrorHandlerInterface {
// 	return &ErrorHandler{
// 		ErrorLogger:   NewErrorLogger(dbWriter),
// 		ErrorRenderer: NewErrorRenderer(localization),
// 	}
// }
