package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func Compile(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("make", "app=helloev3").Output()
	if err != nil {
		jsonRes, _ := json.Marshal(map[string]string{"status": "error", "type": "Shell command error", "error": err.Error()})
		w.Write(jsonRes)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Compilation successful
	//send result back to client
	fileBytes, err := ioutil.ReadFile("app")

	if err != nil {
		jsonRes, _ := json.Marshal(map[string]string{"status": "error", "type": "Couldn't read compiled binary", "error": err.Error()})
		w.Write(jsonRes)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}
