package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/models/inputmodel"
	"github.com/karkitirtha10/simplebank/repositories"
	"net/http"
)

type AuthHandler struct {
	DB             *sqlx.DB
	UserRepository repositories.IUserRepository
}

func (yo AuthHandler) Login(c *gin.Context) {
	var loginInput inputmodel.LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ch := make(chan repositories.UserResult)
	yo.UserRepository.FindForEmail(ch, loginInput.Email, "*")
	userResult := <-ch
	if userResult.Err != nil {
		panic("something went wrong")
	}

	s //now generate refrsh token

	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginInput.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message":      "successful",
			"access_token": accessToken,
			"user":         user,
		})
		return
	}
}
