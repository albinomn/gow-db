package main

import "fmt"

type Player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Model
}

func main() {
	db := NewDB[*Player]("player.json")

	newPlayer := Player{Name: "Albino", Age: 30}

	fmt.Printf("New player %s\n", newPlayer.Name)

	db.Push(&newPlayer)
	db.Save()
}
