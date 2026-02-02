package main

import (
	"encoding/json"
	"os"
	"strconv"
)

type DB[T Entity] struct {
	LastID   int    `json:"lastID,string"`
	Filename string `json:"-"`
	Data     []T    `json:"data"`
}

func NewDB[T Entity](name string) *DB[T] {
	data, err := os.ReadFile(name)

	var decodedData DB[T]

	if err != nil {
		if os.IsNotExist(err) {
			return &DB[T]{
				LastID:   0,
				Filename: name,
				Data:     make([]T, 0),
			}
		}
		panic(err)
	}
	if err := json.Unmarshal(data, &decodedData); err != nil {
		panic(err)
	}

	decodedData.Filename = name

	return &decodedData
}

func (d *DB[T]) Push(data T) {
	i, err := strconv.Atoi(data.GetID())
	if err != nil && i == 0 {
		if d.LastID == 0 {
			d.LastID = 1
			data.SetID(strconv.Itoa(d.LastID))
		} else {
			d.LastID = d.LastID + 1
			data.SetID(strconv.Itoa(d.LastID))
		}
	}
	d.Data = append(d.Data, data)
}

func (d *DB[T]) Save() {
	bytes, err := json.Marshal(&d)
	if err != nil {
		panic(err)
	}
	os.WriteFile(d.Filename, bytes, 0644)
}
