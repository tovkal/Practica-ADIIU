package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gorilla/context"
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
	makePath("login"),
))

//const staticPath = "http://staticadiiu.tovkal.com"
const staticPath = "../static"
const port = ":3000"

var api = Api{}

func main() {
	setupLogger()

	api.initDB()
	api.initSessionStore()

	// API
	router := NewRouter()

	log.Info("Ready to serve on %s!\n", port)
	log.Critical(http.ListenAndServe(port, context.ClearHandler(router)).Error())
}

func makePath(name string) string {
	return path.Join("tmpl", name+".html")
}
