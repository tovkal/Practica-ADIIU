package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func EncodeJson(v interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}

	return b, nil
}

func WriteHeader(w http.ResponseWriter, code int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}

	w.WriteHeader(code)
}

func WriteStatuslessJson(w http.ResponseWriter, v interface{}) error {
	b, err := EncodeJson(v)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func WriteJson(w http.ResponseWriter, v interface{}) error {
	WriteHeader(w, http.StatusOK)

	if err := WriteStatuslessJson(w, v); err != nil {
		return err
	}

	return nil
}

func Error(w http.ResponseWriter, error string, code int) {
	WriteHeader(w, code)

	err := WriteStatuslessJson(w, map[string]string{"error": error})
	if err != nil {
		log.Fatal(err)
	}

}

func ResourceNotFound(w http.ResponseWriter) {
	Error(w, "Resource not found", http.StatusNotFound)
}

func DecodeJson(r *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
