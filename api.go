package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type api struct {
	DB gorm.DB
}

// Init the database
func (api *api) initDB() {
	var err error
	api.DB, err = gorm.Open("mysql", "wsfarmacia:wsfarmacia@/wsfarmacia?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	api.DB.LogMode(true)
}
