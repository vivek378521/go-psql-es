package models

import "gorm.io/gorm"

type Hashtag struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}
