package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.SetContentTypeMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	router.HandleFunc("/personalities/{id}", controllers.GetOnePersonality).Methods("Get")
	router.HandleFunc("/personalities", controllers.Create).Methods("Post")
	router.HandleFunc("/personalities/{id}", controllers.Delete).Methods("Delete")
	router.HandleFunc("/personalities/{id}", controllers.Update).Methods("Put")

	log.Fatalln(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
