package main

import "fmt"

type Player struct {
	Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db := New[*Player](Options{Filename: "player.json", AutoIncrement: true})

	newPlayer := Player{Name: "Albino", Age: 30}

	fmt.Printf("New player %s\n", newPlayer.Name)

	db.Insert(&newPlayer)
	db.Save()
	newPlayer.Age = 31
	db.Update(&newPlayer)
}
