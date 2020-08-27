package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"os"
)

var Gdb *gorm.DB

func InitSqlite(dbPath string) {

	var err error
	fi, err := os.Stat(dbPath)
	if err != nil {
		panic(err)
	}

	log.Info("open db file", fi.Name())

	Gdb, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	err = Gdb.DB().Ping()
	if err != nil {
		log.Error("ping db failed")
		return
	}
	log.Info("init engine succeed")
	return
}
