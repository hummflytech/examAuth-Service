package service

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	adminuserrepo "github.com/Dawit0/examAuth/internal/infrastructure/repository/adminUser_repo"
)

type AdminResetService struct {
	repo   *adminuserrepo.AdminUserResetRepo
	mailer SendMailer
	otpTTl time.Duration
}

func NewAdminResetService(repo *adminuserrepo.AdminUserResetRepo, mailer SendMailer) *AdminResetService {
	return &AdminResetService{repo: repo, mailer: mailer, otpTTl: 15 * time.Minute}
}

func (ar *AdminResetService) RequestResetPasswordEmail(email string) error {
	user, err := ar.repo.GetByEmail(email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("user not found")
		}
		return err
	}

	otp, err := generateNumericOTP(6)
	if err != nil {
		return err
	}
	expiredAt := time.Now().Add(ar.otpTTl)
	err = ar.repo.SavePasswordReset(email, user.ID.Hex(), otp, expiredAt)
	if err != nil {
		return err
	}
	return ar.mailer.SendMail(email, otp)

}

func (ar *AdminResetService) ResetPassword(email string, otp string, newPassword string) error {
	token, err := ar.repo.FindValidResetByEmailAndOTP(email, otp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("invalid reset token")
		}
		return err
	}
	if token == nil || token.Used {
		return errors.New("invalid reset token")
	}
	err = ar.repo.MarkPasswordResetUsed(token.ID.Hex())
	if err != nil {
		return err
	}
	user, err := ar.repo.GetByEmail(token.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = ar.repo.UpdateAdminPassword(user.ID.Hex(), string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}
