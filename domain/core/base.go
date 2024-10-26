package core

import (
	"fmt"
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
