package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var originalSalida = Salidas{Id: 3, Idmedicamento: 1, Cantidad: 5, Idfarmacia: 1}
var testSalida = Salidas{Idmedicamento: 1, Cantidad: 666, Idfarmacia: 1}

func TestGetAllSalidas(t *testing.T) {
	resp := test("salidas", "GET", "")

	isOK(resp, t)
}

func TestGetSalida(t *testing.T) {
	setup()

	resp := test("salidas/2014-11-29/2014-11-30", "GET", "")

	isOK(resp, t)

	responseStruct := []Salidas{}
	decodeJsonPayload(resp, &responseStruct, t)
}

func TestPostSalida(t *testing.T) {
	setup()

	jsonBytes, _ := json.Marshal(testSalida)
	resp := test("salidas", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Salidas{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedSalidaExpected(testSalida, responseStruct, t)
}

// Private functions

func isReturnedSalidaExpected(compareTo Salidas, responseStruct Salidas, t *testing.T) {
	if !compareSalidas(compareTo, responseStruct) {
		a, _ := json.Marshal(responseStruct)
		b, _ := json.Marshal(compareTo)
		t.Errorf("Salida retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}

func setup() {
	fechaHora, err := time.Parse("2006-01-02 15:04:05", "2014-11-29 14:10:04")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	originalSalida.Fechahora = fechaHora

	fechaHora, _ = time.Parse("2006-01-02 15:04:05", "2014-12-12 12:12:12")
	testSalida.Fechahora = fechaHora
}
