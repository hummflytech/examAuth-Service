package dto

import (
	"time"
)

type UserCreate struct {
	Username string   `json:"username" binding:"required"`
	Phone    string   `json:"phone" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Badge    *string  `json:"badge"`
	IsActive *bool    `json:"is_active"`
	Score    *float64 `json:"score"`
}

type AdminUserCreate struct {
	Username string   `json:"username" binding:"required"`
	Phone    string   `json:"phone" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
}

type AdminUserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminUserResponse struct {
	ID        string
	Username  string
	Phone     string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
}

type UserResponse struct {
	ID        string
	Username  string
	Phone     string
	Email     string
	Password  string
	Badge     string
	IsActive  bool
	Score     float64
	CreatedAt time.Time
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestResetDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordDTO struct {
	Email       string `json:"email" binding:"required,email"`
	OTP         string `json:"otp" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

