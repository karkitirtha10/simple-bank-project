package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/model"
)

func (ctlr AccountHandler) Add(c *gin.Context) {
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
	query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"
	// result, err := ctlr.DB.NamedExec(query, &account)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	var accountId uint64
	err := ctlr.DB.QueryRowx(query, accountInput.Owner, float64(0.00), accountInput.Currency).Scan(&accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	/////////////

	// accountId, err := result.LastInsertId()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully added new account",
		"data":    accountId,
	})

}
