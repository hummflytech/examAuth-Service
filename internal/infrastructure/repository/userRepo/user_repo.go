package userRepo

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

type UserRepo struct {
	DB *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{DB: db}
}

func (ur *UserRepo) CreateUser(user *domain.User) (*domain.User, error) {
	models, errs := mapper.MapDomainToModel(*user)
	if errs != nil {
		return nil, errs
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ur.DB.Collection("users").InsertOne(ctx, models)
	if err != nil {
		return nil, err
	}

	models.ID = res.InsertedID.(primitive.ObjectID)

	val, err := mapper.MapModelToDomain(*models)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (ur *UserRepo) FindByEmail(email string) (*domain.User, error) {
	var models model.UserModel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ur.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&models)
	if err != nil {
		return nil, err
	}

	domain, err := mapper.MapModelToDomain(models)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (ur *UserRepo) FindByID(id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var models model.UserModel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ur.DB.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&models)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := ur.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var models []model.UserModel
	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	domainUsers := make([]domain.User, 0, len(models))
	for _, item := range models {
		val, err := mapper.MapModelToDomain(item)
		if err != nil {
			continue
		}
		domainUsers = append(domainUsers, *val)
	}

	return domainUsers, nil
}

func (ur *UserRepo) DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = ur.DB.Collection("users").DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (ur *UserRepo) UpdateUser(id string, user *domain.User) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Check if the email is already taken by ANOTHER user
	// Filter: { email: "new@email.com", _id: { $ne: current_user_id } }
	duplicateFilter := bson.M{
		"email": user.Email(),
		"_id":   bson.M{"$ne": objID},
	}

	count, err := ur.DB.Collection("users").CountDocuments(ctx, duplicateFilter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email is already in use by another account")
	}

	// 2. Map and Prepare Update
	models, err := mapper.MapDomainToModel(*user)
	if err != nil {
		return nil, err
	}
	models.ID = objID

	update := bson.M{
		"$set": models,
	}

	// 3. Execute Update
	_, err = ur.DB.Collection("users").UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	domainRes, er := mapper.MapModelToDomain(*models)
	if er != nil {
		return nil, er
	}

	return domainRes, nil
}
