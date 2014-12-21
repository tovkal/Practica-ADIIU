package main

import (
	"net/http"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *Api) GetAllEntradas(w rest.ResponseWriter, r *rest.Request) {
	entradas := []Entradas{}
	api.DB.Find(&entradas)
	w.WriteJson(&entradas)
}

func (api *Api) GetEntrada(w rest.ResponseWriter, r *rest.Request) {
	fromDate := r.PathParam("fromDate")
	toDate := r.PathParam("toDate")
	entrada := Entradas{}
	if api.DB.Where("fechahora >= ? and fechahora <=", fromDate, toDate).First(&entrada).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&entrada)
}

func (api *Api) PostEntrada(w rest.ResponseWriter, r *rest.Request) {
	entrada := Entradas{}
	if err := r.DecodeJsonPayload(&entrada); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&entrada).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&entrada)
}
