package systemerror

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/app/localization"
)

type ErrorHandlerInterface interface {
	Handle(
		err error,
		c *gin.Context,
		localization localization.ILocalization,
	)
}

type ErrorHandler struct {
	ErrorLogger   ErrorLoggerInterface
	ErrorRenderer ErrorRendererInterface
}

func (yo *ErrorHandler) Handle(
	err error,
	c *gin.Context,
	localization localization.ILocalization,
) {
	yo.ErrorLogger.Log(err)
	yo.ErrorRenderer.Response(err, c, localization)
}
