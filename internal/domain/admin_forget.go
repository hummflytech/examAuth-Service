package domain


import "time"

type AdminForgetPassword struct {
	id        uint
	userID    uint
	email     string
	otp       string
	expiredAt time.Time
	used      bool
}

func NewAdminForgetPassword(userID uint, email string, otp string, expiredAt time.Time, used bool) (*AdminForgetPassword, error) {
	return &AdminForgetPassword{
		userID:    userID,
		email:     email,
		otp:       otp,
		expiredAt: expiredAt,
		used:      used,
	}, nil
}

func (f ForgetPassword) Id() uint {
	return f.id
}

func (f ForgetPassword) UserId() uint {
	return f.userID
}

func (f ForgetPassword) Email() string {
	return f.email
}

func (f ForgetPassword) Otp() string {
	return f.otp
}

func (f ForgetPassword) ExpiredAt() time.Time {
	return f.expiredAt
}

func (f ForgetPassword) Used() bool {
	return f.used
}

func (f *ForgetPassword) Set_Id(id uint) {
	f.id = id
}

