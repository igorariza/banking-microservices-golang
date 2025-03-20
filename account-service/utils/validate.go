package utils

import (
	"banking-system/account-service/data/models"
	"errors"
	"regexp"
	"strings"
)

var regexName string = `^[a-z][-_0-9a-z]+`

func ValidateCreateAccount(account models.Account) error {

	stringArray := []string{}

	if account.Balance <= 0 {
		stringArray = append(stringArray, "initial balance cannot be negative")
	}

	if !regexp.MustCompile(regexName).MatchString(account.Name) {
		stringArray = append(stringArray, "account name is invalid (only lowercase letters, numbers, hyphen, and underscore are allowed)")
	}
	if account.Name == "" {
		stringArray = append(stringArray, "account name cannot be empty")
	}

	if len(stringArray) > 0 {
		justString := strings.Join(stringArray, " ")
		return errors.New(justString)
	}
	return nil
}
