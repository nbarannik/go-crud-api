package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/models"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/utils"
	"net/http"
	"strconv"
)

var NewGame models.Game //Todo: delete

func GetGame(w http.ResponseWriter, r *http.Request) {
	games := models.GetGames()
	res, _ := json.Marshal(games)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func ParseIdFromRequest(r *http.Request) int64 {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	id, _ := strconv.ParseInt(gameId, 0, 0)
	return id
}

func GetGameById(w http.ResponseWriter, r *http.Request) {
	id := ParseIdFromRequest(r)
	game, _ := models.GetGameById(id)
	res, _ := json.Marshal(game)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	gameModel := &models.Game{}
	utils.ParseBody(r, gameModel)
	game := gameModel.CreateGame()
	res, _ := json.Marshal(game)
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json") ???
	w.Write(res)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	id := ParseIdFromRequest(r)
	game := models.DeleteGame(id)
	res, _ := json.Marshal(game)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
