package main

import (
	"challenge-12/database"
	"challenge-12/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8000")
}