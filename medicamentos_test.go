package main

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalMedicamento = Medicamentos{Id: 1, Idcategoria: 1, Texto: "Antihistamínico con actividad antagonista selectiva sobre los receptores H1 periféricos.", Imagen: "allegra-r-pediatrico.png", Nombre: "Allegra ® pediátrico", Codigo: "000001", Enalmacen: 22}
var testMedicamento = Medicamentos{Idcategoria: 1, Texto: "Test.", Imagen: "allegra-r-pediatrico.png", Nombre: "Allegra ® pediátrico", Codigo: "000001", Enalmacen: 666}

func TestGetAllMedicamentos(t *testing.T) {
	resp := sendTest("medicamentos", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}

func TestGetMedicamento(t *testing.T) {
	resp := sendTest("medicamentos/1", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, originalMedicamento, t)
}

func TestPostMedicamento(t *testing.T) {
	jsonBytes, _ := json.Marshal(testMedicamento)
	resp := sendTest("medicamentos", "POST", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	responseStruct := Medicamentos{}
	decodeJsonPayload(resp, &responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testMedicamento = responseStruct
}

func TestPutMedicamento(t *testing.T) {
	updated := testMedicamento
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := sendTest("medicamentos/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, updated, t)

	testMedicamento = updated
}

func TestSumaEnAlmancen(t *testing.T) {
	jsonBytes, _ := json.Marshal(testMedicamento)
	resp := sendTest("medicamentos/"+strconv.FormatInt(testMedicamento.Id, 10)+"/stock/999", "PUT", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	testMedicamento.Enalmacen += 999

	isResponseExpected(resp, testMedicamento, t)
}

func TestDeleteMedicamento(t *testing.T) {
	resp := sendTest("medicamentos/"+strconv.FormatInt(testMedicamento.Id, 10), "DELETE", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}
