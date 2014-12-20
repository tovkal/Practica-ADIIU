package main

import (
	"github.com/tovkal/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {

	api := Api{}
	api.InitDB()

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/categorias", api.GetAllCategorias},
		&rest.Route{"GET", "/categorias/:id", api.GetCategoria},
		&rest.Route{"POST", "/reminders", api.PostReminder},
		&rest.Route{"PUT", "/reminders/:id", api.PutReminder},
		&rest.Route{"DELETE", "/reminders/:id", api.DeleteReminder},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", &handler))
}
