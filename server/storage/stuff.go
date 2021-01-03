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

func AddNotes(notes Notes) error {

	notes.CreateTime = time.Now()
	return Gdb.Create(&notes).Error

}

func GetNotesByUid(uid string) ([]Notes, error) {
	var notes []Notes
	err := Gdb.Where("uid = ?", uid).Find(&notes).Error
	return notes, err
}

func GetNotesById(id, uid string) (Notes, error) {
	var notes Notes
	err := Gdb.Where("id = ? AND uid = ?", id, uid).Find(&notes).Error
	return notes, err
}
