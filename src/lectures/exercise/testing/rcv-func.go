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
/*
package main

import "fmt"

type player struct {
	name              string
	health, MaxHealth uint
	energy, MaxEnergy uint
}

func (player *player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.MaxHealth {
		player.health = player.MaxHealth
	}
	fmt.Println(player.name, "Add", amount, "health ->", player.health)
}
func (player *player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "->", player.health)
}

func (player *player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.MaxEnergy {
		player.energy = player.MaxEnergy
	}
	fmt.Println(player.name, "Add", amount, "energy ->", player.energy)
}

func (player *player) consumeEnergy(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Energy consumed", amount, "->", player.energy)
}

func main() {
	newPlayer := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	newPlayer.applyDamage(99)
	newPlayer.addHealth(10)
	newPlayer.consumeEnergy(20)
	newPlayer.addEnergy(10)
	newPlayer.consumeEnergy(9999)

}

/*
package main

import (
    "fmt"
)

type direction int

const (
    North direction = iota
    East
    South
    West
)

func (d direction) String() string {
    switch d {
    case North:
        return "North"
    case East:
        return "East"
	case South:
        return "South"
    case West:
        return "West"
    default:
        return "Unknown"
    }
}

func main() {
    dir := East
    fmt.Println("Direction:", dir) // Uses the String() method to convert to a string
}
*/
