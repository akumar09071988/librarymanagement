package routing

import (
	"librarymangement/pkg/handler"
	"net/http"

	"github.com/go-chi/chi"
)

// Router struct to do router func
type Router struct{}

var bookHandler = new(handler.BookHandler)

func bookRouter(router *chi.Mux) *chi.Mux {
	router.Get("/{bookID}", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.Get(w, r)
	})

	router.Post("/update", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.Update(w, r)
	})
	router.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.Create(w, r)
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		bookHandler.GetAll(w, r)
	})
	return router
}

// Routes : initialize routes to all
func (r *Router) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/book", bookRouter(router))
	})
	return router
}
