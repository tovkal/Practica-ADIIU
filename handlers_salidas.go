package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAllSalidas(w http.ResponseWriter, r *http.Request) {
	salidas := []Salidas{}
	api.DB.Find(&salidas)
	WriteJson(w, &salidas)
}

func getSalidas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromDate := vars["fromDate"]
	toDate := vars["toDate"]
	salidas := []Salidas{}
	if api.DB.Where("(fechahora BETWEEN ? AND ?)", fromDate+" 00:00:00", toDate+" 23:59:59").Find(&salidas).Error != nil {
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &salidas)
}

func postSalida(w http.ResponseWriter, r *http.Request) {
	salida := Salidas{}
	if err := DecodeJson(r, &salida); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	salida.Fechahora = time.Now()

	if err := api.DB.Save(&salida).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &salida)
}
