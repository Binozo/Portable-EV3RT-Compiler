package routes

import (
	"github.com/gorilla/mux"
	"portable-ev3rt-compiler/server/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/compile", controllers.Compile).Methods("POST")
}
