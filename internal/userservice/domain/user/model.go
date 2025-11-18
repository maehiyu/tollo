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
	Biography   string
}

func (p *ProfessionalProfile) IsProfile() {}

type GeneralProfile struct {
	Points       uint32
	Introduction string
}

func (g *GeneralProfile) IsProfile() {}

type User struct {
	ID        string
	Name      string
	Email     Email
	Profile   Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id string, email Email, name string, profile Profile) (*User, error) {
	if id == "" {
		return nil, ErrEmptyID
	}
	if name == "" {
		return nil, ErrEmptyName
	}
	if profile == nil {
		return nil, ErrNilProfile
	}

	now := time.Now()
	return &User{
		ID:        id,
		Email:     email,
		Name:      name,
		Profile:   profile,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
