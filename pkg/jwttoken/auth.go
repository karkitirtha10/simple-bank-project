package services

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/karkitirtha10/simplebank/services"
)

// declare
type IAuth interface {
	CreateToken(*rsa.PrivateKey) (string, error)
	ClaimsFromToken(tokenString string, publicKey string) (jwt.Claims, error)
}

type LoadPrivateKeyFileResult struct {
	PrivateKey *rsa.PrivateKey
	err        error
}

// define
type Auth struct {
	RSAGenerator services.IRSAGenerator
}

func (a Auth) CreateToken(claims jwt.MapClaims) (string, error) {
	//   key = /* Load key from somewhere, for example a file SigningMethodRS256*/

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// t := jwt.NewWithClaims(jwt.SigningMethodRS256,
	// 	jwt.MapClaims{
	// 		"expiresIn": "john",
	// 		"id":        2,
	// 	})

	ch := make(chan LoadPrivateKeyFileResult)
	go func(chan LoadPrivateKeyFileResult) {
		var lPKFResult LoadPrivateKeyFileResult
		lPKFResult.PrivateKey, lPKFResult.err = a.RSAGenerator.LoadPrivateKeyFromFile("storage/private.key")
		ch <- lPKFResult
	}(ch)

	loadPrivateKeyFileResult := <-ch

	return t.SignedString(loadPrivateKeyFileResult.PrivateKey)
}

func (a Auth) ClaimsFromToken(tokenString string, publicKey string) (jwt.Claims, error) {
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

// func new() (Auth){
// 	return Auth{}
// }
