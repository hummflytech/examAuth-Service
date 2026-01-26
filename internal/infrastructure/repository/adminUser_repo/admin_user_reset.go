package adminuserrepo

import (
	"context"
	"time"

	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminUserResetRepo struct {
	DB *mongo.Database
}

func NewAdminUserResetRepo(db *mongo.Database) *AdminUserResetRepo {
	return &AdminUserResetRepo{DB: db}
}

func (aur *AdminUserResetRepo) GetByEmail(email string) (*model.AdminUserModel, error) {
	var user model.AdminUserModel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := aur.DB.Collection("admin_users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (aur *AdminUserResetRepo) SavePasswordReset(email string, userID string, otp string, expiredAt time.Time) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	reset := model.AdminPasswordResetModel{
		UserID:    objID,
		Email:     email,
		OTP:       otp,
		ExpiresAt: expiredAt,
		Used:      false,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = aur.DB.Collection("admin_password_resets").InsertOne(ctx, reset)
	return err
}

func (aur *AdminUserResetRepo) FindValidResetByEmailAndOTP(email, otp string) (*model.AdminPasswordResetModel, error) {
	var reset model.AdminPasswordResetModel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := aur.DB.Collection("admin_password_resets").FindOne(ctx, bson.M{
		"email":      email,
		"otp":        otp,
		"used":       false,
		"expires_at": bson.M{"$gt": time.Now()},
	}).Decode(&reset)

	if err != nil {
		return nil, err
	}
	return &reset, nil
}

func (aur *AdminUserResetRepo) MarkPasswordResetUsed(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = aur.DB.Collection("admin_password_resets").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"used": true}})
	return err
}

func (aur *AdminUserResetRepo) UpdateAdminPassword(userID string, hashedPassword string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = aur.DB.Collection("admin_users").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"password": hashedPassword}})
	return err
}
