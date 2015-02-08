package main

import (
	"encoding/gob"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

type Api struct {
	DB    gorm.DB
	store *sessions.CookieStore
}

// Init the database
func (api *Api) initDB() {
	var err error
	api.DB, err = gorm.Open("mysql", "wsfarmacia:wsfarmacia@/wsfarmacia?charset=utf8")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	api.DB.LogMode(true)
}

func (api *Api) initSessionStore() {
	api.store = sessions.NewCookieStore([]byte("something-very-secret"))

	gob.Register(&Farmacias{})
}
