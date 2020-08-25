package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var Gdb *gorm.DB

func InitSqlite(dbPath string) {
	if len(dbPath) == 0 {
		panic("no db path")
	}
	var err error
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
