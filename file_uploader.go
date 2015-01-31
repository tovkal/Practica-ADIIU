package main

import (
	"io"
	"net/http"
	"os"

	"github.com/tovkal/go-json-rest/rest"
)

type fileResponse struct {
	Status   string `json:"status"`
	FileName string `json:"filename"`
}

func (api *api) uploadHandler(w rest.ResponseWriter, r *rest.Request) {
	// the FormFile function takes in the POST input id file
	file, header, err := r.Request.FormFile("file")

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	out, err := os.Create("./static/img/uploads/" + header.Filename)
	if err != nil {
		rest.Error(w, "Unable to create the file. Check your write permissions", http.StatusInternalServerError)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	if _, err = io.Copy(out, file); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fileResponse{}
	response.Status = "success"
	response.FileName = header.Filename

	w.WriteJson(response)
}
