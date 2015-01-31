package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/tovkal/go-json-rest/rest"
)

// Template caching
var templates = template.Must(template.ParseFiles(
	makePath("head"),
	makePath("footer"),
	makePath("header"),
	makePath("menu"),
	makePath("index"),
	makePath("inici"),
	makePath("categorias"),
	makePath("entradas"),
	makePath("salidas"),
	makePath("farmacias"),
	makePath("medicamentos"),
	makePath("noticias"),
))

// TODO walk tmpl directory amb build caches from what's found... but order...

const staticPath = "http://staticadiiu.tovkal.com"
const port = ":8080"

func main() {

	api := api{}
	api.initDB()

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		EnableStatusService:      true,
	}

	// API routes
	err := handler.SetRoutes(
		// Categorias
		&rest.Route{"GET", "/categorias", api.getAllCategorias},
		&rest.Route{"GET", "/categorias/:id", api.getCategoria},
		&rest.Route{"POST", "/categorias", api.postCategoria},
		&rest.Route{"PUT", "/categorias/:id", api.putCategoria},
		&rest.Route{"DELETE", "/categorias/:id", api.deleteCategoria},

		// Entradas
		&rest.Route{"GET", "/entradas", api.getAllEntradas},
		&rest.Route{"GET", "/entradas/:fromDate/:toDate", api.getEntrada},
		&rest.Route{"POST", "/entradas", api.postEntrada},

		// Salidas
		&rest.Route{"GET", "/salidas", api.getAllSalidas},
		&rest.Route{"GET", "/salidas/:fromDate/:toDate", api.getSalida},
		&rest.Route{"POST", "/salidas", api.postSalida},

		// Farmacias
		&rest.Route{"GET", "/farmacias", api.getAllFarmacias},
		&rest.Route{"GET", "/farmacias/:nik", api.getFarmacia},
		&rest.Route{"POST", "/farmacias", api.postFarmacia},
		&rest.Route{"PUT", "/farmacias/:nik", api.putFarmacia},
		&rest.Route{"DELETE", "/farmacias/:nik", api.deleteFarmacia},

		// Medicamentos
		&rest.Route{"GET", "/medicamentos", api.getAllMedicamentos},
		&rest.Route{"GET", "/medicamentos/:id", api.getMedicamento},
		&rest.Route{"POST", "/medicamentos", api.postMedicamento},
		&rest.Route{"PUT", "/medicamentos/:id", api.putMedicamento},
		&rest.Route{"DELETE", "/medicamentos/:id", api.deleteMedicamento},
		&rest.Route{"PUT", "/medicamentos/:id/stock/:quantity", api.sumaEnAlmancen},

		// Noticias
		&rest.Route{"GET", "/noticias", api.getAllNoticias},
		&rest.Route{"GET", "/noticias/:id", api.getNoticia},
		&rest.Route{"GET", "/noticias/date/:date", api.getNoticiasFromDate},
		&rest.Route{"POST", "/noticias", api.postNoticia},
		&rest.Route{"PUT", "/noticias/:id", api.putNoticia},

		//File upload
		&rest.Route{"POST", "/upload", api.uploadHandler},

		// Status
		&rest.Route{"GET", "/.status",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(handler.GetStatus())
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// API
	http.Handle("/api/", http.StripPrefix("/api", &handler))

	// Web pages
	http.HandleFunc("/", renderTemplate)

	log.Printf("Ready to serve on %s!\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving template with url = " + r.URL.Path)
	url := strings.TrimPrefix(r.URL.Path, "/")

	// For root, show the index page
	if len(url) == 0 {
		url = "index"
	}

	if err := templates.ExecuteTemplate(w, url, staticPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makePath(name string) string {
	return path.Join("tmpl", name+".html")
}
