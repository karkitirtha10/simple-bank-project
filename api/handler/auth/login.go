package handler

import (
	"github.com/karkitirtha10/simplebank/models/inputmodel"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/repositories"
)

func (handler AuthHandler) Login(c *gin.Context) {
	var loginInput inputmodel.LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ch := make(chan repositories.UserResult)
	handler.UserRepository.FindForEmail(ch, loginInput.Email, "*")
	//user, err := <-ch
	//_,_ := <-ch

	//if err != nil {
	//
	//}

	// go func(ch chan repositories.UserResult) {

	// }(ch)

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginInput.Password), bcrypt.DefaultCost)
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": "failed to create new user",
	//		})
	//		return
	//	}
}
