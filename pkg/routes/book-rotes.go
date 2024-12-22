package routes

import (
	"github.com/gorilla/mux"
	"github.com/vipunsanjana/book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
}
