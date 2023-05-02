// File Name: main
package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

// Write middleware
func middlewareA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is executed going to the handler
		log.Println("Executing first middleware going to handler")

		next.ServeHTTP(w, r)
		// This is executed going back to the client
		log.Println("Executing first middleware going back to client")
	})
}

// Create handler function
func testHandler(w http.ResponseWriter, r *http.Request) {
	// Arrived at handler
	log.Println("Executing the handler...")
	w.Write([]byte("Reached handler"))
}

func main() {
	// Create middleware with authorization username and password
	authHandler := httpauth.SimpleBasicAuth("joshua", "pa$$word")
	finalHandler := http.HandlerFunc(final)

	// Create server
	mux := http.NewServeMux()

	// Call handler function
	mux.Handle("/", authHandler(finalHandler))
	mux.Handle("/next", middlewareA(http.HandlerFunc(testHandler)))

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Handler function that redirects to other handler function if login successful
func final(w http.ResponseWriter, r *http.Request) {
	redirectFunc(w, r)
}

// Redirect function
func redirectFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/next", http.StatusSeeOther)
}
