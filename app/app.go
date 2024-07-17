package app

import (
	"encoding/json"
	"io"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/api/controller"
	"github.com/karkitirtha10/simplebank/app/handler"

	// handler "github.com/karkitirtha10/simplebank/app/api/controller/accountcontroller"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/app/services"
	"github.com/karkitirtha10/simplebank/app/systemerror"
	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/db"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Application struct {
	Config *config.Config
	DB     *sqlx.DB
	// Router                      *gin.Engine
	I18nBundle                  *i18n.Bundle
	UserRepository              repositories.UserRepositoryInterface
	PostgressDBLogWriter        io.Writer
	ErrorLogger                 systemerror.ErrorLoggerInterface
	ErrorRenderer               systemerror.ErrorRendererInterface
	ErrorHandler                systemerror.ErrorHandlerInterface
	OAuthAccessTokenRepository  repositories.IOAuthAccessTokenRepository
	RSAGenerator                services.IRSAGenerator
	JWTService                  services.IJWTService
	OAuthClientRepository       repositories.IOAuthClientRepository
	OAuthRefreshTokenRepository repositories.IOAuthRefreshTokenRepository
	LoginHandler                handler.LoginHandlerInterface
	UserController              controller.UserControllerInterface
	AuthController              controller.AuthControllerInterface
	PermissionRepository        repositories.PermissionRepositoryInterface
	RoleRepository              repositories.RoleRepositoryInterface
	RolePermissionRepository    repositories.RolePermissionRepositoryInterface
	RolePermissionPersister     services.RolePermissionPersisterInterface
}

func InitializeApp() Application {
	c := config.LoadConfig()
	db := db.Connection(c.DbUrl)
	userRepo := repositories.NewUserRepository(db)
	postgressDBLogWriter := systemerror.NewDBLogWriter(db)
	errorLogger := systemerror.NewErrorLogger(postgressDBLogWriter)
	errorRenderer := systemerror.NewErrorRenderer()
	errorHandler := systemerror.NewErrorHandler(errorLogger, errorRenderer)
	accessTokenRepository := repositories.NewOAuthAccessTokenRepository(db)
	rsaGenerator := services.NewRSAGenerator()
	jwtService := services.NewJWTService(rsaGenerator)
	userController := controller.NewUserController(userRepo, errorHandler)
	oAuthClientRepository := repositories.NewOAuthClientRepository(db)
	oauthRefreshTokenRepository := repositories.NewOAuthRefreshTokenRepository(db)
	oauthRefreshTokenService := services.NewOAuthRefreshTokenService(
		oAuthClientRepository,
		jwtService,
		oauthRefreshTokenRepository,
	)
	personalAccessClientService := services.NewPersonalAccessClientService(
		oAuthClientRepository,
		jwtService,
		accessTokenRepository,
		oauthRefreshTokenService,
		userRepo,
		c,
	)
	loginHandler := handler.NewLoginHandler(userRepo, personalAccessClientService)
	authController := controller.NewAuthController(loginHandler, errorHandler)
	roleRepository := repositories.NewRoleRepository(db)
	permissionRepository := repositories.NewPermissionRepository(db)
	rolePermissionRepository := repositories.NewRolePermissionRepository(db)
	rolePermissionPersister := services.NewRolePermissionPersister(
		permissionRepository,
		roleRepository,
		rolePermissionRepository,
	)
	return Application{
		Config: c,
		// Router:                      gin.Default(),
		DB:                          db,
		I18nBundle:                  InitializeInternationalisation(),
		UserRepository:              userRepo,
		PostgressDBLogWriter:        postgressDBLogWriter,
		ErrorLogger:                 errorLogger,
		ErrorRenderer:               errorRenderer,
		ErrorHandler:                errorHandler,
		OAuthAccessTokenRepository:  accessTokenRepository,
		RSAGenerator:                rsaGenerator,
		JWTService:                  jwtService,
		OAuthClientRepository:       oAuthClientRepository,
		OAuthRefreshTokenRepository: oauthRefreshTokenRepository,
		LoginHandler:                loginHandler,
		UserController:              userController,
		AuthController:              authController,
		RoleRepository:              roleRepository,
		PermissionRepository:        permissionRepository,
		RolePermissionRepository:    rolePermissionRepository,
		RolePermissionPersister:     rolePermissionPersister,
	}
}

func InitializeInternationalisation() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English) //default: english
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("lang/en.json")
	bundle.MustLoadMessageFile("lang/ne.json")
	return bundle
	// ~/.air --build.cmd "go build -o ./dist/simple-bank ./cmd/main.go" --build.bin "./dist/simple-bank"
}
