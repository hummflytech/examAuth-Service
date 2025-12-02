package dto

import "time"

type UserCreate struct {
	Username string   `json:"username" binding:"required"`
	Phone    string   `json:"phone" binding:"required"`
	Email    *string  `json:"email"`
	Password string   `json:"password" binding:"required"`
	Badge    *string  `json:"badge"`
	IsActive *bool    `json:"is_active"`
	Score    *float64 `json:"score"`
}

type UserResponse struct {
	ID        uint
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
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}
