package main

import (
	"io"
	"log"
	"net/http"
	"os"

	// Logging handler middleware package
	"github.com/gorilla/handlers"
)

func logHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(dst, h)
	}
}

func main() {
	// Record request
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	loggingHandler := logHandler(logFile)
	// Create server
	mux := http.NewServeMux()
	// Call handler function
	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", loggingHandler(finalHandler))

	log.Print("Listening on :4000...")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Create handler function
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successful login"))
}
