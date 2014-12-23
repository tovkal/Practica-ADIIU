package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func test(operation string, method string, json string) (response *http.Response) {
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

func isOK(resp *http.Response, t *testing.T) {
	if resp.StatusCode != http.StatusOK {
		t.Error("Operation did not return OK, got ", resp.Status)
	}
}

func decodeJsonPayload(r *http.Response, v interface{}, t *testing.T) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("Error reading body from response.", err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(content, v)
	if err != nil {
		t.Error("Error unmarshaling response.", err)
	}
}
