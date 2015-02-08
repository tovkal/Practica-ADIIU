package main

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalNoticia = Noticias{Id: 1, Texto: "Soy una noticia que va del 16 al 25", Inicio: "2014-11-16", Fin: "2014-11-25"}
var testNoticia = Noticias{Texto: "Test", Inicio: "2015-02-08", Fin: "2015-03-02"}

func TestGetAllNoticias(t *testing.T) {
	resp := sendTest("noticias", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}

func TestGetNoticia(t *testing.T) {
	resp := sendTest("noticias/1", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, originalNoticia, t)
}

func TestPostNoticia(t *testing.T) {
	jsonBytes, _ := json.Marshal(testNoticia)
	resp := sendTest("noticias", "POST", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	responseStruct := Noticias{}
	decodeJsonPayload(resp, &responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testNoticia = responseStruct
}

func TestPutNoticia(t *testing.T) {
	updated := testNoticia
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := sendTest("noticias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, updated, t)
}
