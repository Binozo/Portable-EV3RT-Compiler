package controllers

import "net/http"

func Compile(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}