package repositories

import "github.com/jmoiron/sqlx"

type IAccountRepository interface {
	create(string, float64, string) *sqlx.Row
}

type AccountRepository struct {
	DB *sqlx.DB
}

func (r AccountRepository) create(owner string, balance float64, currency string) *sqlx.Row {
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

	return r.DB.QueryRowx(query, owner, balance, currency)
}
