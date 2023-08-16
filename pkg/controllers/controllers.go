package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/models"
	"github.com/nbarannik/gameshop-go-crud-api/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

func SetResponse(w http.ResponseWriter, res []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		log.Printf("Failed to set response: %s", err)
	}
}

func GetGame(w http.ResponseWriter, _ *http.Request) {
	games := models.GetGames()
	res, _ := json.Marshal(games)
	SetResponse(w, res)
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
	SetResponse(w, res)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	gameModel := &models.Game{}
	err := utils.ParseBody(r, gameModel)
	if err != nil {

	}
	game := gameModel.CreateGame()
	res, _ := json.Marshal(game)
	SetResponse(w, res)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	id := ParseIdFromRequest(r)
	game := models.DeleteGame(id)
	res, _ := json.Marshal(game)
	SetResponse(w, res)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	updateGame := &models.Game{}
	err := utils.ParseBody(r, updateGame)
	if err != nil {
		http.Error(w, "Failed to update game", http.StatusInternalServerError)
		return
	}
	id := ParseIdFromRequest(r)
	curGame, db := models.GetGameById(id)
	if updateGame.Title != "" {
		curGame.Title = updateGame.Title
	}
	if updateGame.Author != "" {
		curGame.Author = updateGame.Author
	}
	db.Save(&curGame)
	res, _ := json.Marshal(curGame)
	SetResponse(w, res)
}
