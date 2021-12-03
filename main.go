package main

import (
	"webapi-with-go/database"
	"webapi-with-go/server"
)

func main() {
	database.StartDB()
	defer database.CloseConn()

	server := server.NewServer()
	server.Run()
}
