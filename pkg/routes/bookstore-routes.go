package routes

import (
	"github.com/boxercool/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRouter = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

}
