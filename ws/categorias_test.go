package ws

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalCategoria = Categorias{Id: 1, Nombre: "Pediatría", Texto: "Productos para niños de 0 a 3 años", Imagen: "pediatria.png"}
var testCategoria = Categorias{Nombre: "Nombre", Texto: "Descripción", Imagen: "image.png"}

func TestGetAllCategorias(t *testing.T) {
	resp := test("categorias", "GET", "")

	CodeIs(resp, 200, t)
}

func TestGetCategoria(t *testing.T) {
	resp := test("categorias/1", "GET", "")

	isOK(resp, t)

	response := Categorias{}
	decodeJsonPayload(resp, &response, t)
	isReturnedStructExpected(response, originalCategoria, t)
}

func TestPostCategoria(t *testing.T) {
	jsonBytes, _ := json.Marshal(testCategoria)
	resp := test("categorias", "POST", string(jsonBytes))

	isOK(resp, t)

	response := Categorias{}
	decodeJsonPayload(resp, &response, t)
	isReturnedStructExpected(response, testCategoria, t)

	// Save returned struct, which includes the id for later tests
	testCategoria = response
}

func TestPutCategoria(t *testing.T) {
	updated := testCategoria
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := test("categorias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	isOK(resp, t)

	response := Categorias{}
	decodeJsonPayload(resp, &response, t)
	isReturnedStructExpected(response, updated, t)
}

func TestDeleteCategoria(t *testing.T) {
	resp := test("categorias/"+strconv.FormatInt(testCategoria.Id, 10), "DELETE", "")

	isOK(resp, t)
}

// Private functions

func isReturnedStructExpected(response Categorias, expected Categorias, t *testing.T) {
	if !response.isEqualTo(&expected) {
		responseJson, _ := json.Marshal(response)
		expectedJson, _ := json.Marshal(expected)
		t.Errorf("Returned result does not match expected result. Got \n%s, expected \n%s\n", responseJson, expectedJson)
	}
}
