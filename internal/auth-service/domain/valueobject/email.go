package valueobject

import (
	"errors"
	"strings"
)

type Email string

func NewEmail(v string) (Email, error) {
	if len(v) < 5 || !strings.Contains(v, "@") {
		return "", errors.New("invalid email")
	}
	return Email(v), nil
}
