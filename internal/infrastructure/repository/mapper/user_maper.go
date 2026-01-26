package mapper

import (
	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func MapDomainToModel(user domain.User) (*model.UserModel, error) {
	email := user.Email()
	badge := user.Badge()
	isActive := user.IsActive()
	score := user.Score()
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password()), bcrypt.DefaultCost)

	var id primitive.ObjectID
	if user.ID() != "" {
		var err error
		id, err = primitive.ObjectIDFromHex(user.ID())
		if err != nil {
			// If it's not a valid hex, we might want to handle it,
			// but for new users it will be empty and omitempty will handle it.
		}
	}

	return &model.UserModel{
		ID:        id,
		Username:  user.Username(),
		Phone:     user.Phone(),
		Email:     email,
		Password:  string(pass),
		CreatedAt: user.CreatedAt(),
		IsActive:  &isActive,
		Badge:     &badge,
		Score:     &score,
	}, nil
}

func MapModelToDomain(model model.UserModel) (*domain.User, error) {
	domain_val, err := domain.NewUser(model.Email, model.Password, model.Badge, model.Username, model.Phone, model.IsActive, model.Score)
	if err != nil {
		return nil, err
	}

	domain_val.Id_Set(model.ID.Hex())
	return domain_val, nil
}

func MapAdminDomainToModel(admin_user domain.AdminUser) (*model.AdminUserModel, error) {
	email := admin_user.Email()
	username := admin_user.Username()
	phone := admin_user.Phone()
	isActive := admin_user.IsActive()
	pass, _ := bcrypt.GenerateFromPassword([]byte(admin_user.Password()), bcrypt.DefaultCost)

	var id primitive.ObjectID
	if admin_user.Id() != "" {
		var err error
		id, err = primitive.ObjectIDFromHex(admin_user.Id())
		if err != nil {
			// If it's not a valid hex, we might want to handle it,
			// but for new users it will be empty and omitempty will handle it.
		}
	}

	return &model.AdminUserModel{
		ID:        id,
		Username:  username,
		Phone:     phone,
		Email:     email,
		Password:  string(pass),
		CreatedAt: admin_user.CreatedAt(),
		IsActive:  &isActive,
	}, nil
}

func MapAdminModelToDomain(model model.AdminUserModel) (*domain.AdminUser, error) {
	admin_user, err := domain.NewAdminUser(model.Email, model.Password, model.IsActive, model.Username, model.Phone)
	if err != nil {
		return nil, err
	}

	admin_user.Set_Id(model.ID.Hex())
	return admin_user, nil
}
