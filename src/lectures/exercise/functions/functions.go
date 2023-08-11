//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var reader *bufio.Reader

func perName(name string) string {
	return name
}

func name() string {
	input, _ := reader.ReadString(',')
	return input
}

func welcome() string {
	input, _ := reader.ReadString('\n')
	return input
}

func add(x int, y int, z int) int {
	result := x + y + z
	return result
}

func randomSingleNumber() int {
	randomNumber := rand.Intn(101)
	return randomNumber
}

func randomDoubleNumber() (int, int) {
	num1 := rand.Intn(101)
	num2 := rand.Intn(100)
	return num1, num2
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	personName := perName(name())
	fmt.Printf("Hi, welcome %s to getting better at go", personName)
	fmt.Println(welcome())
	fmt.Println(add(1, 2, 4))
	num1 := randomSingleNumber()
	num2, num3 := randomDoubleNumber()
	fmt.Println(num1)
	fmt.Println(num2, num3)
	fmt.Println(add(num1, num2, num3))

}
