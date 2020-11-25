package storage

import "time"

type Notes struct {
	Id         string    `json:"id" gorm:"column:id"`
	Uid        string    `json:"uid" gorm:"column:uid"`
	Content    string    `json:"content" gorm:"column:content"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time; type:datetime"`
}
