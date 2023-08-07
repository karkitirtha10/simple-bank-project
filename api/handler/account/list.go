package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/model"
)

type GetAccountsResult struct {
	accounts []model.Account
	err      error
}

func (h AccountHandler) List(c *gin.Context) {
	h.Repository.GetAll()

	ch := make(chan GetAccountsResult)

	go func(ch chan GetAccountsResult) {
		var accountsError GetAccountsResult
		accountsError.accounts, accountsError.err = h.Repository.GetAll()
		ch <- accountsError
	}(ch)

	getAccountsResult := <-ch

	if getAccountsResult.err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": getAccountsResult.err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfull",
		"data":    getAccountsResult.accounts,
	})
}
