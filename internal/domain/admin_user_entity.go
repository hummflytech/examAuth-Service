package domain

import (
	"regexp"
	"time"
)

type AdminUser struct {
	id        string
	username  string
	email     string
	password  string
	phone     string
	createdAT time.Time
	isActive  *bool
}

func NewAdminUser(email string, password string, isActive *bool, username string, phone string) (*AdminUser, error) {
	if len(email) != 0 {
		regex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

		if !regex.MatchString(email) {
			return nil, ErrInvalidEmail
		}
	}

	if len(username) == 0 {
		return nil, ErrInvalidUsername
	}

	phonregex := regexp.MustCompile(`^(?:\+?251|0)?9\d{8}$`)

	if !phonregex.MatchString(phone) {
		return nil, ErrInvalidPhone
	}

	if len(password) < 4 {
		return nil, ErrInvalidPassword
	}

	nows := time.Now()
	return &AdminUser{
		email:     email,
		password:  password,
		isActive:  isActive,
		phone:     phone,
		username:  username,
		createdAT: nows,
	}, nil
}

func (a AdminUser) Id() string {
	return a.id
}

func (a AdminUser) Username() string {
	return a.username
}

func (a AdminUser) Email() string {
	return a.email
}

func (a AdminUser) Password() string {
	return a.password
}

func (a AdminUser) Phone() string {
	return a.phone
}

func (a AdminUser) CreatedAt() time.Time {
	return a.createdAT
}

func (a AdminUser) IsActive() bool {
	if a.isActive == nil {
		return false
	}
	return *a.isActive
}

func (a *AdminUser) Set_Id(id string) {
	a.id = id
}
