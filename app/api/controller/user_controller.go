package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/app/systemerror"

	"github.com/go-playground/validator/v10"
	"github.com/karkitirtha10/simplebank/app/localization"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/inputmodel"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/language"
)

type UserControllerInterface interface {
	Add(c *gin.Context)
}

type UserController struct {
	UserRepository repositories.UserRepositoryInterface
	ErrorHandler   systemerror.ErrorHandlerInterface
}

func (handler *UserController) Add(c *gin.Context) {
	localization := localization.NewLocalization(language.Nepali.String())
	var userInput inputmodel.AddUserInput
	//validation
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": localization.Translate("The given data is invalid"),
			"errors": systemerror.ToValidationErrorBag(
				err.(validator.ValidationErrors),
				localization,
				userInput,
			),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"systemerror": "failed to create new usercontroller",
		})
		return
	}

	ch := make(chan datamodel.InsertUserResult)
	go handler.UserRepository.Create(
		ch,
		userInput.Name,
		userInput.Email,
		hashedPassword,
		time.Now(),
		1,
	)
	userResult := <-ch
	// 	if err == sql.ErrNoRows {
	if userResult.Err != nil {

		ch := make(chan datamodel.UserResult)
		go handler.UserRepository.FindForEmail(ch, userInput.Email, "u_email")

		if (<-ch).User.Email == userInput.Email {
			c.JSON(http.StatusCreated, gin.H{
				"systemerror": "email already used",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create usercontroller",
		})
		return
	}

	c.Header("Location", userResult.UserID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully added new usercontroller",
	})

}

func NewUserController(
	userRepository repositories.UserRepositoryInterface,
	errorHandler systemerror.ErrorHandlerInterface,
) UserControllerInterface {
	return &UserController{
		UserRepository: userRepository,
		ErrorHandler:   errorHandler,
	}
}
