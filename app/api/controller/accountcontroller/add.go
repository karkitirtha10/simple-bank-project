package handler

import (
	"net/http"

	"github.com/karkitirtha10/simplebank/app/models/inputmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (yo AccountHandler) Add(c *gin.Context) {
	var accountInput inputmodel.AddAccountInput

	//validation. binding systemerror
	if err := c.ShouldBindJSON(&accountInput); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
			//return
			//locales.All
			//v := validator.New()
			//v.RegisterTranslation("required", "{0} is a required field", func(ut validator.TranslationFunc) error {
			//	return ut("{0} is a required field", "Field")
			//}, func(ut validator.TranslationFunc, fe validator.FieldError) string {
			//	return fmt.Sprintf("%s is a required field", fe.Field())
			//})
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"systemerror": err.Error(),
		})

	}

	// fmt.Println(accountInput.Owner, accountInput.Currency)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "successfully added new accountcontroller",
	// 	"data":    accountInput,
	// })

	// accountcontroller := Account{
	// 	Owner:    accountInput.Owner,
	// 	Balance:  0,
	// 	Currency: accountInput.Currency,
	// }

	// query := "INSERT INTO accounts (owner,balance,currency) VALUES (:owner,:balance,:currency) RETURNING id"
	//($1,$2,$3) shoud be used instead of (?,?,?) for DB.QueryRow
	// query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"
	// result, err := controller.DB.NamedExec(query, &accountcontroller)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"systemerror": err.Error(),
	// 	})
	// 	return
	// }

	// query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"
	// var accountId uint64
	// err := controller.DB.QueryRowx(query, accountInput.Owner, float64(0.00), accountInput.Currency).Scan(&accountId)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"systemerror": err.Error(),
	// 	})
	// 	return
	// }
	/////////////

	// accountId, err := result.LastInsertId()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"systemerror": err.Error(),
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
			"systemerror": insertAccountResult.Err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully added new accountcontroller",
		"data":    insertAccountResult.AccountId,
	})

}

// go func(ch chan int) {

// 	ch <- function1()

// }(ch)
