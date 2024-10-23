package core

import (
	"fmt"
	"strings"
)

type (
	Option struct {
		Label string
		Value string
	}
	Reference struct {
		Id    string
		Value string
	}
)

func Duplicate(list []string) error {
	available := make(map[string]bool)
	duplicate := make([]string, 0)

	for i, _ := range list {
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
