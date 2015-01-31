package main

import (
	"net/http"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *api) getAllNoticias(w rest.ResponseWriter, r *rest.Request) {
	noticias := []Noticias{}
	api.DB.Find(&noticias)
	w.WriteJson(&noticias)
}

func (api *api) getNoticia(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	noticias := Noticias{}
	if api.DB.Find(&noticias, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&noticias)
}

func (api *api) postNoticia(w rest.ResponseWriter, r *rest.Request) {
	noticia := Noticias{}
	if err := r.DecodeJsonPayload(&noticia); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&noticia).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&noticia)
}

func (api *api) putNoticia(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	noticia := Noticias{}
	if api.DB.First(&noticia, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Noticias{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	noticia.Texto = updated.Texto
	noticia.Inicio = updated.Inicio
	noticia.Fin = updated.Fin

	if err := api.DB.Save(&noticia).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&noticia)
}

func (api *api) getNoticiasFromDate(w rest.ResponseWriter, r *rest.Request) {
	date := r.PathParam("date")
	noticias := []Noticias{}
	if api.DB.Where("inicio <= ? and fin >= ?", date+" 00:00:00", date+" 23:59:59").First(&noticias).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&noticias)
}
