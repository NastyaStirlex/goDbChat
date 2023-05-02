package db

import (
	"awesomeProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbase *gorm.DB

const dsn = "host=localhost user=user password=12345 dbname=mems_db sslmode=disable"

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(models.Mem{}, models.User{}, models.Result{})
	if err != nil {
		return nil
	}

	return db
}

func GetDb() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep *= 2
			fmt.Printf("Database is unavailable. Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
