package main

import (
	"net/http"
	"time"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *api) getAllSalidas(w rest.ResponseWriter, r *rest.Request) {
	salidas := []Salidas{}
	api.DB.Find(&salidas)
	w.WriteJson(&salidas)
}

func (api *api) getSalida(w rest.ResponseWriter, r *rest.Request) {
	fromDate := r.PathParam("fromDate")
	toDate := r.PathParam("toDate")
	salidas := []Salidas{}
	if api.DB.Where("(fechahora BETWEEN ? AND ?)", fromDate+" 00:00:00", toDate+" 23:59:59").Find(&salidas).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&salidas)
}

func (api *api) postSalida(w rest.ResponseWriter, r *rest.Request) {
	salida := Salidas{}
	if err := r.DecodeJsonPayload(&salida); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	salida.Fechahora = time.Now()

	if err := api.DB.Save(&salida).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&salida)
}
