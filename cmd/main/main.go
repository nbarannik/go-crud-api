package main

import (
	"github.com/gorilla/mux"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterGameStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:5000", router))
}
