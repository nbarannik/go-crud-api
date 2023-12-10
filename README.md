# Golang, MySQL CRUD API

Example of simple CRUD API using Golang, MySQL, [go-gorm/gorm](https://github.com/go-gorm/gorm) and [gorilla/mux](https://github.com/gorilla/mux). 

## Set up:

1. Clone the application  
```
git@github.com:nbarannik/gameshop-go-crud-api.git
```
2. Create MySQL database
```
CREATE DATABASE gameshop;
```
3. Change mysql username and password (username: "```root```", password: "```password```" by default)
   
 - Change ```var DBConfig``` in ```pkg/config/database.go```

4. Build and run

```
go run cmd/main/main.go
```

## CRUD APIs

| Endpoint  | HTTP Method | Result |
| ------------- | ------------- | ------------- |
| /game/ | GET  | Get all games |
| /game/  | POST  | Create a new game |
| /game/{gameId}  | GET  | Get a game by id |
| /game/{gameId}  | DELETE  | Delete a game by id |
| /game/{gameId}  | PUT  | Update a game by id |
