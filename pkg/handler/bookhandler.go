package handler

import (
	"encoding/json"
	"librarymangement/pkg/models"
	"librarymangement/pkg/service"
	"net/http"

	"github.com/go-chi/chi"
)

//BookHandler struct
type BookHandler struct{}

var bookService = new(service.BookService)

//GetAll gets all books from service
func (bh *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books := bookService.GetAll()
	json.NewEncoder(w).Encode(books)
}

//Get gets  book from service
func (bh *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	bookIDString := chi.URLParam(r, "bookID")
	book := bookService.Get(bookIDString)
	json.NewEncoder(w).Encode(book)
}

//Update gets  book from service
func (bh *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	error := bookService.Update(book)
	if error != nil {
		json.NewEncoder(w).Encode(error)
	}
	json.NewEncoder(w).Encode("Update successfull")
}

//Create gets  book from service
func (bh *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	error := bookService.Create(book)
	if error != nil {
		json.NewEncoder(w).Encode(error)
	}
	json.NewEncoder(w).Encode("Created successfull")
}
