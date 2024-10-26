package core

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/moriba-build/ose/ddd/config"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type (
	Option struct {
		Label string
		Value string
	}
	PasswordState string
	Password      struct {
		Value string
		Hash  string
		State PasswordState
	}
	Reference struct {
		Id    string
		Value string
	}
)

const (
	CHANGE  PasswordState = "CHANGE"
	CHANGED PasswordState = "CHANGED"
)

func Duplicate(list []string) error {
	available := make(map[string]bool)
	duplicate := make([]string, 0)

	for i := range list {
		if available[list[i]] == true {
			duplicate = append(duplicate, fmt.Sprintf("%s already selected", list[i]))
		} else {
			available[list[i]] = true
		}
	}

	if len(duplicate) > 0 {
		return fmt.Errorf(strings.Join(duplicate, ", "))
	}

	return nil
}

func GeneratePassword() (*Password, error) {
	pass, err := password.Generate(8, 2, 2, true, false)
	if err != nil {
		return nil, err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return &Password{
		Value: pass,
		Hash:  string(bytes),
		State: CHANGE,
	}, err
}

func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyAccessToken(tokenStr string) (jwt.MapClaims, error) {
	secret := config.NewEnvs().EnvStr("ACCESS_SECRET_KEY")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid login deatils")
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

	return nil, fmt.Errorf("invalid token")
}

func VerifyRefreshToken(tokenStr string) (jwt.MapClaims, error) {
	secret := config.NewEnvs().EnvStr("REFRESH_SECRET_KEY")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid login deatils")
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

	return nil, fmt.Errorf("invalid token")
}
