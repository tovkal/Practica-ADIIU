package main

import (
	"net/http"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *Api) GetAllSalidas(w rest.ResponseWriter, r *rest.Request) {
	salidas := []Salidas{}
	api.DB.Find(&salidas)
	w.WriteJson(&salidas)
}

func (api *Api) GetSalida(w rest.ResponseWriter, r *rest.Request) {
	fromDate := r.PathParam("fromDate")
	toDate := r.PathParam("toDate")
	salida := Salidas{}
	if api.DB.Where("fechahora >= ? AND fechahora <= ?", fromDate, toDate).First(&salida).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&salida)
}

func (api *Api) PostSalida(w rest.ResponseWriter, r *rest.Request) {
	salida := Salidas{}
	if err := r.DecodeJsonPayload(&salida); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&salida).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&salida)
}
