package controllers

import (
	"alura/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	products := models.SearchAllProducts()
	err := templates.ExecuteTemplate(writer, "Index", products)
	if err != nil {
		fmt.Println("Ocorreu um erro ao executar o template", err)
	}
}

func New(writer http.ResponseWriter, request *http.Request) {
	err := templates.ExecuteTemplate(writer, "New", nil)
	if err != nil {
		fmt.Println("Ocorreu um erro ao executar o template", err)
	}
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		name := request.FormValue("name")
		description := request.FormValue("description")
		price := request.FormValue("price")
		quantity := request.FormValue("quantity")

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Panicln("erro na coversao do price: ", err)
		}

		convertQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Panicln("erro na coversao da quantity: ", err)
		}

		models.CreateNewProduct(name, description, convertPrice, convertQuantity)
	}
	http.Redirect(writer, request, "/", 302)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	models.Delete(id)
	http.Redirect(writer, request, "/", 302)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	product := models.FindById(id)
	err := templates.ExecuteTemplate(writer, "Edit", product)
	if err != nil {
		fmt.Println("Ocorreu um erro ao executar o template", err)
	}
}

func Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id := request.FormValue("id")
		name := request.FormValue("name")
		description := request.FormValue("description")
		price := request.FormValue("price")
		quantity := request.FormValue("quantity")

		convertId, err := strconv.Atoi(id)
		if err != nil {
			log.Panicln("erro na coversao do id: ", err)
		}

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Panicln("erro na coversao do price: ", err)
		}

		convertQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Panicln("erro na coversao da quantity: ", err)
		}

		models.Update(convertId, name, description, convertPrice, convertQuantity)
	}
	http.Redirect(writer, request, "/", 302)
}
