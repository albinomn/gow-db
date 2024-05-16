package main

import (
	"fmt"
	"log"
	"os"
)

type gowDB struct {
	filePath string
	file *os.File
	Data []byte
}

func New(filePath string) gowDB {
	var f *os.File
	f, err := os.Open(filePath)
	if err != nil {
		f, err = os.Create(filePath)
	}
	return gowDB{filePath: filePath, file: f}
}

func (d* gowDB) Close() {
	d.file.Close()
}

func (d* gowDB) Write() {
	d.file.WriteString(string(d.Data))
}

func (d* gowDB) Read() {
	data, err := os.ReadFile(d.filePath)
	if err != nil {
		log.Fatal("File doesn't exist")
	}

	d.Data = data
}

func main () {
	newDb := New("data.json")

	defer newDb.Close()
	
	//ogMsg := map[string]map[string]string{"Data": {"Boa": "noite"}}
	//msg, err := json.Marshal(ogMsg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//newDb.Data = msg

	//newDb.Write()

	newDb.Read()

	fmt.Printf("%s", string(newDb.Data))
}
