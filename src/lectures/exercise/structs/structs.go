//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func RectAreaPeri(a rectangle) (int, int) {
	fmt.Scan(&a.length)
	fmt.Scan(&a.width)
	return (a.length * a.width), (a.length * 2) + (a.width * 2)
}

func DoubleRectAreaPeri(d rectangle) (int, int) {
	fmt.Scan(&d.length)
	fmt.Scan(&d.width)
	return 4 * (d.length * d.width), (2*(d.length*2) + 2*(d.width*2))
}

func main() {
	var area rectangle
	fmt.Println(RectAreaPeri(area))
	fmt.Println(DoubleRectAreaPeri(area))
}
