package repositories

import (
	"github.com/karkitirtha10/simplebank/models/dbmodel"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (repo UserRepository) FindForEmail(ch chan UserResult, email string, cols string) {
	var user dbmodel.User
	err := repo.DB.QueryRowx("SELECT "+cols+" FROM users WHERE u_email = $1 LIMIT 1", email).StructScan(&user)
	ch <- UserResult{user, err}
}

// func (repo UserRepository) ExistsForEmail(email string) UserResult {
// 	var user model.User
// 	err := repo.DB.QueryRowx("select 1 from users where email = $1").StructScan(&user)
// 	return UserResult{user, err}
// }

func (repo UserRepository) Create(
	ch chan InsertUserResult,
	name string,
	email string,
	password []byte,
	emailVerifiedAt time.Time,
	active uint8,
) {

	var userId string
	err := repo.DB.QueryRowx(
		"INSERT INTO users(u_name, u_email, u_password, u_email_verified_at, u_active) values ($1, $2, $3, $4, $5) RETURNING u_id",
		name,
		email,
		password,
		emailVerifiedAt,
		active,
	).Scan(&userId)

	ch <- InsertUserResult{userId, err}
}

type IUserRepository interface {
	FindForEmail(ch chan UserResult, email string, cols string)

	Create(
		ch chan InsertUserResult,
		name string,
		email string,
		password []byte,
		emailVerifiedAt time.Time,
		active uint8,
	)
}

// user result set
type InsertUserResult struct {
	UserID string
	Err    error
}

type UserResult struct {
	User dbmodel.User
	Err  error
}

// user repo
