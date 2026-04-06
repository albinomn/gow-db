package main

type Model struct {
	Id string `json:"id,string"`
}

type Entity interface {
	GetID() string
	SetID(string)
}

func (m Model) GetID() string {
	return m.Id
}

func (m *Model) SetID(id string) {
	m.Id = id
}
