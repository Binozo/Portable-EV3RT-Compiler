package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadZipToCompilationServer(zipFilePath string, destPath string) {
	client := &http.Client{}
	file, err := os.Open(zipFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = io.Copy(part, file)
	writer.Close()
	req, err := http.NewRequest("POST", "http://localhost:5321/api/compile", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server Error: %v", string(content))
		return
	}
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
