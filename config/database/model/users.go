package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID             string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Role           string `gorm:"column:role;default:'User'"`
	Username       string `gorm:"column:username;unique;not null"`
	Email          string `gorm:"column:email;unique;not null"`
	HashedPassword string `gorm:"column:hashed_password;type:text"`
	UpdatedAt      time.Time
	CreatedAt      time.Time
	UserDetail     UserDetail `gorm:"foreignKey:UserID"` // Define the relationship
}

type UserDetail struct {
	gorm.Model
	ID             string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         string `gorm:"column:user_id;type:uuid"` // Foreign key to Users
	FirstName      string `gorm:"column:first_name"`
	MiddleName     string `gorm:"column:middle_name"`
	LastName       string `gorm:"column:last_name"`
	DateOfBirth    time.Time
	Address        string
	PhoneNumber    string `gorm:"column:phone_number;unique"`
	ProfilePicture string `gorm:"column:profile_picture;type:text"`
	UpdatedAt      time.Time
	CreatedAt      time.Time
}
