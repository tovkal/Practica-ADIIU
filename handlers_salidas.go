package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAllSalidas(w http.ResponseWriter, r *http.Request) {
	salidas := []SalidasJoin{}
	api.DB.Table("salidas").Select("salidas.id, salidas.idmedicamento, medicamentos.nombre as nombremedicamento, salidas.fechahora, salidas.cantidad, salidas.idfarmacia, farmacias.nik as nombrefarmacia").Joins("INNER JOIN medicamentos ON salidas.idmedicamento = medicamentos.id INNER JOIN farmacias ON salidas.idfarmacia = farmacias.id").Scan(&salidas)
	WriteJson(w, &salidas)
}

func getSalidas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromDate := vars["fromDate"]
	toDate := vars["toDate"]
	salidas := []SalidasJoin{}
	if api.DB.Table("salidas").Select("salidas.id, salidas.idmedicamento, medicamentos.nombre as nombremedicamento, salidas.fechahora, salidas.cantidad, salidas.idfarmacia, farmacias.nik as nombrefarmacia").Joins("INNER JOIN medicamentos ON salidas.idmedicamento = medicamentos.id INNER JOIN farmacias ON salidas.idfarmacia = farmacias.id").Where("(fechahora BETWEEN ? AND ?)", fromDate+" 00:00:00", toDate+" 23:59:59").Scan(&salidas).Error != nil {
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

	salida.Fechahora = time.Now().Format("2006-01-02 15:04:05")

	if err := api.DB.Save(&salida).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	salidaOutput, err := getSalidaById(salida.Id)
	if err != nil {
		ResourceNotFound(w)
	}

	WriteJson(w, &salidaOutput)
}
