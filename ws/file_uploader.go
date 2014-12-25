package ws

import (
	"io"
	"net/http"
	"os"

	"github.com/tovkal/go-json-rest/rest"
)

func (api *Api) UploadHandler(w rest.ResponseWriter, r *rest.Request) {

	// the FormFile function takes in the POST input id file
	file, header, err := r.Request.FormFile("file")

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	out, err := os.Create("/files/images/" + header.Filename)
	if err != nil {
		rest.Error(w, "Unable to create the file for writing. Check your write access privilege", http.StatusInternalServerError)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson("File uploaded successfully : " + header.Filename)
	w.WriteHeader(http.StatusOK)
}
