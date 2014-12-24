package ws

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalFarmacia = Farmacias{Id: 1, Nik: "Milano", Pass: "milano01", Nivel: 0}
var testFarmacia = Farmacias{Nik: "Test", Pass: "test", Nivel: 666}

func TestGetAllFarmacias(t *testing.T) {
	resp := test("farmacias", "GET", "")

	isOK(resp, t)
}

func TestGetFarmacia(t *testing.T) {
	resp := test("farmacias/1", "GET", "")

	isOK(resp, t)

	responseStruct := Farmacias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedFarmaciaExpected(originalFarmacia, responseStruct, t)
}

func TestPostFarmacia(t *testing.T) {
	jsonBytes, _ := json.Marshal(testFarmacia)
	resp := test("farmacias", "POST", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Farmacias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedFarmaciaExpected(testFarmacia, responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testFarmacia = responseStruct
}

func TestPutFarmacia(t *testing.T) {
	updated := testFarmacia
	updated.Pass = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := test("farmacias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	isOK(resp, t)

	responseStruct := Farmacias{}
	decodeJsonPayload(resp, &responseStruct, t)
	isReturnedFarmaciaExpected(updated, responseStruct, t)
}

func TestDeleteFarmacia(t *testing.T) {
	resp := test("farmacias/"+strconv.FormatInt(testFarmacia.Id, 10), "DELETE", "")

	isOK(resp, t)
}

// Private functions

func isReturnedFarmaciaExpected(compareTo Farmacias, responseStruct Farmacias, t *testing.T) {
	if !compareFarmacias(compareTo, responseStruct) {
		a, _ := json.Marshal(responseStruct)
		b, _ := json.Marshal(compareTo)
		t.Errorf("Farmacia retrieved does not match the expected result. Got \n%s, expected \n%s", a, b)
	}
}
