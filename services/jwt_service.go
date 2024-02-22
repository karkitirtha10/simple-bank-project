package services

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	RSAGenerator IRSAGenerator
}

func (a JWTService) CreateToken(claims jwt.MapClaims, privateKeyPath string) (string, error) {
	//   key = /* Load key from somewhere, for example a file SigningMethodRS256*/

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// jwt.MapClaims{
	// 	"expiresIn": "john",
	// 	"id":        2,
	// }

	ch := make(chan LoadPrivateKeyFileResult)
	go func(chan LoadPrivateKeyFileResult) {
		var lPKFResult LoadPrivateKeyFileResult
		lPKFResult.PrivateKey, lPKFResult.err = a.RSAGenerator.LoadPrivateKeyFromFile(privateKeyPath) //eg: "storage/private.key"
		ch <- lPKFResult
	}(ch)

	loadPrivateKeyFileResult := <-ch

	return t.SignedString(loadPrivateKeyFileResult.PrivateKey)
}

func (a JWTService) ClaimsFromToken(tokenString string, publicKey string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// _, ok := t.Method.(*jwt.SigningMethodRS256)
		_, ok := t.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.New("could not parse token")
		}
		return publicKey, nil
	})
	/////////
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Provide the secret key or the public key to verify the token's signature
	// 	return []byte("your-secret-key"), nil
	// })

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// declare
type IOAuthService interface {
	CreateToken(claims jwt.MapClaims, privateKeyPath string) (string, error)
	ClaimsFromToken(tokenString string, publicKey string) (jwt.Claims, error)
}

type LoadPrivateKeyFileResult struct {
	PrivateKey *rsa.PrivateKey
	err        error
}
