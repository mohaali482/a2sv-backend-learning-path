package models

import "errors"

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

func (m *Member) FindBook(bookId int) (int, error) {
	for idx, book := range m.BorrowedBooks {
		if book.ID == bookId {
			return idx, nil
		}
	}

	return 0, errors.New("book not borrowed by member")
}

func (m *Member) RemoveBook(idx int) (bool, error) {
	if idx < len(m.BorrowedBooks) {
		m.BorrowedBooks[idx], m.BorrowedBooks[len(m.BorrowedBooks)-1] = m.BorrowedBooks[len(m.BorrowedBooks)-1], m.BorrowedBooks[idx]
		m.BorrowedBooks = m.BorrowedBooks[:idx]
		return true, nil
	}

	return false, errors.New("index out of bound")
}
