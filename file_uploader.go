package main

import (
	"io"
	"net/http"
	"os"
)

type fileResponse struct {
	Status   string `json:"status"`
	FileName string `json:"filename"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	out, err := os.Create("./static/img/uploads/" + header.Filename)
	if err != nil {
		Error(w, "Unable to create the file. Check your write permissions", http.StatusInternalServerError)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	if _, err = io.Copy(out, file); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fileResponse{}
	response.Status = "success"
	response.FileName = header.Filename

	WriteJson(w, response)
}
