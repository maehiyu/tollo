package user

import "time"

type UserType int

const (
	GeneralUser UserType = iota + 1
	ProfessionalUser
)

type Profile interface {
	IsProfile()
}

type ProfessionalProfile struct {
	ProBadgeURL string
	Biography string
}

func (p *ProfessionalProfile) IsProfile() {}

type GeneralProfile struct {
	Points uint32
	Introduction string
}

func (g *GeneralProfile) IsProfile() {}


type User struct {
	ID		string
	Name		string
	Email		Email
	Profile Profile
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
