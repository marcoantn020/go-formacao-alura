package main

import (
	"alura/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.LoaderRoute()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Ocorreu um erro no servidor", err)
	}
}
