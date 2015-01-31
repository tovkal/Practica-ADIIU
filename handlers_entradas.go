package main

import (
	"net/http"
	"time"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *api) getAllEntradas(w rest.ResponseWriter, r *rest.Request) {
	entradas := []Entradas{}
	api.DB.Find(&entradas)
	w.WriteJson(&entradas)
}

func (api *api) getEntrada(w rest.ResponseWriter, r *rest.Request) {
	fromDate := r.PathParam("fromDate")
	toDate := r.PathParam("toDate")
	entradas := []Entradas{}
	if api.DB.Where("(fechahora BETWEEN ? AND ?)", fromDate+" 00:00:00", toDate+" 23:59:59").Find(&entradas).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&entradas)
}

func (api *api) postEntrada(w rest.ResponseWriter, r *rest.Request) {
	entrada := Entradas{}
	if err := r.DecodeJsonPayload(&entrada); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entrada.Fechahora = time.Now()

	if err := api.DB.Save(&entrada).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&entrada)
}
