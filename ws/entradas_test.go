package ws

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var originalEntrada = Entradas{Id: 1, Idmedicamento: 1, Cantidad: 22}
var testEntrada = Entradas{Idmedicamento: 1, Cantidad: 666}

func init() {
	fechaHora, err := time.Parse("2006-01-02 15:04:05", "2014-11-26 16:18:16")
	if err != nil {
		fmt.Println("Error %s", err)
	}
	originalEntrada.Fechahora = fechaHora

	fechaHora, _ = time.Parse("2006-01-02 15:04:05", "2014-12-12 12:12:12")
	testEntrada.Fechahora = fechaHora
}

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
