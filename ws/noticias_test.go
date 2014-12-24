package ws

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

var originalNoticia = Noticias{Id: 1, Texto: "Soy una noticia que va del 16 al 25"}
var testNoticia = Noticias{Texto: "Test"}

func init() {
	fechaHora, err := time.Parse("2006-01-02 15:04:05", "2014-11-16 00:00:00")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	originalNoticia.Inicio = fechaHora

	fechaHora, err = time.Parse("2006-01-02 15:04:05", "2014-11-25 00:00:00")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	originalNoticia.Fin = fechaHora

	fechaHora, err = time.Parse("2006-01-02 15:04:05", "2014-11-25 00:00:00")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	testNoticia.Inicio = fechaHora

	fechaHora, err = time.Parse("2006-01-02 15:04:05", "2014-11-25 00:00:00")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	testNoticia.Fin = fechaHora
}

func TestGetAllNoticias(t *testing.T) {
	resp := test("noticias", "GET", "")

	isOK(resp, t)
}

func TestGetNoticia(t *testing.T) {
	resp := test("noticias/1", "GET", "")

	isOK(resp, t)
}

func TestPostNoticia(t *testing.T) {
	jsonBytes, _ := json.Marshal(testNoticia)
	resp := test("noticias", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Noticias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedNoticiaExpected(testNoticia, responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testNoticia = responseStruct
}

func TestPutNoticia(t *testing.T) {
	updated := testNoticia
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := test("noticias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Noticias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedNoticiaExpected(updated, responseStruct, t)
}

// Private functions

func isReturnedNoticiaExpected(compareTo Noticias, responseStruct Noticias, t *testing.T) {
	if !compareNoticias(compareTo, responseStruct) {
		a, _ := json.Marshal(responseStruct)
		b, _ := json.Marshal(compareTo)
		t.Errorf("Noticia retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}
