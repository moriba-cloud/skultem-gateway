package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/moriba-build/ose/ddd/config"
	"time"
)

func New(id string) (*Domain, error) {
	access, err := AccessToken(id)
	if err != nil {
		return nil, err
	}
	refresh, err := RefreshToken(id)
	if err != nil {
		return nil, err
	}

	return &Domain{
		access:  access,
		refresh: refresh,
	}, nil
}

func Existing(refresh string) (*Domain, error) {
	token, err := VerifyRefreshToken(refresh)
	if err != nil {
		return nil, err
	}

	id := token["id"].(string)

	access, err := AccessToken(id)
	return &Domain{
		access:  access,
		refresh: refresh,
	}, nil
}

func AccessToken(id string) (string, error) {
	value := config.NewEnvs().EnvStr("ACCESS_SECRET_KEY")
	var secret = []byte(value)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	return token.SignedString(secret)
}

func ActiveUser(ctx context.Context, key string) *User {
	val := ctx.Value(key)
	if val != nil {
		return val.(*User)
	}

	return nil
}

func RefreshToken(id string) (string, error) {
	value := config.NewEnvs().EnvStr("REFRESH_SECRET_KEY")
	var secret = []byte(value)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	return token.SignedString(secret)
}

func VerifyAccessToken(tokenStr string) (jwt.MapClaims, error) {
	value := config.NewEnvs().EnvStr("ACCESS_SECRET_KEY")
	var secret = []byte(value)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}

		return secret, nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("not authorized")
}

func VerifyRefreshToken(tokenStr string) (jwt.MapClaims, error) {
	value := config.NewEnvs().EnvStr("REFRESH_SECRET_KEY")
	var secret = []byte(value)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}

		return secret, nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("not authorized")
}
