package domain

import "time"

type ForgetPassword struct {
	id        uint
	userID    uint
	email     string
	otp       string
	expiredAt time.Time
	used      bool
}

func NewForgetPassword(userID uint, email string, otp string, expiredAt time.Time, used bool) (*ForgetPassword, error) {
	return &ForgetPassword{
		userID:    userID,
		email:     email,
		otp:       otp,
		expiredAt: expiredAt,
		used:      used,
	}, nil
}

func (f ForgetPassword) AdminId() uint {
	return f.id
}

func (f ForgetPassword) AdminUserId() uint {
	return f.userID
}

func (f ForgetPassword) AdminEmail() string {
	return f.email
}

func (f ForgetPassword) AdminOtp() string {
	return f.otp
}

func (f ForgetPassword) AdminExpiredAt() time.Time {
	return f.expiredAt
}

func (f ForgetPassword) AdminUsed() bool {
	return f.used
}

func (f *ForgetPassword) Set_AdminId(id uint) {
	f.id = id
}
