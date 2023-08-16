package models

import (
	"github.com/nbarannik/gameshop-go-crud-api/pkg/config"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Game struct {
	gorm.Model
	Title  string `gorm:"" json:"title"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Game{})
	if err != nil {
		log.Fatal("Failed to connect to db: %w", err)
	}
}

func (g *Game) CreateGame() *Game {
	db.Create(&g)
	return g
}

func GetGames() []Game {
	var games []Game
	db.Find(&games)
	return games
}

func GetGameById(Id int64) (*Game, *gorm.DB) {
	var game Game
	db := db.Where("ID=?", Id).Find(&game)
	return &game, db
}

func DeleteGame(Id int64) Game {
	var game Game
	db.Where("ID=?", Id).Delete(&game)
	return game
}
