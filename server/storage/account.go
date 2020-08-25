package storage

import "golang.org/x/crypto/bcrypt"

func (User) TableName() string { return "user" }

type User struct {
	Id   string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Pass string `json:"pass" gorm:"column:pass"`
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
