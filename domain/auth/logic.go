package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/moriba-build/ose/ddd/config"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"time"
)

func New(id string, school string) (*Domain, error) {
	access, err := AccessToken(id, school)
	if err != nil {
		return nil, err
	}
	refresh, err := RefreshToken(id, school)
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
	school := token["school"].(string)

	access, err := AccessToken(id, school)
	return &Domain{
		access:  access,
		refresh: refresh,
	}, nil
}

func AccessToken(id string, school string) (string, error) {
	value := config.NewEnvs().EnvStr("ACCESS_SECRET_KEY")
	var secret = []byte(value)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["school"] = school
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	return token.SignedString(secret)
}

func ActiveUser(ctx context.Context, key string) *User {
	val := ctx.Value(key)
	if val != nil {
		return val.(*User)
	}

	return &User{
		Id:         "",
		GivenNames: "",
		FamilyName: "",
		Phone:      0,
		Email:      "",
		Role: core.Reference{
			Id:    "",
			Value: "",
		},
		School: "",
		State:  "",
	}
}

func ActiveAccessToken(ctx context.Context) string {
	val := ctx.Value("access")
	if val != nil {
		return val.(string)
	}

	return ""
}

func ActiveRefreshToken(ctx context.Context) string {
	val := ctx.Value("refresh")
	if val != nil {
		return val.(string)
	}

	return ""
}

func RefreshToken(id string, school string) (string, error) {
	value := config.NewEnvs().EnvStr("REFRESH_SECRET_KEY")
	var secret = []byte(value)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["school"] = school
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
