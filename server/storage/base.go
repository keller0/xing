package storage

import (
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var JWTSigningKey = []byte("tom go")

type JwtClaims struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	jwt.StandardClaims
}

const CtxUserKey = "userInfoKeyInCtx"

type CtxUserInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Gdb *gorm.DB

func InitDB(dsn string) {

	var err error

	//dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	Gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Info("init engine succeed")
	return
}
