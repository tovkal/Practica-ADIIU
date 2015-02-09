package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAllEntradas(w http.ResponseWriter, r *http.Request) {
	entradas := []EntradasJoin{}
	api.DB.Table("entradas").Select("entradas.id, entradas.idmedicamento, medicamentos.nombre as nombremedicamento, entradas.cantidad, entradas.fechahora").Joins("INNER JOIN medicamentos ON entradas.idmedicamento = medicamentos.id").Scan(&entradas)
	WriteJson(w, &entradas)
}

func getEntradas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromDate := vars["fromDate"]
	toDate := vars["toDate"]
	entradas := []EntradasJoin{}
	if api.DB.Table("entradas").Select("entradas.id, entradas.idmedicamento, medicamentos.nombre as nombremedicamento, entradas.cantidad, entradas.fechahora").Joins("INNER JOIN medicamentos ON entradas.idmedicamento = medicamentos.id").Where("(fechahora BETWEEN ? AND ?)", fromDate+" 00:00:00", toDate+" 23:59:59").Scan(&entradas).Error != nil {
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

	entrada.Fechahora = time.Now().Format("2006-01-02 15:04:05")

	if err := api.DB.Save(&entrada).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entradaOutput, err := getEntradaById(entrada.Id)
	if err != nil {
		ResourceNotFound(w)
	}

	WriteJson(w, &entradaOutput)
}
