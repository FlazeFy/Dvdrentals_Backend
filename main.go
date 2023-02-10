package main

import (
	"dvdrentals_backend/database"
	"dvdrentals_backend/routes"
)

func main() {
	database.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
