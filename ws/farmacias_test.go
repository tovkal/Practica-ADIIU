package ws

import (
	"encoding/json"
	"strconv"
	"testing"
)

var originalFarmacia = Farmacias{Id: 1, Nik: "Milano", Pass: "milano01", Nivel: 0}
var testFarmacia = Farmacias{Nik: "Test", Pass: "test", Nivel: 666}

func TestGetAllFarmacias(t *testing.T) {
	resp := sendTest("farmacias", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}

func TestGetFarmacia(t *testing.T) {
	resp := sendTest("farmacias/1", "GET", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, originalFarmacia, t)
}

func TestPostFarmacia(t *testing.T) {
	jsonBytes, _ := json.Marshal(testFarmacia)
	resp := sendTest("farmacias", "POST", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	responseStruct := Farmacias{}
	decodeJsonPayload(resp, &responseStruct, t)

	// Save returned struct, which includes the id for later tests
	testFarmacia = responseStruct
}

func TestPutFarmacia(t *testing.T) {
	updated := testFarmacia
	updated.Pass = "Updated"
	jsonBytes, _ := json.Marshal(updated)
	resp := sendTest("farmacias/"+strconv.FormatInt(updated.Id, 10), "PUT", string(jsonBytes))

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)

	isResponseExpected(resp, updated, t)
}

func TestDeleteFarmacia(t *testing.T) {
	resp := sendTest("farmacias/"+strconv.FormatInt(testFarmacia.Id, 10), "DELETE", "")

	CodeIs(resp, 200, t)
	ContentTypeIsJson(resp, t)
}
