package models

import "gorm.io/gorm"

type UserProject struct {
	gorm.Model
	UserId    uint `json:"userId" gorm:"uniqueIndex:idx_userid_projectid"`
	ProjectId uint `json:"projectId" gorm:"uniqueIndex:idx_userid_projectid"`
}
