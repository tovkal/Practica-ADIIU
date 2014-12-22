package main

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalCategoria = Categorias{Id: 1, Nombre: "Pediatría", Texto: "Productos para niños de 0 a 3 años", Imagen: "pediatria.png"}
var testCategoria = Categorias{Nombre: "Nombre", Texto: "Descripción", Imagen: "image.png"}

func TestGetAllCategorias(t *testing.T) {
	resp := test("categorias", "GET", "")

	isOK(resp, t)
}

func TestGetCategoria(t *testing.T) {
	resp := test("categorias/1", "GET", "")

	isOK(resp, t)

	responseStruct := Categorias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedCategoriaExpected(originalCategoria, responseStruct, t)
}

func TestPostCategoria(t *testing.T) {
	jsonBytes, _ := json.Marshal(testCategoria)
	resp := test("categorias", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Categorias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedCategoriaExpected(testCategoria, responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testCategoria = responseStruct
}

func TestPutCategoria(t *testing.T) {
	updated := testCategoria
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := test("categorias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Categorias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedCategoriaExpected(updated, responseStruct, t)
}

func TestDeleteCategoria(t *testing.T) {
	resp := test("categorias/"+strconv.FormatInt(testCategoria.Id, 10), "DELETE", "")

	isOK(resp, t)
}

// Private functions

func isReturnedCategoriaExpected(compareTo Categorias, responseStruct Categorias, t *testing.T) {
	if !compareCategorias(compareTo, responseStruct) {
		a, _ := json.Marshal(responseStruct)
		b, _ := json.Marshal(compareTo)
		t.Errorf("Categoria retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}
