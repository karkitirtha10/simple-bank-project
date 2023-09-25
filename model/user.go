package model

import (
	"database/sql"
)

type User struct {
	Id              string       `db:"u_id"`
	Name            string       `db:"u_name"`
	Email           string       `db:"u_email"`
	Password        string       `db:"u_password"`
	EmailVerifiedAt sql.NullTime `db:"u_email_verified_at"`
	Active          uint8        `db:"u_active"`
	Created_at      sql.NullTime `db:"u_created_at"`
	Updated_at      sql.NullTime `db:"u_updated_at"`
	// u_updated_at TIMESTAMPTZ
}

// type User struct {
// 	Id              string    `db:"u_id"`
// 	Name            string    `db:"u_name"`
// 	Email           string    `db:"u_email"`
// 	Password        string    `db:"u_password"`
// 	EmailVerifiedAt time.Time `db:"u_email_verified_at"`
// 	Active          uint8     `db:"u_active"`
// 	Created_at      time.Time `db:"u_created_at"`
// 	Updated_at      time.Time `db:"u_updated_at"`
// 	// u_updated_at TIMESTAMPTZ
// }

//binding:"required"  is equvalent to present validation in laravel
// validate:"required" is equvalent to required validation in laravel

type AddUserInput struct {
	Name     string `json:"name" binding:"required" validate:"required,max:2"`
	Email    string `json:"email" binding:"required" validate:"required,max:2"`
	Password string `json:"password" binding:"required" validate:"required,max:2"`
	// u_updated_at TIMESTAMPTZ
}
