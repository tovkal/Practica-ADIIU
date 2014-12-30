package main

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalCategoria = Categorias{Id: 1, Nombre: "Pediatría", Texto: "Productos para niños de 0 a 3 años", Imagen: "pediatria.png"}
var categoriasList = [2]Categorias{originalCategoria, Categorias{Id: 2, Nombre: "Higiene", Texto: "Productos para la igiene corporal", Imagen: "higiene.png"}}
var testCategoria = Categorias{Nombre: "Nombre", Texto: "Descripción", Imagen: "image.png"}

func TestGetAllCategorias(t *testing.T) {
	resp := sendTest("categorias", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, categoriasList, t)
}

func TestGetCategoria(t *testing.T) {
	resp := sendTest("categorias/1", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, originalCategoria, t)
}

func TestPostCategoria(t *testing.T) {
	jsonBytes, _ := json.Marshal(testCategoria)
	resp := sendTest("categorias", "POST", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	response := Categorias{}
	decodeJsonPayload(resp, &response, t)

	// Save returned struct, which includes the id for later tests
	testCategoria = response
}

func TestPutCategoria(t *testing.T) {
	updated := testCategoria
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := sendTest("categorias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, updated, t)
}

func TestDeleteCategoria(t *testing.T) {
	resp := sendTest("categorias/"+strconv.FormatInt(testCategoria.Id, 10), "DELETE", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}
