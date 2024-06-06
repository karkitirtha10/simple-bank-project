package handler

import (
	"database/sql"
	"errors"

	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/inputmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/app/services"
	"github.com/karkitirtha10/simplebank/app/systemerror"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	UserRepository              repositories.UserRepositoryInterface
	PersonalAccessClientService services.PersonalAccessClientServiceInterface
}

func (yo LoginHandler) Handle(loginInput inputmodel.LoginInput) (*datamodel.LoginHandlerResult, error) {

	ch := make(chan datamodel.UserResult)
	go yo.UserRepository.FindForEmail(ch, loginInput.Email, "*")
	userResult := <-ch

	//todo: add error handler except sql.ErrNoRows
	if errors.Is(userResult.Err, sql.ErrNoRows) {
		return nil, systemerror.NewNotFound(
			"user",
			"email = "+loginInput.Email,
			userResult.Err,
		)
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(userResult.User.Password),
		[]byte(loginInput.Password),
	)

	if err != nil {
		return nil, err
	}

	// return &datamodel.LoginHandlerResult{User: userResult.User}, nil
	oAuthTokenPair, err := yo.PersonalAccessClientService.Generate(userResult.User)
	if err != nil {
		return nil, err
	}

	return &datamodel.LoginHandlerResult{
		OAuthTokenPair: oAuthTokenPair,
		User:           userResult.User,
	}, nil

	//now generate refrsh token
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginInput.Password), bcrypt.DefaultCost)

	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message":        "successful",
	//		"token":          oAuthTokenPair,
	//		"usercontroller": userResult.User,
	//	})
	//	return
	//}
}

func NewLoginHandler(
	userRepository repositories.UserRepositoryInterface, personalAccessClientService services.PersonalAccessClientServiceInterface,
) LoginHandlerInterface {
	return &LoginHandler{
		UserRepository:              userRepository,
		PersonalAccessClientService: personalAccessClientService,
	}
}
