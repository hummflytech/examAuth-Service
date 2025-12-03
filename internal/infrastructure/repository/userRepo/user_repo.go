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

func (ur *UserRepo) FindByID(id uint) (*domain.User, error) {
	var models model.UserModel
	err := ur.DB.Model(&model.UserModel{}).Where("id = ?", id).First(&models).Error
	if err != nil {
		return nil, err
	}

	domain, err := mapper.MapModelToDomain(models)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (ur *UserRepo) AllUsers() ([]domain.User, error) {
	var models []model.UserModel

	err := ur.DB.Model(&model.UserModel{}).Find(&models).Error
	if err != nil {
		return nil, err
	}

	domain := make([]domain.User, 0, len(models))

	for _, item := range models {
		val, err := mapper.MapModelToDomain(item)

		if err != nil {
			continue
		}

		domain = append(domain, *val)
	}

	return domain, nil
}

func (ur *UserRepo) DeleteUser(id uint) error {
	err := ur.DB.Model(&model.UserModel{}).Delete(&model.UserModel{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) UpdateUser(id uint, user *domain.User) (*domain.User, error) {
	models, err := mapper.MapDomainToModel(*user)
	if err != nil {
		return nil, err
	}

	errs := ur.DB.Model(&model.UserModel{}).Where("id = ?", id).Updates(&models).Error
	if errs != nil {
		return nil, errs
	}

	domain, er := mapper.MapModelToDomain(*models)
	if er != nil {
		return nil, er
	}

	return domain, nil
}
