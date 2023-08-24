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
package testing

import "testing"

func TestChangeStat(t *testing.T) {
	q := player{
		name:      "Kylo",
		health:    90,
		MaxHealth: 100,
		energy:    700,
		MaxEnergy: 1000,
	}
	if q.health > 100 {
		t.Errorf("Incorrect player health stat, want %v , got %v ", q.MaxHealth, q.health)
	}
	if q.energy > 1000 {
		t.Errorf("Invalid player energy stat")
	}
	if q.health <= 0 && q.energy <= 0 {
		t.Errorf("Invalid player stat parameter")
	}

}
