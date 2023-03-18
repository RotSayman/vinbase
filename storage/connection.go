package storage

import (
	"log"

	"github.com/RotSayman/vinbase/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	if dbase == nil {
		dbase, err := gorm.Open(sqlite.Open("base.db"), &gorm.Config{})
		if res := dbase.Exec("PRAGMA case_sensitive_like = OFF; PRAGMA encoding = 'UTF-8';", nil); res.Error != nil {
			log.Printf("%v \n", res.Error)
		}
		if err != nil {
			log.Fatal(err)
		}
		dbase.AutoMigrate(&models.Vin{})
		return dbase
	}
	return dbase
}
