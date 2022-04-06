package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"portable-ev3rt-compiler/server/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
