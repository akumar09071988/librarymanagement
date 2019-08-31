package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"librarymangement/pkg/models"
	"os"
	"path/filepath"
)

const dataSrc string = "../librarymangement/pkg/dummydata/book.json"

var bookList []models.Book

// BookDao struct
type BookDao struct{}

//GetAll gets all books
func (b *BookDao) GetAll() ([]models.Book, error) {
	return parseJSONFile(dataSrc)
}

//Get get book of id
func (b *BookDao) Get(id string) (models.Book, error) {
	books, error := parseJSONFile(dataSrc)
	if error != nil {
		return models.Book{}, errors.New("Unable to get")
	}
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("Unable to find book with")
}

//Create creates new book its temporary
func (b *BookDao) Create(book models.Book) error {
	books, error := parseJSONFile(dataSrc)
	if error != nil {
		return errors.New("Unable to add")
	}

	bookList = append(books, book)

	return nil
}

//Update to books list its temporary
func (b *BookDao) Update(book models.Book) error {
	books, error := parseJSONFile(dataSrc)
	if error != nil {
		return errors.New("Unable to update")
	}
	for index, innerbook := range books {
		if book.ID == innerbook.ID {
			books[index] = book
			break
		}
	}

	return nil
}

func parseJSONFile(path string) ([]models.Book, error) {
	if bookList == nil || len(bookList) == 0 {
		path, error := filepath.Abs(path)
		if error != nil {
			fmt.Println("abs file path invalid", error)
			return []models.Book{}, error
		}
		jsonFile, err := os.Open(path)
		defer jsonFile.Close()
		if err != nil {
			fmt.Println(err)
			return []models.Book{}, err
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &bookList)
		return bookList, nil
	}
	return bookList, nil
}

func saveToJSON(books []models.Book) {

}
