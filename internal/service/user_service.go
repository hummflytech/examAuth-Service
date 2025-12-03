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
	val, _ := uc.UserRepo.FindByPhone(user.Phone())
	// if err != nil {
	// 	return nil, err
	// }
	if val != nil {
		return nil, errors.New("user already exist at this phone number")
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

func (uc *UserService) FindByID(id uint) (*domain.User, error) {
	return uc.UserRepo.FindByID(id)
}

func (uc *UserService) AllUsers() ([]domain.User, error) {
	return uc.UserRepo.AllUsers()
}

func (uc *UserService) DeleteUser(id uint) error {
	val, err := uc.UserRepo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}
	if val == nil {
		return errors.New("user not found")
	}
	return uc.UserRepo.DeleteUser(id)
}

func (uc *UserService) UpdateUser(id uint, user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, nil
	}
	val, _ := uc.UserRepo.FindByID(id)
	if val == nil {
		return nil, errors.New("user not found")
	}
	return uc.UserRepo.UpdateUser(id, user)
}
