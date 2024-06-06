package repositories

import (
	"database/sql"
	"time"

	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (yo UserRepository) FindForEmail(
	ch chan datamodel.UserResult,
	email string,
	cols string,
) {
	var user dbmodel.User
	err := yo.DB.QueryRowx("SELECT "+cols+" FROM users WHERE u_email = $1 LIMIT 1", email).StructScan(&user)
	ch <- datamodel.UserResult{User: user, Err: err}
}

func (yo UserRepository) FindForId(
	ch chan datamodel.UserResult,
	id string,
	cols string,
) {
	var user dbmodel.User
	err := yo.DB.QueryRowx("SELECT "+cols+" FROM users WHERE u_id = $1 LIMIT 1", id).
		StructScan(&user)
	ch <- datamodel.UserResult{User: user, Err: err}
}

// func (yo UserRepository) ExistsForEmail(email string) UserResult {
// 	var usercontroller model.User
// 	err := yo.DB.QueryRowx("select 1 from users where email = $1").StructScan(&usercontroller)
// 	return UserResult{usercontroller, err}
// }

func (yo UserRepository) Create(
	ch chan datamodel.InsertUserResult,
	name string,
	email string,
	password []byte,
	emailVerifiedAt time.Time,
	active uint8,
) {

	var userId string
	err := yo.DB.QueryRowx(
		"INSERT INTO users(u_name, u_email, u_password, u_email_verified_at, u_active) values ($1, $2, $3, $4, $5) RETURNING u_id",
		name,
		email,
		password,
		emailVerifiedAt,
		active,
	).Scan(&userId)

	ch <- datamodel.InsertUserResult{UserID: userId, Err: err}
}

func (yo UserRepository) InsertIfEmailNotExists(
	// ch chan datamodel.InsertOAuthRefreshTokenResult,
	user dbmodel.User,
	// cols string,
) (sql.Result, error) {
	return yo.DB.Exec(
		`
		INSERT INTO users(
			u_id,
			u_name, 
			u_email, 
			u_password, 
			u_email_verified_at, 
			u_active
		) 
		values ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (u_email) DO NOTHING
		`,
		user.Id,
		user.Name,
		user.Email,
		user.Password,
		user.EmailVerifiedAt,
		user.Active,
	)

	// _ = user
	// ch <- datamodel.InsertOAuthRefreshTokenResult{
	// 	OAuthRefreshToken: user,
	// 	Err:               err,
	// }

}

func (yo UserRepository) FindForEmailSync(
	email string,
	cols string,
) *datamodel.UserResult {
	var user dbmodel.User
	err := yo.DB.QueryRowx("SELECT "+cols+" FROM users WHERE u_email = $1 LIMIT 1", email).StructScan(&user)
	return &datamodel.UserResult{User: user, Err: err}
}

type UserRepositoryInterface interface {
	FindForEmail(
		ch chan datamodel.UserResult,
		email string,
		cols string,
	)

	FindForEmailSync(
		email string,
		cols string,
	) *datamodel.UserResult

	Create(
		ch chan datamodel.InsertUserResult,
		name string,
		email string,
		password []byte,
		emailVerifiedAt time.Time,
		active uint8,
	)

	InsertIfEmailNotExists(
		// ch chan datamodel.InsertOAuthRefreshTokenResult,
		user dbmodel.User,
		// cols string,
	) (sql.Result, error)

	FindForId(
		ch chan datamodel.UserResult,
		id string,
		cols string,
	)
}
func NewUserRepository(dbase *sqlx.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: dbase,
	}
}