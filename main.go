package main

import (
	"encoding/json"
	"log"
	"os"
)

type DB struct {
	filePath string
	file *os.File
}

func (d* DB) Create(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	d.file = f
}

func (d* DB) Close() {
	d.file.Close()
}
func main () {
	newDb := DB{filePath: "data.json"}
	f, err := os.Create(newDb.filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	ogMsg := map[string]map[string]string{"Data": {"Boa": "noite"}}
	msg, err := json.Marshal(ogMsg)
	_, err = f.WriteString(string(msg))

	if err != nil {
		log.Fatal(err)
	}
}
