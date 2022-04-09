package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"portable-ev3rt-compiler/server/pkg/utils"
)

func Compile(w http.ResponseWriter, r *http.Request) {
	// Delete old zipped project file
	os.Remove("zippedProject.zip")
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	if err != nil {
		jsonRes, _ := json.Marshal(map[string]string{"status": "error", "type": "Couldn't parse uploaded zip file", "error": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonRes)
		return
	}
	defer file.Close() // Close the file when we finish

	// This is path which we want to store the file
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		jsonRes, _ := json.Marshal(map[string]string{"status": "error", "type": "Couldn't save the file", "error": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonRes)
		return
	}

	// Copy the file to the destination path
	io.Copy(f, file)

	// Clear the "project" directory
	os.RemoveAll("project")
	// Unzip the file
	err = utils.Unzip(handler.Filename, "./project")
	if err != nil {
		jsonRes, _ := json.Marshal(map[string]string{"status": "error", "type": "Couldn't unzip the project zip", "error": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonRes)
		return
	}

	out, err := exec.Command("make", "app=project").Output()
	if err != nil {
		jsonRes, _ := json.Marshal(map[string]string{"status": "error", "type": "Shell command error", "error": err.Error(), "output": string(out)})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonRes)
		return
	}
	fmt.Printf("Console output: %s\n", out)
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
