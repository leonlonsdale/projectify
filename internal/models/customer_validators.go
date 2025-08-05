package models

import (
	"regexp"
	"strings"
)

const MinPasswordLength = 8

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (c *CustomerRegistration) Validate() map[string]string {
	validationErrors := make(map[string]string)

	if strings.TrimSpace(c.Name) == "" {
		validationErrors["name"] = "name is required"
	}

	if strings.TrimSpace(c.Email) == "" {
		validationErrors["email"] = "email is required"
	} else if !EmailRegex.MatchString(c.Email) {
		validationErrors["email"] = "invalid email format"
	}

	if strings.TrimSpace(c.Password) == "" {
		validationErrors["password"] = "password is required"
	} else if len(c.Password) < MinPasswordLength {
		validationErrors["password"] = "password must be at least 8 characters"
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
