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
))

func main() {

	api := Api{}
	api.InitDB()

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		EnableStatusService:      true,
	}

	// API routes
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
		&rest.Route{"GET", "/medicamentos/:id", api.GetMedicamento},
		&rest.Route{"POST", "/medicamentos", api.PostMedicamento},
		&rest.Route{"PUT", "/medicamentos/:id", api.PutMedicamento},
		&rest.Route{"DELETE", "/medicamentos/:id", api.DeleteMedicamento},
		&rest.Route{"PUT", "/medicamentos/:id/stock/:quantity", api.SumaEnAlmancen},

		// Noticias
		&rest.Route{"GET", "/noticias", api.GetAllNoticias},
		&rest.Route{"GET", "/noticias/:id", api.GetNoticia},
		&rest.Route{"GET", "/noticias/date/:date", api.GetNoticiasFromDate},
		&rest.Route{"POST", "/noticias", api.PostNoticia},
		&rest.Route{"PUT", "/noticias/:id", api.PutNoticia},

		//File upload
		&rest.Route{"POST", "/upload", api.UploadHandler},

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
	http.Handle("/img/uploads/", http.FileServer(http.Dir("static/")))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	log.Print("Ready to serve!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving template with url = " + r.URL.Path)
	url := strings.TrimPrefix(r.URL.Path, "/")

	// For root, show the index page
	if len(url) == 0 {
		url = "index"
	}

	if err := templates.ExecuteTemplate(w, url, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makePath(name string) string {
	return path.Join("tmpl", name+".html")
}
