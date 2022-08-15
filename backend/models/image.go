package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Path      string `gorm:"unique;not null;index" json:"path"`
	WebsiteID uint   `json:"website_id"`
	Website   Website
}
