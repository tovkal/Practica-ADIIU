package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getAllFarmacias(w http.ResponseWriter, r *http.Request) {
	farmacias := []Farmacias{}
	api.DB.Find(&farmacias)
	WriteJson(w, &farmacias)
}

func getFarmacia(w http.ResponseWriter, r *http.Request) {
	nik := mux.Vars(r)["nik"]
	farmacia, err := getFarmaciaByNik(nik)
	if err != nil {
		log.Error(err.Error())
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &farmacia)
}

func postFarmacia(w http.ResponseWriter, r *http.Request) {
	farmacia := Farmacias{}
	if err := DecodeJson(r, &farmacia); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&farmacia).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &farmacia)
}

func putFarmacia(w http.ResponseWriter, r *http.Request) {
	nik := mux.Vars(r)["nik"]
	farmacia := Farmacias{}
	if api.DB.First(&farmacia, nik).Error != nil {
		ResourceNotFound(w)
		return
	}

	updated := Farmacias{}
	if err := DecodeJson(r, &updated); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	farmacia.Nik = updated.Nik
	farmacia.Pass = updated.Pass
	farmacia.Nivel = updated.Nivel

	if err := api.DB.Save(&farmacia).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &farmacia)
}

func deleteFarmacia(w http.ResponseWriter, r *http.Request) {
	nik := mux.Vars(r)["nik"]
	farmacia := Farmacias{}
	if api.DB.First(&farmacia, nik).Error != nil {
		ResourceNotFound(w)
		return
	}
	if err := api.DB.Delete(&farmacia).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, "")
}
