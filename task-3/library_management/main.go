package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mohaali482/a2sv-backend-learning-path/task-3/library_management/controllers"
	"github.com/mohaali482/a2sv-backend-learning-path/task-3/library_management/models"
	"github.com/mohaali482/a2sv-backend-learning-path/task-3/library_management/services"
)

func main() {
	library := controllers.NewLibrary()

	welcomeMessage()
	command := getCommand()
	for command != "0" {
		switch command {
		case "1":
			addBook(library)
		case "2":
			removeBook(library)
		case "3":
			borrowBook(library)
		case "4":
			returnBook(library)
		case "5":
			listAvailableBooks(library)
		case "6":
			listBorrowedBooks(library)
		default:
			fmt.Println("🚫🚫Invalid command🚫🚫")
		}

		command = getCommand()
	}

}

func welcomeMessage() {
	fmt.Println("Welcome to library manager")
	fmt.Println()
}

func inputReader() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func getCommand() string {
	fmt.Println()
	fmt.Println("Here are the list of commands.")
	fmt.Println()
	fmt.Println("(1) Add book")
	fmt.Println("(2) Remove book")
	fmt.Println("(3) Borrow book")
	fmt.Println("(4) Return book")
	fmt.Println("(5) List available books")
	fmt.Println("(6) List borrowed books")
	fmt.Println()
	fmt.Println("(0) Exit")
	fmt.Println()

	fmt.Print("Enter your command number: ")
	command := inputReader()
	return command
}

func addBook(l services.LibraryManager) {
	fmt.Println()
	fmt.Print("Enter book title: ")
	bookTitle := inputReader()
	fmt.Print("Enter book author: ")
	bookAuthor := inputReader()
	fmt.Println()

	book := models.Book{
		Title:  bookTitle,
		Author: bookAuthor,
	}

	l.AddBook(book)

	fmt.Println("✅ Done")
	fmt.Println()
}

func removeBook(l services.LibraryManager) {
	fmt.Println()
	fmt.Print("Enter book id: ")
	bookId := inputReader()
	intBookId, err := strconv.Atoi(bookId)

	for err != nil {
		fmt.Println("Invalid book id")
		fmt.Print("Enter book id: ")
		bookId = inputReader()
		intBookId, err = strconv.Atoi(bookId)
	}

	fmt.Println()

	l.RemoveBook(intBookId)

	fmt.Println()
}

func borrowBook(l services.LibraryManager) {
	fmt.Println()
	fmt.Print("Enter book id: ")
	bookId := inputReader()
	intBookId, err := strconv.Atoi(bookId)

	for err != nil {
		fmt.Println("Invalid book id")
		fmt.Print("Enter book id: ")
		bookId = inputReader()
		intBookId, err = strconv.Atoi(bookId)
	}
	fmt.Println()

	fmt.Print("Enter member id: ")
	memberId := inputReader()
	intMemberId, err := strconv.Atoi(memberId)

	for err != nil {
		fmt.Println("Invalid member id")
		fmt.Print("Enter member id: ")
		memberId = inputReader()
		intMemberId, err = strconv.Atoi(memberId)
	}

	fmt.Println()

	err = l.BorrowBook(intBookId, intMemberId)

	if err != nil {
		fmt.Println("🚫 " + err.Error())
	} else {
		fmt.Println("✅ Done")
	}

	fmt.Println()
}

func returnBook(l services.LibraryManager) {
	fmt.Println()
	fmt.Print("Enter book id: ")
	bookId := inputReader()
	intBookId, err := strconv.Atoi(bookId)

	for err != nil {
		fmt.Println("Invalid book id")
		fmt.Print("Enter book id: ")
		bookId = inputReader()
		intBookId, err = strconv.Atoi(bookId)
	}
	fmt.Println()

	fmt.Print("Enter member id: ")
	memberId := inputReader()
	intMemberId, err := strconv.Atoi(memberId)

	for err != nil {
		fmt.Println("Invalid member id")
		fmt.Print("Enter member id: ")
		memberId = inputReader()
		intMemberId, err = strconv.Atoi(memberId)
	}

	fmt.Println()

	err = l.ReturnBook(intBookId, intMemberId)

	if err != nil {
		fmt.Println("🚫 " + err.Error())
	} else {
		fmt.Println("✅ Done")
	}

	fmt.Println()
}

func listAvailableBooks(l services.LibraryManager) {
	fmt.Println()
	fmt.Println("Here is the list of available books.")
	fmt.Println()

	availableBooks := l.ListAvailableBooks()

	if len(availableBooks) == 0 {
		fmt.Println("🚫 No books available")
		fmt.Println()
		return
	}

	fmt.Printf("%-16s%-16s%-16s\n", "Book ID", "Title", "Author")
	for _, book := range availableBooks {
		fmt.Printf("%-16v%-16s%-16s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println()
}

func listBorrowedBooks(l services.LibraryManager) {
	fmt.Println()

	fmt.Print("Enter member id: ")
	memberId := inputReader()
	intMemberId, err := strconv.Atoi(memberId)

	for err != nil {
		fmt.Println("Invalid member id")
		fmt.Print("Enter member id: ")
		memberId = inputReader()
		intMemberId, err = strconv.Atoi(memberId)
	}
	fmt.Println()
	fmt.Println("Here is the list of borrowed books.")
	fmt.Println()

	borrowedBooks := l.ListBorrowedBooks(intMemberId)

	if len(borrowedBooks) == 0 {
		fmt.Println("🚫 No books borrowed")
		fmt.Println()
		return
	}

	fmt.Printf("%-16s%-16s%-16s\n", "Book ID", "Title", "Author")
	for _, book := range borrowedBooks {
		fmt.Printf("%-16v%-16s%-16s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println()
}
