package main

import (
	"flag"
	"fmt"
	"github.com/keller0/xing/server/config"
	"github.com/keller0/xing/server/storage"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func testHash(password string) {
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

var (
	dbPath   string
	userName string
	pass     string
)

func main() {
	var configFile = flag.String("conf", "./config.yaml", "config file path")
	flag.StringVar(&userName, "u", "", "new user name")
	flag.StringVar(&pass, "p", "", "new user pass")

	flag.Parse()

	var configs config.AllConf
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic("read config file failed")
	}
	if err := yaml.Unmarshal(content, &configs); err != nil {
		panic("unmarshal config content failed")
	}

	storage.InitDB(configs.App.MysqlDSN)

	u := storage.User{Name: userName, Pass: pass}

	err = u.Add()
	if err != nil {
		panic(err)
	}
	fmt.Println("add user ", userName, "done")

}
