package adminuserrepo

import (
	"context"
	"errors"
	"time"

	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/mapper"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminUserRepo struct {
	DB *mongo.Database
}

func NewAdminUserRepo(db *mongo.Database) *AdminUserRepo {
	return &AdminUserRepo{DB: db}
}

func (au *AdminUserRepo) CreateAdmins(adminuser *domain.AdminUser) (*domain.AdminUser, error) {
	adminModel, err := mapper.MapAdminDomainToModel(*adminuser)
	if err != nil {
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	res, err := au.DB.Collection("admin_users").InsertOne(ctx, adminModel)
	if err != nil {
		return nil, err
	}

	adminModel.ID = res.InsertedID.(primitive.ObjectID)

	val, err := mapper.MapAdminModelToDomain(*adminModel)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (au *AdminUserRepo) FindByEmail(email string) (*domain.AdminUser, error) {
	var adminModel model.AdminUserModel

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	err := au.DB.Collection("admin_users").FindOne(ctx, bson.M{"email": email}).Decode(&adminModel)
	if err != nil {
		return nil, err
	}

	val, err := mapper.MapAdminModelToDomain(adminModel)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (au *AdminUserRepo) FindById(id string) (*domain.AdminUser, error) {
	var adminModel model.AdminUserModel

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = au.DB.Collection("admin_users").FindOne(ctx, bson.M{"_id": objID}).Decode(&adminModel)
	if err != nil {
		return nil, err
	}

	val, err := mapper.MapAdminModelToDomain(adminModel)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (au *AdminUserRepo) UpdateAdmins(id string, adminuser *domain.AdminUser) (*domain.AdminUser, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	// 1. Check if the email is already taken by ANOTHER user
	// Filter: { email: "new@email.com", _id: { $ne: current_user_id } }
	duplicateFilter := bson.M{
		"email": adminuser.Email(),
		"_id":   bson.M{"$ne": objID},
	}

	count, err := au.DB.Collection("admin_users").CountDocuments(ctx, duplicateFilter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email is already in use by another account")
	}

	adminModel, err := mapper.MapAdminDomainToModel(*adminuser)
	if err != nil {
		return nil, err
	}
	adminModel.ID = objID

	update := bson.M{
		"$set": adminModel,
	}

	_, err = au.DB.Collection("admin_users").UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	val, err := mapper.MapAdminModelToDomain(*adminModel)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (au *AdminUserRepo) DeleteAdmins(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	_, err = au.DB.Collection("admin_users").DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (au *AdminUserRepo) AllAdmins() ([]domain.AdminUser, error) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	cursor, err := au.DB.Collection("admin_users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var models []model.AdminUserModel
	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	adminUsers := make([]domain.AdminUser, 0, len(models))
	for _, item := range models {
		val, err := mapper.MapAdminModelToDomain(item)
		if err != nil {
			continue
		}
		adminUsers = append(adminUsers, *val)
	}

	return adminUsers, nil
}
