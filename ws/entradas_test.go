package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var originalEntrada = Entradas{Id: 1, Idmedicamento: 1, Cantidad: 22}
var testEntrada = Entradas{Idmedicamento: 1, Cantidad: 666}

func TestGetAllEntradas(t *testing.T) {
	resp := test("entradas", "GET", "")

	isOK(resp, t)
}

func TestGetEntrada(t *testing.T) {
	resp := test("entradas/2014-11-25/2014-11-27", "GET", "")

	isOK(resp, t)

	responseStruct := Entradas{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedEntradaExpected(originalEntrada, responseStruct, t)
}

func TestPostEntrada(t *testing.T) {
	jsonBytes, _ := json.Marshal(testEntrada)
	resp := test("entradas", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Entradas{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedEntradaExpected(testEntrada, responseStruct, t)
}

// Private functions

func isReturnedEntradaExpected(compareTo Entradas, responseStruct Entradas, t *testing.T) {
	if !compareEntradas(compareTo, responseStruct) {
		a, _ := json.Marshal(responseStruct)
		b, _ := json.Marshal(compareTo)
		t.Errorf("Entrada retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}

func setup() {
	fechaHora, err := time.Parse("2006-01-02 15:04:05", "2014-11-26 16:18:16")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	originalEntrada.Fechahora = fechaHora

	fechaHora, _ = time.Parse("2006-01-02 15:04:05", "2014-12-12 12:12:12")
	testEntrada.Fechahora = fechaHora
}
