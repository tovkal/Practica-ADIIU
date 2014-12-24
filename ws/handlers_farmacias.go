package ws

import (
	"net/http"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *Api) GetAllFarmacias(w rest.ResponseWriter, r *rest.Request) {
	farmacias := []Farmacias{}
	api.DB.Find(&farmacias)
	w.WriteJson(&farmacias)
}

func (api *Api) GetFarmacia(w rest.ResponseWriter, r *rest.Request) {
	nik := r.PathParam("nik")
	farmacia := Farmacias{}
	if api.DB.First(&farmacia, nik).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&farmacia)
}

func (api *Api) PostFarmacia(w rest.ResponseWriter, r *rest.Request) {
	farmacia := Farmacias{}
	if err := r.DecodeJsonPayload(&farmacia); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&farmacia).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&farmacia)
}

func (api *Api) PutFarmacia(w rest.ResponseWriter, r *rest.Request) {
	nik := r.PathParam("nik")
	farmacia := Farmacias{}
	if api.DB.First(&farmacia, nik).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Farmacias{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	farmacia.Nik = updated.Nik
	farmacia.Pass = updated.Pass
	farmacia.Nivel = updated.Nivel

	if err := api.DB.Save(&farmacia).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&farmacia)
}

func (api *Api) DeleteFarmacia(w rest.ResponseWriter, r *rest.Request) {
	nik := r.PathParam("nik")
	farmacia := Farmacias{}
	if api.DB.First(&farmacia, nik).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := api.DB.Delete(&farmacia).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
