package models

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	Domain string `gorm:"unique;not null;index" json:"domain"`
	Title  string `gorm:"not null" json:"title"`
}
