package storage

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (User) TableName() string { return "user" }

type User struct {
	Id         string    `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	Pass       string    `json:"pass" gorm:"column:pass"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (user *User) CheckUserAuth() bool {
	pass := user.Pass
	Gdb.Where("name = ?", user.Name).First(&user)
	if len(user.Id) == 0 {
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
	user.Id = u.String()
	user.CreateTime = time.Now()

	return Gdb.Create(*user).Error

}
