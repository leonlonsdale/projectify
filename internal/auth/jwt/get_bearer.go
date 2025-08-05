package jwt

import (
	"net/http"
	"strings"

	"github.com/leonlonsdale/projectify/internal/errs"
)

func (j *JWT) GetBearerToken(headers http.Header) (string, error) {

	authorisation := headers.Get("Authorization")
	if authorisation == "" {
		return "", errs.ErrorNoBearerToken
	}

	token := strings.Fields(authorisation)

	if len(token) < 2 || token[0] != "Bearer" {
		return "", errs.ErrorNoBearerToken
	}

	return token[1], nil
}
