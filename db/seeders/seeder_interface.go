package seeders

type SeederInterface interface {
	Seed() error
	// Name() string
}
