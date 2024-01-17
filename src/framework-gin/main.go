package main

import (
	"gin/database"
	"gin/routes"
)

func main() {
	database.Connect()
	routes.HandlerRequests()
}
