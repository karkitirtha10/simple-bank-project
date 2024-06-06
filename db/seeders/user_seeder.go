package seeders

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserSeeder struct {
	repositories.UserRepositoryInterface
}

func (yo UserSeeder) Seed() {

	//seed superadmin
	hashedPassword, err := bcrypt.
		GenerateFromPassword(
			[]byte("admin@123"),
			bcrypt.DefaultCost,
		)
	if err != nil {
		panic("could not hash password. Failed to seed UserSeeder")
	}

	superAdminId, err1 := uuid.NewV7()
	if err1 != nil {
		panic("could not generate v7 uuid. Failed to seed UserSeeder")
	}

	superAdmin := dbmodel.User{
		Id:              superAdminId.String(),
		Name:            "superadmin user",
		Email:           "superadmin@hisabkitab.com",
		Password:        string(hashedPassword),
		EmailVerifiedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		Active:          1,
	}

	_, err = yo.UserRepositoryInterface.InsertIfEmailNotExists(superAdmin)
	if err != nil {
		panic("Failed to seed UserSeeder. "+ err.Error())
	}

	//seed admin

	fmt.Println("successfully seeded UserSeeder")
}

func NewUserSeeder(userRepository repositories.UserRepositoryInterface) SeederInterface {
	return &UserSeeder{
		UserRepositoryInterface: userRepository,
	}
}
