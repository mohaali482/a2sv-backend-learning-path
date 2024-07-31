package controllers

import (
	"errors"

	"github.com/mohaali482/a2sv-backend-learning-path/task-3/library_management/models"
)

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func (s *Library) AddBook(book models.Book) {
	s.Books[book.ID] = book
}

func (s *Library) RemoveBook(bookID int) {
	delete(s.Books, bookID)
}

func (s *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := s.Books[bookID]
	if !ok {
		return errors.New("book is not available")
	}

	member, ok := s.Members[memberID]

	if !ok {
		member = models.Member{
			ID: memberID,
		}
		s.Members[memberID] = member
	}

	member.BorrowedBooks = append(member.BorrowedBooks, book)

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
	member.RemoveBook(idx)

	s.Books[book.ID] = book

	return nil
}

func (s *Library) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0)

	for _, book := range s.Books {
		availableBooks = append(availableBooks, book)
	}

	return availableBooks
}

func (s *Library) ListBorrowedBooks(memberID int) []models.Book {
	if _, ok := s.Members[memberID]; !ok {
		return []models.Book{}
	}

	return s.Members[memberID].BorrowedBooks
}
