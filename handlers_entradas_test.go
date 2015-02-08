package main

import (
	"encoding/json"
	"testing"
)

var originalEntrada = Entradas{Id: 1, Idmedicamento: 1, Cantidad: 22, Fechahora: "2014-11-26 16:18:16"}
var testEntrada = Entradas{Idmedicamento: 1, Cantidad: 666}

func TestGetAllEntradas(t *testing.T) {
	resp := sendTest("entradas", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}

func TestGetEntrada(t *testing.T) {
	resp := sendTest("entradas/2014-11-25/2014-11-27", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, [1]Entradas{originalEntrada}, t)
}

func TestPostEntrada(t *testing.T) {
	jsonBytes, _ := json.Marshal(testEntrada)
	resp := sendTest("entradas", "POST", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}
