package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}
