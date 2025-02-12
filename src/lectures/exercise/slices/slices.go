//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func showInfo(line []Part) {
	fmt.Println()
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Print(part, " ")
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	assemblyLine := make([]Part, x)
	assemblyLine[0] = "pipe"
	assemblyLine[1] = "bolt"
	assemblyLine[2] = "nut"
	showInfo(assemblyLine)
	assemblyLine = append(assemblyLine, "washer", "screw")
	showInfo(assemblyLine)
	newAssemblyLine := assemblyLine[3:]
	showInfo(newAssemblyLine)
}
