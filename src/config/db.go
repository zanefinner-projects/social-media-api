package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //We need this for the dialect
)

//Migrate sets up the database
func Migrate(uname, pword, dbname string) {
	db, err := gorm.Open("mysql", uname+":"+pword+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})
	if err != nil {
		log.Panic(err)
	}
	db.Close()
}

//GetDBCreds gives creds to access the db
func GetDBCreds() (string, string, string) {
	return "zane", "5245", "devapi"
}
