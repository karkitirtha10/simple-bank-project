package systemerror

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	// "github.com/karkitirtha10/simplebank/app/services"
	"github.com/karkitirtha10/simplebank/app/localization"
)

type ErrorRendererInterface interface {
	Response(err error, c *gin.Context, localization localization.ILocalization)
}

type ErrorRenderer struct {
	localization localization.ILocalization
}

func (yo *ErrorRenderer) Response(
	err error,
	c *gin.Context,
	localization localization.ILocalization,
) {
	yo.localization = localization
	switch err.(type) {
	case *NotFoundError:
		notFound := err.(*NotFoundError)
		yo.NotFoundErrorResponse(notFound, c)
	case *SystemError:
		systemError := err.(*SystemError)
		yo.SystemErrorResponse(systemError, c)
	case *UnAuthenticatedError:
		// unAuthError := err.(*UnAuthenticatedError)
		yo.UnAuthenticatedErrorrResponse(c)
	case *AccessDeniedError:
		yo.UnForbiddenErrorrResponse(c)
	case validator.ValidationErrors:
		yo.ValidationErrorResponse(
			err.(validator.ValidationErrors),
			c,
		)
	default:
		yo.UnconvertedErrorResponse(c)
	}
}

func (yo *ErrorRenderer) NotFoundErrorResponse(err *NotFoundError, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("%s not found.", err.Resource),
	})
}

func (yo *ErrorRenderer) SystemErrorResponse(err *SystemError, c *gin.Context) {
	msg := "oops! something went wrong."
	if err.shortMessage == "" {
		msg = err.shortMessage
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": msg,
	})
}

func (yo *ErrorRenderer) UnconvertedErrorResponse(c *gin.Context) {
	// c.JSON(http.StatusInternalServerError, gin.H{
	// 	"message": "oops! something went wrong.",
	// })
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "oops! something went wrong.",
	})
}

func (yo *ErrorRenderer) UnAuthenticatedErrorrResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthenticated",
	})
}

func (yo *ErrorRenderer) UnForbiddenErrorrResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "Unauthorized action.",
	})
}

func (yo *ErrorRenderer) ValidationErrorResponse(v validator.ValidationErrors, c *gin.Context) {

	c.JSON(http.StatusForbidden, gin.H{
		"message": "Unauthorized action.",
		// "errors":  ToValidationErrorBag(v, yo.localization),
		"errors": "asdas",
	})
}
