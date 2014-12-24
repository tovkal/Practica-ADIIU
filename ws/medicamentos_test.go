package ws

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalMedicamento = Medicamentos{Id: 1, Idcategoria: 1, Texto: "Antihistamínico con actividad antagonista selectiva sobre los receptores H1 periféricos.", Imagen: "allegra-r-pediatrico.png", Nombre: "Allegra ® pediátrico", Codigo: "000001", Enalmacen: 22}
var testMedicamento = Medicamentos{Idcategoria: 1, Texto: "Test.", Imagen: "allegra-r-pediatrico.png", Nombre: "Allegra ® pediátrico", Codigo: "000001", Enalmacen: 666}

func TestGetAllMedicamentos(t *testing.T) {
	resp := test("medicamentos", "GET", "")

	isOK(resp, t)
}

func TestGetMedicamento(t *testing.T) {
	resp := test("medicamentos/1", "GET", "")

	isOK(resp, t)

	responseStruct := Medicamentos{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedMedicamentoExpected(originalMedicamento, responseStruct, t)
}

func TestPostMedicamento(t *testing.T) {
	jsonBytes, _ := json.Marshal(testMedicamento)
	resp := test("medicamentos", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Medicamentos{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedMedicamentoExpected(testMedicamento, responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testMedicamento = responseStruct
}

func TestPutMedicamento(t *testing.T) {
	updated := testMedicamento
	updated.Texto = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := test("medicamentos/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Medicamentos{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedMedicamentoExpected(updated, responseStruct, t)
}

func TestSumaEnAlmancen(t *testing.T) {
	jsonBytes, _ := json.Marshal(testMedicamento)
	resp := test("medicamentos/"+strconv.FormatInt(testMedicamento.Id, 10)+"/stock/999", "PUT", string(jsonBytes))

	isOK(resp, t)
}

func TestDeleteMedicamento(t *testing.T) {
	resp := test("medicamentos/"+strconv.FormatInt(testMedicamento.Id, 10), "DELETE", "")

	isOK(resp, t)
}

// Private functions

func isReturnedMedicamentoExpected(compareTo Medicamentos, responseStruct Medicamentos, t *testing.T) {
	if !compareMedicamentos(compareTo, responseStruct) {
		a, _ := json.Marshal(responseStruct)
		b, _ := json.Marshal(compareTo)
		t.Errorf("Medicamento retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}
