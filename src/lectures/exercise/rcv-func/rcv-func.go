//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type player struct {
	name              string
	health, MaxHealth int
	energy, MaxEnergy int
}

func stats(data *player) {
	fmt.Println(*data)
}

func (data *player) changeStat(x, y int) {
	data.health += x
	data.energy += y
	fmt.Println("The new health is :", data.health, "\nThe energy is :", data.energy)
}

func main() {
	newPlayer := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	stats(&newPlayer)
	newPlayer.changeStat(5, 100)
}
