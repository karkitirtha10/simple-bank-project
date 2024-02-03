package handler

import (
	"github.com/karkitirtha10/simplebank/models/inputmodel"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/karkitirtha10/simplebank/repositories"
)

func (yo AccountHandler) Add(c *gin.Context) {
	var accountInput inputmodel.AddAccountInput

	//validation. binding error
	if err := c.ShouldBindJSON(&accountInput); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fmt.Println(accountInput.Owner, accountInput.Currency)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "successfully added new account",
	// 	"data":    accountInput,
	// })

	// account := Account{
	// 	Owner:    accountInput.Owner,
	// 	Balance:  0,
	// 	Currency: accountInput.Currency,
	// }

	// query := "INSERT INTO accounts (owner,balance,currency) VALUES (:owner,:balance,:currency) RETURNING id"
	//($1,$2,$3) shoud be used instead of (?,?,?) for DB.QueryRow
	// query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"
	// result, err := handler.DB.NamedExec(query, &account)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"
	// var accountId uint64
	// err := handler.DB.QueryRowx(query, accountInput.Owner, float64(0.00), accountInput.Currency).Scan(&accountId)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	/////////////

	// accountId, err := result.LastInsertId()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	ch := make(chan repositories.InsertAccountResult)

	go func(ch chan repositories.InsertAccountResult) {
		ch <- yo.Repository.Create(
			//accountInput.Owner,
			float64(0.00),
			accountInput.Currency,
		)
	}(ch)

	insertAccountResult := <-ch

	if insertAccountResult.Err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insertAccountResult.Err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully added new account",
		"data":    insertAccountResult.AccountId,
	})

}

// go func(ch chan int) {

// 	ch <- function1()

// }(ch)
