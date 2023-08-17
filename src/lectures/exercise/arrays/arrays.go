//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type shoppingList struct {
	cost     int
	itemName string
}

func printInfo(list [4]shoppingList) {
	var cost, totalItems int
	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.cost
		if item.itemName != "" {
			totalItems += 1
		}
	}
	fmt.Println("The last item on list is", list[totalItems-1])
	fmt.Println("Total item in the list are", totalItems)
	fmt.Println("The total cost of items is", cost)
}

func main() {

	shoppingListArray := [4]shoppingList{
		{4, "kela"},
		{5, "baigan"},
		{7, "tamatar"},
	}
	printInfo(shoppingListArray)
	shoppingListArray[3] = shoppingList{5, "roti"}
	printInfo(shoppingListArray)
}
