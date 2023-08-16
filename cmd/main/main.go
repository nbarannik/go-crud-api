package main

import (
	"github.com/gorilla/mux"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/config"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Printf("Run server on %s", config.DBConfig.Server)
	log.Fatal(http.ListenAndServe(config.DBConfig.Server, router))
}
