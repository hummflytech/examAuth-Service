package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminUserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Phone     string             `bson:"phone"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	IsActive  *bool              `bson:"is_active"`
}
