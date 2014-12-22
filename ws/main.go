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
		// Categorias
		&rest.Route{"GET", "/categorias", api.GetAllCategorias},
		&rest.Route{"GET", "/categorias/:id", api.GetCategoria},
		&rest.Route{"POST", "/categorias", api.PostCategoria},
		&rest.Route{"PUT", "/categorias/:id", api.PutCategoria},
		&rest.Route{"DELETE", "/categorias/:id", api.DeleteCategoria},

		// Entradas
		&rest.Route{"GET", "/entradas", api.GetAllEntradas},
		&rest.Route{"GET", "/entradas/:fromDate/:toDate", api.GetEntrada},
		&rest.Route{"POST", "/entradas", api.PostEntrada},

		// Salidas
		&rest.Route{"GET", "/salidas", api.GetAllSalidas},
		&rest.Route{"GET", "/salidas/:fromDate/:toDate", api.GetSalida},
		&rest.Route{"POST", "/salidas", api.PostSalida},

		// Farmacias
		&rest.Route{"GET", "/farmacias", api.GetAllFarmacias},
		&rest.Route{"GET", "/farmacias/:nik", api.GetFarmacia},
		&rest.Route{"POST", "/farmacias", api.PostFarmacia},
		&rest.Route{"PUT", "/farmacias/:nik", api.PutFarmacia},
		&rest.Route{"DELETE", "/farmacias/:nik", api.DeleteFarmacia},

		// Medicamentos
		&rest.Route{"GET", "/medicamentos", api.GetAllMedicamentos},
		&rest.Route{"GET", "/medicamentos/:nombre", api.GetMedicamento},
		&rest.Route{"POST", "/medicamentos", api.PostMedicamento},
		&rest.Route{"PUT", "/medicamentos/:nombre", api.PutMedicamento},
		&rest.Route{"DELETE", "/medicamentos/:nombre", api.DeleteMedicamento},
		&rest.Route{"PUT", "/medicamentos/:nombre/stock/:quantity", api.SumaEnAlmancen},

		// Noticias
		&rest.Route{"GET", "/noticias", api.GetAllNoticias},
		&rest.Route{"GET", "/noticias/:id", api.GetNoticia},
		&rest.Route{"GET", "/noticias/date/:date", api.GetNoticiasFromDate},
		&rest.Route{"POST", "/noticias", api.PostNoticia},
		&rest.Route{"PUT", "/noticias/:id", api.PutNoticia},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Running")

	log.Fatal(http.ListenAndServe(":8080", &handler))
}
