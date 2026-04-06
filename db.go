package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrNotFound = errors.New("Record not found")
var ErrAlreadyExists = errors.New("Record already exists")

type DB[T Entity] struct {
	LastID   int          `json:"lastID,string,omitempty"`
	Filename string       `json:"-"`
	Data     map[string]T `json:"data"`
}

type Options struct {
	Filename      string
	AutoIncrement bool
}

func New[T Entity](opt Options) *DB[T] {
	data, err := os.ReadFile(opt.Filename)

	if err != nil {
		if os.IsNotExist(err) {
			if opt.AutoIncrement {
				return &DB[T]{
					LastID:   0,
					Filename: opt.Filename,
					Data:     make(map[string]T),
				}
			}
			return &DB[T]{
				Filename: opt.Filename,
				Data:     make(map[string]T),
				LastID:   0,
			}
		}
		panic(err)
	}

	var decodedData DB[T]

	if err := json.Unmarshal(data, &decodedData); err != nil {
		panic(err)
	}

	decodedData.Filename = opt.Filename

	return &decodedData
}

func (d *DB[T]) Insert(data T) error {
	dataId := data.GetID()
	if _, exists := d.Data[dataId]; exists {
		return ErrAlreadyExists
	}

	if d.LastID >= 0 {
		if dataId == "" {
			d.LastID += 1
			data.SetID(strconv.Itoa(d.LastID))
			dataId = data.GetID()
		} else {
			convertedId, err := strconv.Atoi(dataId)
			if err != nil {
				return err
			}
			if d.LastID < convertedId {
				d.LastID = convertedId
			}
		}
	}

	d.Data[dataId] = data
	return nil
}

func (d DB[T]) Save() {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	os.WriteFile(d.Filename, bytes, 0644)
}

func (d *DB[T]) Update(newData T) error {
	id := newData.GetID()
	for i, oldData := range d.Data {
		if id == i {
			oldData = newData
		}
	}
	d.Save()
	return nil
}
