package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/api/handler"
	"github.com/karkitirtha10/simplebank/config"
)

// /////////////////new code///////////////////////////////
func Register(r *gin.Engine, db *sqlx.DB, config config.Config) {
	//register api routes here
	accountHandler := handler.NewAccountHandler(db)
	r.POST("/api/v1/accounts", accountHandler.Add)
}

///////////////////old code///////////////////////////////

/*
import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/api/handler"
	"github.com/karkitirtha10/simplebank/config"
	//√
	//"github.com/karkitirtha10/gotodo/pkg/common/factory"
)

type Api struct {
	R              *gin.Engine
	DB             *sqlx.DB
	Config         config.Config
	HandlerFactory handler.Factory
	// ServiceFactory services.Factory
	// RepositoryFactory repositories.Factory
}

func (api Api) Register(r *gin.Engine, db *sqlx.DB, config config.Config) {
	//register api routes here

	r.POST("/api/v1/accounts")
}

func NewApi(
	router *gin.Engine,
	db *sqlx.DB,
	config config.Config,
	handlerFactory handler.Factory,
) *Api {
	return &Api{
		R:              router,
		DB:             db,
		Config:         config,
		HandlerFactory: handlerFactory,
	}
}
*/
