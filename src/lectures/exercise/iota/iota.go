//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import (
	"fmt"
)

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type operations int

func (op operations) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("Unhandled operations")
}

func main() {
	add := operations(Add)
	fmt.Println(add.calculate(2, 2)) // = 4
	sub := operations(Subtract)
	fmt.Println(sub.calculate(10, 3)) // = 7
	mul := operations(Multiply)
	fmt.Println(mul.calculate(3, 3)) // = 9
	div := operations(Divide)
	fmt.Println(div.calculate(100, 2)) // = 50
}

/*package main

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
