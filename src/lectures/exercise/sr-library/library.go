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
	book.lended += 1
	library.books[title] = book

	member.books[title] = Audit{checkOutTime: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did nit chcek out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkInTime = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:  make(map[Title]BookEntry),
		member: make(map[Name]Member),
	}
	library.books["Webapps in GO"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The little GO book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn GO"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.member["Jayson"] = Member{"Jayson", make(map[Title]Audit)}
	library.member["Tyler"] = Member{"Tyler", make(map[Title]Audit)}
	library.member["Jean"] = Member{"Jean", make(map[Title]Audit)}

	fmt.Println("\nInitial :")
	printLibraryBooks(&library)
	printAllMemberAudit(&library)

	member := library.member["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book :")
	if checkedOut {
		printLibraryBooks(&library)
		printAllMemberAudit(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a book :")
	if returned {
		printLibraryBooks(&library)
		printAllMemberAudit(&library)
	}
}
