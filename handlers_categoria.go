package main

import (
	"net/http"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *api) getAllCategorias(w rest.ResponseWriter, r *rest.Request) {
	categorias := []Categorias{}
	api.DB.Find(&categorias)
	w.WriteJson(&categorias)
}

func (api *api) getCategoria(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	categoria := Categorias{}
	if api.DB.First(&categoria, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&categoria)
}

func (api *api) postCategoria(w rest.ResponseWriter, r *rest.Request) {
	categoria := Categorias{}
	if err := r.DecodeJsonPayload(&categoria); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&categoria).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&categoria)
}

func (api *api) putCategoria(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	categoria := Categorias{}
	if api.DB.First(&categoria, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Categorias{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoria.Nombre = updated.Nombre
	categoria.Texto = updated.Texto
	categoria.Imagen = updated.Imagen

	if err := api.DB.Save(&categoria).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&categoria)
}

func (api *api) deleteCategoria(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	categoria := Categorias{}
	if api.DB.First(&categoria, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := api.DB.Delete(&categoria).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson("{}")
}
