package routes

import (
	"github.com/gorilla/mux"
	"github.com/vipunsanjana/book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
