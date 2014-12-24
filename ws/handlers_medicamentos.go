package ws

import (
	"net/http"
	"strconv"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *Api) GetAllMedicamentos(w rest.ResponseWriter, r *rest.Request) {
	medicamentos := []Medicamentos{}
	api.DB.Find(&medicamentos)
	w.WriteJson(&medicamentos)
}

func (api *Api) GetMedicamento(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&medicamento)
}

func (api *Api) PostMedicamento(w rest.ResponseWriter, r *rest.Request) {
	medicamento := Medicamentos{}
	if err := r.DecodeJsonPayload(&medicamento); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&medicamento).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&medicamento)
}

func (api *Api) PutMedicamento(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Medicamentos{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	medicamento.Idcategoria = updated.Idcategoria
	medicamento.Nombre = updated.Nombre
	medicamento.Texto = updated.Texto
	medicamento.Imagen = updated.Imagen
	medicamento.Codigo = updated.Codigo
	medicamento.Enalmacen = updated.Enalmacen

	if err := api.DB.Save(&medicamento).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&medicamento)
}

func (api *Api) DeleteMedicamento(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := api.DB.Delete(&medicamento).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *Api) SumaEnAlmancen(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	quantity := r.PathParam("quantity")

	enAlmacen, err := strconv.ParseInt(quantity, 10, 64)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	medicamento := Medicamentos{}
	if api.DB.First(&medicamento, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	medicamento.Enalmacen = medicamento.Enalmacen + enAlmacen

	if err := api.DB.Save(&medicamento).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&medicamento)
}
