package handler

import (
	"github.com/karkitirtha10/simplebank/models/dbmodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAccountsResult struct {
	accounts []dbmodel.Account
	err      error
}

func (yo AccountHandler) List(c *gin.Context) {

	ch := make(chan GetAccountsResult)

	go func(ch chan GetAccountsResult) {
		var accountsError GetAccountsResult
		accountsError.accounts, accountsError.err = yo.Repository.GetAll()
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
		"message": "successful",
		"data":    getAccountsResult.accounts,
		//"test_const": weekDayString(Sunday),
	})
}

//
//func weekDayString(weekday ~string) ~string {
//	return weekday
//}

// Weekday - Custom type to hold value for weekday ranging from 1-7
//type Weekday string

// Declare related constants for each weekday starting with index 1
//const (
//	Sunday Weekday = "sunday" // EnumIndex = 1
//	Monday Weekday = "monday" // EnumIndex = 2 	// EnumIndex = 7
//)

//// String - Creating common behavior - give the type a String function
//func (w Weekday) String() string {
//	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w-1]
//}
//
//// EnumIndex - Creating common behavior - give the type a EnumIndex function
//func (w Weekday) EnumIndex() int {
//	return int(w)
//}
