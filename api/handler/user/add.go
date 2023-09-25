package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/model"
	"github.com/karkitirtha10/simplebank/repositories"
	"golang.org/x/crypto/bcrypt"
)

func (handler UserHandler) Add(c *gin.Context) {
	var userInput model.AddUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create new user",
		})
		return
	}

	ch := make(chan repositories.InsertUserResult)
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
		ch := make(chan repositories.UserResult)
		go handler.UserRepository.FindForEmail(ch, userInput.Email, "u_email")

		if (<-ch).User.Email == userInput.Email {
			c.JSON(http.StatusCreated, gin.H{
				"error": "email already used",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully added new user",
	})
	c.Header("Location", userResult.UserID)

}
