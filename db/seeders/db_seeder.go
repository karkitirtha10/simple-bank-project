package seeders

import (
	"fmt"
	"reflect"

	"github.com/karkitirtha10/simplebank/app"
)

type DBSeeder struct {
	App        app.Application
	seederName string
}

func (yo DBSeeder) register() []SeederInterface {
	return []SeederInterface{
		NewUserSeeder(yo.App.UserRepository),
		NewPermissionSeeder(yo.App.RolePermissionPersister),
		//
		//
	}
}

func (yo DBSeeder) Seed() error {
	errch := make(chan error)
	seeders := yo.register()
	var seeder SeederInterface

	for _, seeder = range seeders {
		err := yo.seedOne(seeder, errch)
		if err != nil {
			return err
		}
	}

	return nil

	// var wg sync.WaitGroup
	// seeders := yo.register()

	// var seeder SeederInterface
	// for _, seeder = range seeders {
	// 	wg.Add(1)

	// 	go func(w *sync.WaitGroup) {
	// 		defer w.Done()
	// 		seeder.Seed()
	// 	}(&wg)
	// }

	// wg.Wait()
}

func (yo DBSeeder) seedOne(seeder SeederInterface, errch chan error) error {

	//since SeederInterface is a pointer, we need additional method .Elem()
	reflectSeederName := reflect.TypeOf(seeder).Elem().Name()

	if yo.seederName == "" || yo.seederName == reflectSeederName {
		go func() {
			fmt.Println("seeding  " + reflectSeederName)
			errch <- seeder.Seed()
			fmt.Println("successfully seeded  " + reflectSeederName)
		}()

		return <-errch
	}
	return nil
}

func NewDBSeeder(
	app app.Application,
	seederName string,
) SeederInterface {
	return &DBSeeder{
		App:        app,
		seederName: seederName,
	}
}
