package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/model"
)

type InsertAccountResult struct {
	accountId uint64
	err       error
}

func (handler AccountHandler) Add(c *gin.Context) {
	var accountInput model.AddAccountInput

	//validation. binding error
	if err := c.ShouldBindJSON(&accountInput); err != nil {
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

	ch := make(chan InsertAccountResult)

	go func(ch chan InsertAccountResult) {
		var insertAccountResult InsertAccountResult

		insertAccountResult.accountId,
			insertAccountResult.err = handler.Repository.Create(
			accountInput.Owner,
			float64(0.00),
			accountInput.Currency,
		)
		ch <- insertAccountResult
	}(ch)

	insertAccountResult := <-ch

	if insertAccountResult.err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insertAccountResult.err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully added new account",
		"data":    insertAccountResult.accountId,
	})

}
