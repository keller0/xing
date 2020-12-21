package storage

import (
	"time"
)

func (Notes) TableName() string { return "notes" }

type Notes struct {
	Id         string    `json:"id" gorm:"column:id"`
	Uid        string    `json:"uid" gorm:"column:uid"`
	Content    string    `json:"content" gorm:"column:content"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time; type:datetime"`
}

func (user *User) AddNotes(notes Notes) error {

	notes.Uid = user.Id
	return Gdb.Create(&notes).Error

}
