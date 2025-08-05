package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CustomerRecord struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type CustomerRegistration struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerUpdate struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CustomerSafe struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (c *CustomerRegistration) Validate() map[string]string {
	validationErrors := make(map[string]string)

	if strings.TrimSpace(c.Name) == "" {
		validationErrors["name"] = "name is required"
	}

	if strings.TrimSpace(c.Email) == "" {
		validationErrors["email"] = "email is required"
	} else if !emailRegex.MatchString(c.Email) {
		validationErrors["email"] = "invalid email format"
	}

	if strings.TrimSpace(c.Password) == "" {
		validationErrors["password"] = "password is required"
	} else if len(c.Password) < 8 {
		validationErrors["password"] = "password must be at least 8 characters"
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
