package handler

import (
	"io/fs"
	"io/ioutil"
	"log"
)

func GetFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	/*for _, f := range files {
		fmt.Println(f.Name())
	}*/
	return files
}
