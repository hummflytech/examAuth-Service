package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordResetModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Email     string             `bson:"email"`
	OTP       string             `bson:"otp"`
	ExpiresAt time.Time          `bson:"expires_at"`
	Used      bool               `bson:"used"`
	CreatedAt time.Time          `bson:"created_at"`
}

type AdminPasswordResetModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Email     string             `bson:"email"`
	OTP       string             `bson:"otp"`
	ExpiresAt time.Time          `bson:"expires_at"`
	Used      bool               `bson:"used"`
	CreatedAt time.Time          `bson:"created_at"`
}
