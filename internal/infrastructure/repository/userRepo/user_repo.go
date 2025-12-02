package repository

import (
	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/mapper"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	db.AutoMigrate(&model.UserModel{})
	return &UserRepo{DB: db}
}

func (ur *UserRepo) CreateUser(user *domain.User) (*domain.User, error) {
	models, errs := mapper.MapDomainToModel(*user)
	if errs != nil {
		return nil, errs
	}

	err := ur.DB.Model(&model.UserModel{}).Create(&models).Error
	if err != nil {
		return nil, err
	}

	val, err := mapper.MapModelToDomain(*models)
	if err != nil {
		return nil, err
	}

	return val, nil

}

func (ur *UserRepo) FindByPhone(phone string) (*domain.User, error) {
	var models model.UserModel
	err := ur.DB.Model(&model.UserModel{}).Where("phone=?", phone).First(&models).Error
	if err != nil {
		return nil, err
	}

	domain, err := mapper.MapModelToDomain(models)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
