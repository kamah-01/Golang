package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

    router.HandleFunc("/signup", SignUpHandler).Methods("POST")
    router.HandleFunc("/login", LoginHandler).Methods("POST")

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}