package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/models/dbmodel"
)

type InsertAccountResult struct {
	AccountId string
	Err       error
}

type IAccountRepository interface {
	Create(float64, string) InsertAccountResult
	GetAll() ([]dbmodel.Account, error)
}

type AccountRepository struct {
	DB *sqlx.DB
}

func (r AccountRepository) Create(balance float64, currency string) InsertAccountResult {
	// result, err := ctlr.db.NamedExec(query, &account)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"

	// var accountId uint64
	// err := s.db.QueryRowx(query, owner, balance, currency).Scan(&accountId)
	query := "INSERT INTO accounts (ac_balance, ac_currency) VALUES ($1,$2) RETURNING ac_id"

	var accountId string
	err := r.DB.QueryRowx(query, balance, currency).Scan(&accountId)
	return InsertAccountResult{accountId, err}
	// return accountId, err
}

func (r AccountRepository) GetAll() ([]dbmodel.Account, error) {
	var accounts []dbmodel.Account
	query := "SELECT ac_id, ac_balance, ac_currency, ac_created_at FROM accounts"
	err := r.DB.Select(&accounts, query)
	return accounts, err
}
