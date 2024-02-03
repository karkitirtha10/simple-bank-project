package services

// import (
// "github.com/jmoiron/sqlx"
// handler "github.com/karkitirtha10/simplebank/api/handler/account"
// 	"github.com/karkitirtha10/simplebank/config"
// 	"github.com/karkitirtha10/simplebank/repositories"
// )

// /////////////////new code///////////////////////////////
func NewAuthService() *AuthService {
	return &AuthService{NewRSAGenerator()}
}

func NewRSAGenerator() *RSAGenerator {

	return &RSAGenerator{}
}

///////////////////old code///////////////////////////////

/*
type Factory struct {
	// DB                *sqlx.DB
	Config            config.Config
	RepositoryFactory repositories.Factory
}

func (f Factory) NewAuth() IAuth {
	return &Auth{}
}

func (f Factory) NewRSAGenerator() IRSAGenerator {
	return &RSAGenerator{}
}

// instates factory
func NewFactory(config config.Config, repositoryFactory repositories.Factory) *Factory {
	return &Factory{
		Config:            config,
		RepositoryFactory: repositoryFactory,
	}
}
*/
