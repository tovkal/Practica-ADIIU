package main

import (
	"net/http"
	"time"

	"github.com/tovkal/go-json-rest/rest"
)

type Reminder struct {
	Id        int64     `json:"id"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

func (api *Api) GetAllCategorias(w rest.ResponseWriter, r *rest.Request) {
	categorias := []Categorias{}
	api.DB.Find(&categorias)
	w.WriteJson(&categorias)
}

func (api *Api) GetCategoria(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	categoria := Categorias{}
	if api.DB.First(&categoria, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&categoria)
}

func (api *Api) PostCategoria(w rest.ResponseWriter, r *rest.Request) {
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

func (api *Api) PutCategoria(w rest.ResponseWriter, r *rest.Request) {

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

func (api *Api) DeleteCategoria(w rest.ResponseWriter, r *rest.Request) {
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
	w.WriteHeader(http.StatusOK)
}
