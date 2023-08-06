package model

import "time"

// db
type Account struct {
	Id        int64     `db:"id"`
	Owner     string    `db:"owner"`
	Balance   float32   `db:"balance"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
}

// input
type AddAccountInput struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

// Balance  float32 `json:"balance"`
