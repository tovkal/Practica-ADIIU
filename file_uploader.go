package main

import (
	"io/ioutil"
	"net/http"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
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

	out, err := ioutil.ReadAll(file)

	uploadToBucket(header.Filename, out)

	// Read to file
	/*out, err := os.Create("./static/img/uploads/" + header.Filename)
	if err != nil {
		Error(w, "Unable to create the file. Check your write permissions", http.StatusInternalServerError)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	if _, err = io.Copy(out, file); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	response := fileResponse{}
	response.Status = "success"
	response.FileName = header.Filename

	WriteJson(w, response)
}

func uploadToBucket(fileName string, data []byte) {
	log.Notice("Going to upload image to S3: " + fileName)
	auth := aws.Auth{
		AccessKey: "ACCESS KEY HERE!!",
		SecretKey: "SECRET KEY HERE!!",
	}

	// EUCentral needs V4
	s := s3.New(auth, aws.EUWest)
	bucket := s.Bucket("BUCKET NAME HERE!!")

	err := bucket.Put("/img/uploads/"+fileName, data, "binary/octet-stream", s3.BucketOwnerFull, s3.Options{})
	if err != nil {
		log.Error(err.Error())
	}

	log.Notice("File upload successful")
}
