package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Send request to API, given the operation (one of the routes), the method (PUT, GET, etc.)
// and a json string
func sendTest(operation string, method string, json string) (response *http.Response) {
	url := "http://localhost:8080/" + operation
	fmt.Printf("Method: %s, URL: %s\n", method, url)

	var jsonStr = []byte(json)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

// Check if the status code from the response matches the expected one
func CodeIs(resp *http.Response, expectedCode int, t *testing.T) {
	if resp.StatusCode != expectedCode {
		t.Errorf("Code %d expected, got: %d", expectedCode, resp.StatusCode)
	}
}

// HeaderIs tests the first value for the given headerKey
func HeaderIs(r *http.Response, headerKey, expectedValue string, t *testing.T) {
	value := r.Header.Get(headerKey)
	if value != expectedValue {
		t.Errorf(
			"%s: %s expected, got: %s",
			headerKey,
			expectedValue,
			value,
		)
	}
}

// CHeck if the response's content type is a json with utf8 charset
func ContentTypeIsJson(r *http.Response, t *testing.T) {
	HeaderIs(r, "Content-Type", "application/json; charset=utf8", t)
}

// Gets the body from the http response
func getResponseBody(r *http.Response, t *testing.T) []byte {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("Error reading body from response.", err)
	}
	return content
}

// Decodes the json payload from a http response, and puts it in the struct
func decodeJsonPayload(r *http.Response, v interface{}, t *testing.T) {
	content := getResponseBody(r, t)
	err := json.Unmarshal(content, v)
	if err != nil {
		t.Error("Error unmarshaling response.", err)
	}
}

// Checks if the http response matches the expected struct
func isResponseExpected(r *http.Response, expected interface{}, t *testing.T) {
	responseBody := getResponseBody(r, t)
	defer r.Body.Close()

	var expectedJson []byte
	var err error

	if expectedJson, err = json.MarshalIndent(expected, "", "  "); err != nil {
		t.Errorf("Error marshalling the expected value: \n%s", expected)
	}

	if bytes.Compare(responseBody, expectedJson) != 0 {
		t.Errorf("The response doesn't match the expected result. Got \n%s\n, but expected:\n%s\n", responseBody, expectedJson)
	}
}
