// Filename: main.go
package main

import (
	"log"
	"net/http"

	// A basic authentication third party middleware package
	"github.com/goji/httpauth"
)

func main() {
	// Create middleware with authorization username and password
	authHandler := httpauth.SimpleBasicAuth("joshua", "pa$$word")
	// Create Server
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", authHandler(finalHandler))

	log.Print("Listening on :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Create handler function
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successful login"))
}
