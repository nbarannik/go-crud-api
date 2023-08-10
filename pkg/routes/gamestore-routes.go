package routes

import (
	"github.com/gorilla/mux"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/controllers"
)

func RegisterGameStoreRoutes(router *mux.Router) {
	router.HandleFunc("/game/", controllers.GetGame).Methods("GET")
	router.HandleFunc("/game/", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/game/{gameId}", controllers.GetGameById).Methods("GET")
	router.HandleFunc("/game/{gameId}", controllers.UpdateGame).Methods("PUT")
	router.HandleFunc("/game/{gameId}", controllers.DeleteGame).Methods("DELETE")
}
