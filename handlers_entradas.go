package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAllEntradas(w http.ResponseWriter, r *http.Request) {
	entradas := []Entradas{}
	api.DB.Find(&entradas)
	WriteJson(w, &entradas)
}

func getEntradas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromDate := vars["fromDate"]
	toDate := vars["toDate"]
	entradas := []Entradas{}
	if api.DB.Where("(fechahora BETWEEN ? AND ?)", fromDate+" 00:00:00", toDate+" 23:59:59").Find(&entradas).Error != nil {
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &entradas)
}

func postEntrada(w http.ResponseWriter, r *http.Request) {
	entrada := Entradas{}
	if err := DecodeJson(r, &entrada); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entrada.Fechahora = time.Now()

	if err := api.DB.Save(&entrada).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &entrada)
}
