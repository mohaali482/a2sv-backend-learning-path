package controllers

import (
	"errors"

	"github.com/mohaali482/a2sv-backend-learning-path/task-3/library_management/models"
)

type Library struct {
	NextBookID int
	Books      map[int]models.Book
	Members    map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		NextBookID: 1,
		Books:      make(map[int]models.Book),
		Members:    make(map[int]models.Member),
	}
}

func (s *Library) AddBook(book models.Book) {
	book.ID = s.NextBookID
	book.Status = "available"
	s.Books[book.ID] = book
	s.NextBookID++
}

func (s *Library) RemoveBook(bookID int) {
	delete(s.Books, bookID)
}

func (s *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := s.Books[bookID]
	if !ok || book.Status != "available" {
		return errors.New("book is not available")
	}

	member, ok := s.Members[memberID]

	if !ok {
		member = models.Member{
			ID: memberID,
		}
		s.Members[memberID] = member
	}

	book.Status = "borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	s.Books[bookID] = book
	s.Members[memberID] = member

	return nil
}

func (s *Library) ReturnBook(bookID int, memberID int) error {
	member, ok := s.Members[memberID]
	if !ok {
		return errors.New("member is not registered to the system")
	}

	idx, err := member.FindBook(bookID)
	if err != nil {
		return err
	}

	book := member.BorrowedBooks[idx]
	book.Status = "available"
	s.Books[bookID] = book
	member.RemoveBook(idx)

	return nil
}

func (s *Library) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0)

	for _, book := range s.Books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

func (s *Library) ListBorrowedBooks(memberID int) []models.Book {
	borrowedBooks := make([]models.Book, 0)

	for _, book := range s.Books {
		if book.Status == "borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}

	return borrowedBooks
}
