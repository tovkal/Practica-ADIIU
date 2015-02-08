package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getAllNoticias(w http.ResponseWriter, r *http.Request) {
	noticias := []Noticias{}
	api.DB.Find(&noticias)
	WriteJson(w, &noticias)
}

func getNoticia(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	noticias := Noticias{}
	if api.DB.Find(&noticias, id).Error != nil {
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &noticias)
}

func postNoticia(w http.ResponseWriter, r *http.Request) {
	noticia := Noticias{}
	if err := DecodeJson(r, &noticia); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&noticia).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &noticia)
}

func putNoticia(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	noticia := Noticias{}
	if api.DB.First(&noticia, id).Error != nil {
		ResourceNotFound(w)
		return
	}

	updated := Noticias{}
	if err := DecodeJson(r, &updated); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	noticia.Texto = updated.Texto
	noticia.Inicio = updated.Inicio
	noticia.Fin = updated.Fin

	if err := api.DB.Save(&noticia).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &noticia)
}

func getNoticiasFromDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	noticias := []Noticias{}
	if api.DB.Where("inicio <= ? and fin >= ?", date+" 00:00:00", date+" 23:59:59").Find(&noticias).Error != nil {
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &noticias)
}
