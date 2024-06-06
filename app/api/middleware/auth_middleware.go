package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/localization"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/app/services"
	"github.com/karkitirtha10/simplebank/app/systemerror"
	"github.com/karkitirtha10/simplebank/config"
	"golang.org/x/text/language"
)

func AuthMiddleware(
	jwtService services.IJWTService,
	accessTokenRepo repositories.IOAuthAccessTokenRepository,
	errorHandler systemerror.ErrorHandlerInterface,
	c *config.Config,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		localization := localization.NewLocalization(language.Nepali.String())
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": localization.Translate("unauthenticated"),
			})
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")

		claims, err := jwtService.ClaimsFromToken(token, c.PrivateKeyPath)
		if err != nil {
			convertedErr := systemerror.FromPrevious(
				localization.Translate("unauthenticated"),
				enums.UNAUTHORIZED_ERROR,
				http.StatusUnauthorized,
				err,
			)
			errorHandler.Handle(convertedErr, ctx, localization)
		}

		userId, userIdErr := claims.GetSubject()
		if userIdErr != nil {
			convertedErr := systemerror.FromPrevious(
				localization.Translate("unauthenticated"),
				enums.UNAUTHORIZED_ERROR,
				http.StatusUnauthorized,
				err,
			)
			errorHandler.Handle(convertedErr, ctx, localization)
		}

		tokenId, ok := claims["jti"].(string)
		if !ok {
			convertedErr := systemerror.NewSystemError(
				"could not cast jti claims to string",
				localization.Translate("unauthenticated"),
				enums.UNAUTHORIZED_ERROR,
				http.StatusUnauthorized,
			)
			errorHandler.Handle(convertedErr, ctx, localization)

			// errorLogger.Log(errors.New("could not cast jti claims to string"))
			// ctx.JSON(http.StatusUnauthorized, map[string]string{
			// 	"message": "unauthenticated",
			// })
		}

		ch := make(chan datamodel.OAuthAccessTokenResult)
		accessTokenRepo.FindUnRevokedAndUnExpiredForTokenIdAndUserId(
			ch,
			tokenId,
			userId,
			"*",
		)

		accessTokenResult := <-ch
		if accessTokenResult.Error != nil {
			convertedErr := systemerror.FromPrevious(
				localization.Translate("unauthenticated"),
				enums.UNAUTHORIZED_ERROR,
				http.StatusUnauthorized,
				err,
			)
			errorHandler.Handle(convertedErr, ctx, localization)
		}

		if accessTokenResult.OAuthAccessToken.Revoked {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": localization.Translate("unauthenticated"),
			})
		}

		ctx.Set("access_token", accessTokenResult.OAuthAccessToken)
		ctx.Next()

		//TODO: convert to custom system errors and haldle using error handler
		// claims, ok := token.Claims.(jwt.MapClaims)
		// claims.()
		// claims.GetExpirationTime()

	}
}
