package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.Connect()
	fmt.Println("Server running Go in port 8000")
	routes.HandleRequest()
}
