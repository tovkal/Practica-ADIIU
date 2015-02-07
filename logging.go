package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/op/go-logging"
)

// Global Logger
var log = logging.MustGetLogger("main")

var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05.000} %{shortfile}/%{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)
var fileFormat = logging.MustStringFormatter(
	"%{time:2006-01-02 15:04:05.000} %{shortfile}/%{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}",
)

func setupLogger() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	defer f.Close()

	// One backend for error saved in a file, another for console showing everything.
	backend1 := logging.NewLogBackend(f, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// Set formats for both backends
	backend1Formatter := logging.NewBackendFormatter(backend1, fileFormat)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
}

// Middleware logger
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Notice(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
