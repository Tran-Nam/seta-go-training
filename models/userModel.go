package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	Member  UserRole = "MEMBER"
	Manager UserRole = "MANAGER"
)

type User struct {
	gorm.Model
	UserId   string   `gorm:"not null;uniqueIndex"`
	UserName string   `gorm:"not null"`
	Email    string   `gorm:"not null"`
	Password string   `gorm:"not null"`
	Role     UserRole `gorm:"type:varchar(20);not null;default:'MEMBER'"`
	Teams    []Team   `gorm:"many2many:team_members;"`
}
