package user

import "time"

type UserType int

const (
	GeneralUser UserType = iota + 1
	ProfessionalUser
)

type User struct {
	ID		string
	Name		string
	Email		Email
	Description string
	Type		UserType
	CreatedAt	time.Time
	UpdatedAt	time.Time

	ProBadgeURL string
	Specialty   string

	Points      uint32
}