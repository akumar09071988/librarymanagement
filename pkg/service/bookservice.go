package service

import (
	"librarymangement/pkg/dao"
	"librarymangement/pkg/models"
)

// BookService struct
type BookService struct{}

var bookDao = new(dao.BookDao)

//GetAll all books from dao
func (bs *BookService) GetAll() []models.Book {
	books, err := bookDao.GetAll()
	if err != nil {
		return []models.Book{}
	}
	return books
}

//Get book by id
func (bs *BookService) Get(id string) models.Book {
	book, err := bookDao.Get(id)
	if err != nil {
		return models.Book{}
	}
	return book
}

// Create crate a new book
func (bs *BookService) Create(book models.Book) error {
	err := bookDao.Create(book)
	if err == nil {
		return err
	}
	return nil
}

//Update book
func (bs *BookService) Update(book models.Book) error {
	err := bookDao.Update(book)
	if err == nil {
		return err
	}
	return nil
}
