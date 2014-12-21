package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

var original = Categorias{Id: 1, Nombre: "Pediatría", Texto: "Productos para niños de 0 a 3 años", Imagen: "pediatria.png"}
var testStruct = Categorias{Nombre: "Nombre", Texto: "Descripción", Imagen: "image.png"}

func TestGetAllCategorias(t *testing.T) {
	resp := test("categorias", "GET", "")

	isOK(resp, t)
}

func TestGetCategoria(t *testing.T) {
	resp := test("categorias/1", "GET", "")

	isOK(resp, t)

	responseStruct := Categorias{}
	decodeJsonPayload(resp, &responseStruct, t)
	checkReturnedExpected(original, responseStruct, t)
}

func TestPostCategoria(t *testing.T) {
	jsonBytes, _ := json.Marshal(testStruct)
	resp := test("categorias", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Categorias{}
	decodeJsonPayload(resp, &responseStruct, t)
	checkReturnedExpected(testStruct, responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testStruct = responseStruct
}

func TestPutCategoria(t *testing.T) {
	updated := testStruct
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := test("categorias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Categorias{}
	decodeJsonPayload(resp, &responseStruct, t)
	checkReturnedExpected(updated, responseStruct, t)
}

func TestDeleteCategoria(t *testing.T) {
	resp := test("categorias/"+strconv.FormatInt(testStruct.Id, 10), "DELETE", "")

	isOK(resp, t)
}

// Private functions

func test(operation string, method string, json string) (response *http.Response) {
	url := "http://localhost:8080/" + operation
	fmt.Println("URL:>", url)

	var jsonStr = []byte(json)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func isOK(resp *http.Response, t *testing.T) {
	if resp.StatusCode != http.StatusOK {
		t.Error("Operation did not return OK, got ", resp.Status)
	}
}

func decodeJsonPayload(r *http.Response, v interface{}, t *testing.T) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("Error reading body from response.", err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(content, v)
	if err != nil {
		t.Error("Error unmarshaling response.", err)
	}
}

func checkReturnedExpected(compareTo Categorias, responseStruct Categorias, t *testing.T) {
	if !compareCategorias(compareTo, responseStruct) {
		a, _ := json.Marshal(compareTo)
		b, _ := json.Marshal(responseStruct)
		t.Errorf("Categoria retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}
