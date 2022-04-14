package routes

import (
	"github.com/gorilla/mux"
	"github.com/simox-83/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}
