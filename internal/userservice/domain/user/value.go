package user

import (
	"regexp"
)

type Email string

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func NewEmail(address string) (Email, error) {
	if !emailRegex.MatchString(address) {
		return "", ErrInvalidEmail
	}
	return Email(address), nil
}

func (e Email) String() string {
	return string(e)
}
