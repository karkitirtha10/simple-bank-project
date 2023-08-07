package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/model"
)

type IAccountRepository interface {
	Create(string, float64, string) (uint64, error)
	GetAll() ([]model.Account, error)
}

type AccountRepository struct {
	DB *sqlx.DB
}

func (r AccountRepository) Create(owner string, balance float64, currency string) (uint64, error) {
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
	query := "INSERT INTO accounts (owner,balance,currency) VALUES ($1,$2,$3) RETURNING id"

	var accountId uint64
	err := r.DB.QueryRowx(query, owner, balance, currency).Scan(&accountId)

	return accountId, err
}

func (r AccountRepository) GetAll() ([]model.Account, error) {
	var accounts []model.Account
	query := "SELECT * FROM accounts"
	err := r.DB.Select(&accounts, query)
	return accounts, err
}
