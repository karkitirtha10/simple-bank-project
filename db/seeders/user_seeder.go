package seeders

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserSeeder struct {
	repositories.UserRepositoryInterface
}

func (yo UserSeeder) Seed() error {

	//seed superadmin
	hashedPassword, err := bcrypt.
		GenerateFromPassword(
			[]byte("admin@123"),
			bcrypt.DefaultCost,
		)
	if err != nil {
		return err
	}

	superAdminId, err1 := uuid.NewV7()
	if err1 != nil {
		return err1
	}

	superAdmin := dbmodel.User{
		Id:       superAdminId.String(),
		Name:     "superadmin user",
		Email:    "superadmin@hisabkitab.com",
		Password: string(hashedPassword),
		EmailVerifiedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Active: 1,
	}

	_, err = yo.UserRepositoryInterface.InsertIfEmailNotExists(superAdmin)

	if err != nil {
		return err
	}

	return nil
}

func NewUserSeeder(userRepository repositories.UserRepositoryInterface) SeederInterface {
	return &UserSeeder{
		UserRepositoryInterface: userRepository,
	}
}
