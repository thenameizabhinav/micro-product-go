package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()

	// route for sign up and signin. The Function will come from auth-service package
	mainRouter.HandleFunc("/signup", nil)

	// The Signin will send the JWT Token back as we are making microservices.
	// JWT token will make sure that other services are protected.
	// So, ultimately, we would need a middleware
	mainRouter.HandleFunc("/signin", nil)

	// HTTP Server
	// Add Time outs
	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: mainRouter,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Booting the Server")
	}
}
