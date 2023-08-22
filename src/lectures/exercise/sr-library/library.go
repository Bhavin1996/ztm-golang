//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"

	"go.starlark.net/lib/time"
)

type Title string
type Name string

type Audit struct {
	checkOutTime time.Time
	checkInTime  time.Time
}

type Member struct {
	name  Name
	books map[Title]Audit
}

type BookEntry struct {
	total  int // total books owned
	lended int // total books lended out
}

type Library struct {
	member map[Name]Member
	books  map[Title]BookEntry
}

func printMemberAudit(menber *Member) {
	for title, audit := range menber.books {
		var returnTime string
		if audit.checkInTime.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkInTime.String()
		}
		fmt.Println(menber.name, ":", title, ":", audit.checkOutTime.String(), "through", returnTime)
	}
}

func printAllMemberAudit(library *Library) {
	for _, member := range library.member {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total : ", book.total, "/ lended : ", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not a part of library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more book available to lend")
		return false
	}
}

func main() {

}
