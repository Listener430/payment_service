package models

import "gorm.io/gorm"

type UserFeature struct {
	gorm.Model
	UserID  string `gorm:"uniqueIndex"`
	Feature bool
}
