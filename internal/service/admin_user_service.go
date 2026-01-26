package service

import (
	"errors"

	"github.com/Dawit0/examAuth/internal/domain"
	adminrepo "github.com/Dawit0/examAuth/internal/infrastructure/repository/adminUser_repo"
	"github.com/Dawit0/examAuth/internal/infrastructure/security"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AdminUserService struct {
	repo *adminrepo.AdminUserRepo
}

func NewAdminUserService(repo *adminrepo.AdminUserRepo) *AdminUserService {
	return &AdminUserService{repo: repo}
}

func (as *AdminUserService) CreateAdmins(adminuser *domain.AdminUser) (*domain.AdminUser, error) {
	if adminuser == nil {
		return nil, errors.New("adminuser is nil")
	}
	val, _ := as.repo.FindByEmail(adminuser.Email())
	if val != nil {
		return nil, errors.New("admin already exist at this email")
	}
	return as.repo.CreateAdmins(adminuser)
}

func (as *AdminUserService) AdminLogin(email string, password string) (*domain.AdminUser, error) {
	val, err := as.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if val == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(val.Password()), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return val, nil
}

func (as *AdminUserService) FindById(id string) (*domain.AdminUser, error) {
	return as.repo.FindById(id)
}

func (as *AdminUserService) UpdateAdmins(id string, adminuser *domain.AdminUser) (*domain.AdminUser, error) {
	if adminuser == nil {
		return nil, errors.New("adminuser is nil")
	}
	val, _ := as.repo.FindById(id)
	if val == nil {
		return nil, errors.New("admin not found")
	}
	return as.repo.UpdateAdmins(id, adminuser)
}

func (as *AdminUserService) DeleteAdmins(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	val, _ := as.repo.FindById(id)
	if val == nil {
		return errors.New("admin not found")
	}
	return as.repo.DeleteAdmins(id)
}

func (as *AdminUserService) AllAdmins() ([]domain.AdminUser, error) {
	return as.repo.AllAdmins()
}

func (as *AdminUserService) FindByEmail(email string) (*domain.AdminUser, error) {
	return as.repo.FindByEmail(email)
}

func (as *AdminUserService) ValidateToken(token string) (bool, string, error) {
	val, err := security.VerifyToken(token)

	if err != nil || !val.Valid {
		return false, "", err
	}

	claims := val.Claims.(jwt.MapClaims)

	user_id := claims["user_id"].(string)

	return true, user_id, nil
}
