package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getAllCategorias(w http.ResponseWriter, r *http.Request) {
	categorias := []Categorias{}
	api.DB.Find(&categorias)
	WriteJson(w, &categorias)
}

func getCategoria(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	categoria := Categorias{}
	if err := api.DB.First(&categoria, id).Error; err != nil {
		log.Error(err.Error())
		ResourceNotFound(w)
		return
	}
	WriteJson(w, &categoria)
}

func postCategoria(w http.ResponseWriter, r *http.Request) {
	categoria := Categorias{}
	if err := DecodeJson(r, &categoria); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&categoria).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &categoria)
}

func putCategoria(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	categoria := Categorias{}
	if api.DB.First(&categoria, id).Error != nil {
		ResourceNotFound(w)
		return
	}

	updated := Categorias{}
	if err := DecodeJson(r, &updated); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoria.Nombre = updated.Nombre
	categoria.Texto = updated.Texto
	categoria.Imagen = updated.Imagen

	if err := api.DB.Save(&categoria).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, &categoria)
}

func deleteCategoria(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	categoria := Categorias{}
	if api.DB.First(&categoria, id).Error != nil {
		ResourceNotFound(w)
		return
	}
	if err := api.DB.Delete(&categoria).Error; err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, "{}")
}
