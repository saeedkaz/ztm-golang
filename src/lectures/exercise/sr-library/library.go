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

type Book struct {
	name         string
	checkout     bool
	checkouttime time.Time
	returntime   time.Time
}
type Member struct {
	name string
	book Book
}

func checkout(book *Book) {
	book.checkout = true
	book.checkouttime = time.Now()
}
func checkin(book *Book) {
	book.checkout = false
	book.returntime = time.Now()
}

func printbook(book Book) {
	fmt.Println()
	fmt.Println(book.name)
	fmt.Println(book.checkout)
	fmt.Println(book.checkouttime)
	fmt.Println(book.returntime)

}

func printamem(mem Member) {
	fmt.Println(mem.name)
	fmt.Println(mem.book)

}
func printLib(lib []Book) {
	fmt.Println()
	fmt.Println("printnig library!")
	for i := 0; i < len(lib); i++ {
		element := lib[i]
		printbook(element)
	}
	fmt.Println("Library Printed!")
	fmt.Println()

}

func printMem(mem []Member) {
	fmt.Println()
	fmt.Println("printnig members!")
	for i := 0; i < len(mem); i++ {
		element := mem[i]

		printamem(element)
	}
	fmt.Println("Mems Printed!")
	fmt.Println()

}
func main() {
	//p := fmt.Println
	//then := time.Now()
	//p(then)

	//printLib(lib)
	book1 := Book{}
	book2 := Book{}
	book3 := Book{}
	lib := []Book{book1, book2, book3}
	printLib(lib)
	book1.name = "book1"
	book2.name = "book2"
	book3.name = "book3"
	//book1.checkout = true
	checkin(&book1)
	checkout(&book2)
	lib = []Book{book1, book2, book3}
	printLib(lib)
	//	printbook(book1)
	//	printbook(book2)

	john := Member{"John", book1}
	mems := []Member{john}
	printMem(mems)
}
