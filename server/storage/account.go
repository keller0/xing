package storage

import (
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (User) TableName() string { return "users" }

type User struct {
	Id         int       `json:"id" gorm:"autoIncrement;column:id"`
	UId        string    `json:"uid" gorm:"primaryKey;column:uid"`
	Name       string    `json:"name" gorm:"column:name"`
	Pass       string    `json:"pass" gorm:"column:pass"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime"`
}

func (user *User) CheckUserAuth() bool {
	pass := user.Pass
	Gdb.Where("name = ?", user.Name).First(&user)
	if len(user.UId) == 0 {
		return false
	}
	return CheckPasswordHash(pass, user.Pass)

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) Add() error {

	if len(user.Name) < 3 || len(user.Pass) < 5 {
		return errors.New("username or password is too short")
	}

	hash, err := HashPassword(user.Pass)
	if err != nil {
		return err
	}
	user.Pass = hash
	u, _ := uuid.NewRandom()
	user.UId = u.String()
	user.CreateTime = time.Now()

	res := Gdb.Create(*user)
	log.Debug(res)
	return res.Error

}
