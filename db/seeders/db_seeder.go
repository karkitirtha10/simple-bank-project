package seeders

import (
	"sync"

	"github.com/karkitirtha10/simplebank/app"
)

type DBSeeder struct {
	App app.Application
}

func (yo DBSeeder) register() []SeederInterface {
	return []SeederInterface{
		NewUserSeeder(yo.App.UserRepository),
		//
		//
	}
}

func (yo DBSeeder) Seed() {
	var wg sync.WaitGroup
	seeders := yo.register()

	var seeder SeederInterface
	for _, seeder = range seeders {
		wg.Add(1)

		go func(w *sync.WaitGroup) {
			defer w.Done()
			seeder.Seed()
		}(&wg)
	}

	wg.Wait()
}

func NewDBSeeder(app app.Application) SeederInterface {
	return &DBSeeder{
		App: app,
	}
}
