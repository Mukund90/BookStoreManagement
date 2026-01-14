package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Mukund90/book-store/pkg/models"
	"github.com/Mukund90/book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("🚀 Server starting on http://localhost:9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}
