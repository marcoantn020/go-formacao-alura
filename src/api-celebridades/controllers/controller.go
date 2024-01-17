package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Home(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "Home Page")
}

func AllPersonalities(write http.ResponseWriter, request *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)

	err := json.NewEncoder(write).Encode(personalities)
	if err != nil {
		panic(err.Error())
	}
}

func GetOnePersonality(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var personality models.Personality

	database.DB.Find(&personality, id)
	err := json.NewEncoder(write).Encode(personality)
	if err != nil {
		log.Panicln("Erro ao busca uma personalidade", err)
	}
}

func Create(write http.ResponseWriter, request *http.Request) {
	var newPersonality models.Personality
	err := json.NewDecoder(request.Body).Decode(&newPersonality)
	if err != nil {
		log.Panicln("Houve um erro ao criar personalidade ", err)
	}
	database.DB.Create(&newPersonality)
	json.NewEncoder(write).Encode(newPersonality)
}

func Delete(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(write).Encode(personality)
}

func Update(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewDecoder(request.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(write).Encode(&personality)
}
