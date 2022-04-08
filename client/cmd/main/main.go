package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"portable-ev3rt-compiler/client/pkg/handler"
)

func main() {
	dir, _ := filepath.Abs("./")
	finalExecutableTargetDir := dir
	if len(os.Args) >= 2 {
		dir = os.Args[1] // use path provided by command line
		finalExecutableTargetDir = dir
	}
	if len(os.Args) >= 3 {
		dir = os.Args[2] // use dest path provided by command line
	}
	fmt.Printf("Target Project directory is %v\n", dir)
	fmt.Println("Zipping project files...")
	zipDest := "zippedProject.zip"
	err := handler.ZipFiles(dir, zipDest)
	defer os.Remove(filepath.Join(dir, zipDest))

	if err != nil {
		log.Fatal("Error: couldn't zip files: ", err)
		return
	}
	fmt.Println("Zipping project files... Done")

	fmt.Println("Uploading to local compilation server...")
	handler.UploadZipToCompilationServer(zipDest, finalExecutableTargetDir)
	fmt.Println("Done. Executable file has been created")
}
