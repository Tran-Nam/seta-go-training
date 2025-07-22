package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	TeamId   string `gorm:"not null;uniqueIndex"`
	TeamName string `gorm:"not null"`
	Mangers  []User `gorm:"many2many:team_managers;"`
	Members  []User `gorm:"many2many:team_members;"`
}
