package models

import "gorm.io/gorm"

type HashtagProject struct {
	gorm.Model
	HashtagId uint `json:"hashtagId" gorm:"uniqueIndex:idx_hashtagid_projectid"`
	ProjectId uint `json:"projectId" gorm:"uniqueIndex:idx_hashtagid_projectid"`
}
