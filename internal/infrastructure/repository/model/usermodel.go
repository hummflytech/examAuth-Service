package model

import "time"

type UserModel struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Phone     string
	Email     *string
	Password  string
	CreatedAt time.Time
	IsActive  *bool
	Badge     *string
	Score     *float64
}
