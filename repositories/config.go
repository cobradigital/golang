package repositories

import (
	"os"

	"../configs"
	"../loggers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var log = loggers.Get()

var DBConnect *gorm.DB

func Config() {
	// Init error codes
	DBConn()
	Ping()
}

func DBConn() {
	driver := configs.MustGetString("database.master.driver")
	host := configs.MustGetString("database.master.host")
	user := configs.MustGetString("database.master.username")
	pswd := configs.MustGetString("database.master.password")
	dbnm := configs.MustGetString("database.master.database")

	db, _ := gorm.Open(driver, user+":"+pswd+"@"+host+"/"+dbnm+"?charset=utf8&parseTime=True&loc=Local")

	db.SingularTable(true)

	db.LogMode(true)

	DBConnect = db
}

func Ping() {
	ping := DBConnect.DB().Ping()
	if ping != nil {
		log.Info("Failed Connecting Database.")
		os.Exit(1)
	} else {
		log.Info("Success Connecting Database.")
	}
}
