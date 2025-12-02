package service

import (
	"errors"

	"github.com/Dawit0/examAuth/internal/domain"
	repo "github.com/Dawit0/examAuth/internal/infrastructure/repository/userRepo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService(userrp *repo.UserRepo) *UserService {
	return &UserService{UserRepo: userrp}
}

func (uc *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, nil
	}
	return uc.UserRepo.CreateUser(user)
}

func (uc *UserService) UserLogin(phone string, password string) (*domain.User, error) {
	user, err := uc.UserRepo.FindByPhone(phone)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return user, nil

}
