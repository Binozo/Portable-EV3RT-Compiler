package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadZipToCompilationServer(zipFilePath string, destPath string) {
	client := &http.Client{}
	data, err := os.Open(zipFilePath)
	defer data.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:5321/api/compile", data)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	responseCode := resp.StatusCode
	if responseCode == 200 {
		err := os.WriteFile(filepath.Join(destPath, "app"), content, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
		}
	}

}
