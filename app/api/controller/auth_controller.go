package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/karkitirtha10/simplebank/app/handler"
	"github.com/karkitirtha10/simplebank/app/localization"
	"github.com/karkitirtha10/simplebank/app/models/inputmodel"
	"github.com/karkitirtha10/simplebank/app/systemerror"
	"golang.org/x/text/language"
)

type AuthControllerInterface interface {
	Login(c *gin.Context)
}

// omit empty in json . input model
type AuthController struct {
	LoginHanlder handler.LoginHandlerInterface
	ErrorHandler systemerror.ErrorHandlerInterface
}

func (yo *AuthController) Login(c *gin.Context) {
	//todo take localization from user prefernce or input request
	localization := localization.NewLocalization(language.Nepali.String())
	var loginInput inputmodel.LoginInput
	//validation
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": localization.Translate("The given data is invalid"),
			"errors": systemerror.ToValidationErrorBag(
				err.(validator.ValidationErrors),
				localization,
				loginInput,
			),
		})
		return
	}

	handlerResult, err := yo.LoginHanlder.Handle(loginInput)
	if err != nil {
		yo.ErrorHandler.Handle(err, c, localization)
		return
	}

	c.JSON(http.StatusOK, handlerResult)
}

func NewAuthController(
	loginHandler handler.LoginHandlerInterface,
	errorHandler systemerror.ErrorHandlerInterface,
) AuthControllerInterface {
	return &AuthController{
		LoginHanlder: loginHandler,
		ErrorHandler: errorHandler,
	}
}
