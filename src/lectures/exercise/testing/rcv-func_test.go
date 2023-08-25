// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import (
	"testing"
)

func TestAddHealth(t *testing.T) {
	q := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	q.addHealth(10)
	if q.health > q.MaxHealth {
		t.Errorf("Invalid health added, wanted %v but got %v", q.MaxHealth, q.health)
	}
}

func TestApplyDamage(t *testing.T) {
	q := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	q.applyDamage(99)
	if !(q.health >= q.MaxHealth) {
		t.Errorf("Please enter proper value, wanted: %v", q.health)
	}
}

func TestAddEnergy(t *testing.T) {
	var x uint
	q := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	x = 10
	totalEnergy := x + q.energy
	if q.energy > q.MaxEnergy {
		t.Errorf("Invalid player stat at start")
	}
	if totalEnergy > q.MaxEnergy {
		t.Errorf("Invalid stat to add to current energy pull")
	}
}

func TestConsumeEnergy(t *testing.T) {
	var x uint
	x = 20
	q := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	remEnergy := q.energy - x
	if remEnergy > q.MaxEnergy {
		t.Errorf("All energy has been used")
	}
}
