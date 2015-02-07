package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllMedicamentos(w http.ResponseWriter, r *http.Request) {
	medicamentos := []Medicamentos{}
	api.DB.Find(&medicamentos)
	WriteJson(w, &medicamentos)
}

func getMedicamento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &medicamento)
}

func postMedicamento(w http.ResponseWriter, r *http.Request) {
	medicamento := Medicamentos{}
	if err := DecodeJson(r, &medicamento); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&medicamento).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &medicamento)
}

func putMedicamento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		ResourceNotFound(w)
		return
	}

	updated := Medicamentos{}
	if err := DecodeJson(r, &updated); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	medicamento.Idcategoria = updated.Idcategoria
	medicamento.Nombre = updated.Nombre
	medicamento.Texto = updated.Texto
	medicamento.Imagen = updated.Imagen
	medicamento.Codigo = updated.Codigo
	medicamento.Enalmacen = updated.Enalmacen

	if err := api.DB.Save(&medicamento).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &medicamento)
}

func deleteMedicamento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		ResourceNotFound(w)
		return
	}
	if err := api.DB.Delete(&medicamento).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, "")
}

func sumaEnAlmancen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	quantity := vars["quantity"]

	enAlmacen, err := strconv.ParseInt(quantity, 10, 64)
	if err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		ResourceNotFound(w)
		return
	}

	medicamento.Enalmacen = medicamento.Enalmacen + enAlmacen

	if err := api.DB.Save(&medicamento).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &medicamento)
}
