package main

import (
	"log"
	"net/http"

	"github.com/boxercool/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))

}
