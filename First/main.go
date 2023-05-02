// File Name: main
package main

import (
	"log"
	"net/http"
)

// Write middleware
func middlewareA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is executed going to the handler
		log.Println("Executing first middleware going to handler")
		// Showing flow by exiting after first middleware is executed
		if r.URL.Path == "/exit" {
			return
		}
		next.ServeHTTP(w, r)
		// This is executed going back to the client
		log.Println("Executing first middleware going back to client")
	})
}

func middlewareB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is executed going to the handler
		log.Println("Executing second middleware going to handler")
		// Showing flow by exiting after second middleware is executed
		/*if r.URL.Path == "/exit" {
			return
		}*/
		next.ServeHTTP(w, r)
		// This is executed going back to the client
		log.Println("Executing secong middleware going back to the client")
	})
}

// Create handler function
func testHandler(w http.ResponseWriter, r *http.Request) {
	// Arrived at handler
	log.Println("Executing the handler...")
	w.Write([]byte("Reached handler"))
}

func main() {
	// Create server
	mux := http.NewServeMux()
	// Call handler function
	mux.Handle("/", middlewareA(middlewareB(http.HandlerFunc(testHandler))))

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
