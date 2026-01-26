package mapper

import (
	"github.com/Dawit0/examAuth/internal/delivery/dto"
	"github.com/Dawit0/examAuth/internal/domain"
)

func MapDomaintoResponse(domain domain.User) (*dto.UserResponse, error) {
	return &dto.UserResponse{
		ID:        domain.ID(),
		Username:  domain.Username(),
		Phone:     domain.Phone(),
		Email:     domain.Email(),
		Password:  domain.Password(),
		Badge:     domain.Badge(),
		IsActive:  domain.IsActive(),
		Score:     domain.Score(),
		CreatedAt: domain.CreatedAt(),
	}, nil
}

func MApAdminDomainToResponse(adminuser *domain.AdminUser) (*dto.AdminUserResponse, error)  {
	return &dto.AdminUserResponse{
		ID:        adminuser.Id(),
		Username:  adminuser.Username(),
		Phone:     adminuser.Phone(),
		Email:     adminuser.Email(),
		Password:  adminuser.Password(),
		IsActive:  adminuser.IsActive(),
		CreatedAt: adminuser.CreatedAt(),
	}, nil
}
