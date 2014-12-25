package ws

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Api struct {
	DB gorm.DB
}

// Init the database
func (api *Api) InitDB() {
	var err error
	api.DB, err = gorm.Open("mysql", "mascaro:mascaro@/wsfarmacia?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	api.DB.LogMode(true)
}
