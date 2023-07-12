package routes

import (
	"log"
	"net/http"

	"github.com/OmnGabriel/go-api-rest.git/controllers"
	"github.com/OmnGabriel/go-api-rest.git/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/characters", controllers.AllCharacters).Methods("Get")
	r.HandleFunc("/api/characters/{id}", controllers.ReturnACharacter).Methods("Get")
	r.HandleFunc("/api/characters", controllers.MakeANewCharacter).Methods("Post")
	r.HandleFunc("/api/characters/{id}", controllers.DeletACharacter).Methods("Delete")
	r.HandleFunc("/api/characters/{id}", controllers.EditACharacter).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
