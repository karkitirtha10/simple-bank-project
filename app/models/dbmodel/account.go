package dbmodel

import "time"

// Account db
type Account struct {
	Id string `db:"ac_id"`
	//Owner     string    `db:"ac_owner"`
	Balance   float32   `db:"ac_balance"`
	Currency  string    `db:"ac_currency"`
	CreatedAt time.Time `db:"ac_created_at"`
}

// Balance  float32 `json:"balance"`
