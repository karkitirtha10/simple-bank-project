package services

func NewAuthService() IOAuthService {
	return &JWTService{NewRSAGenerator()}
}

func NewRSAGenerator() IRSAGenerator {

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
