package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OmnGabriel/go-api-rest.git/database"
	"github.com/OmnGabriel/go-api-rest.git/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func AllCharacters(w http.ResponseWriter, r *http.Request) {

	var character []models.Character
	database.DB.Find(&character)
	json.NewEncoder(w).Encode(character)

}

func ReturnACharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var character models.Character

	database.DB.First(&character, id)
	json.NewEncoder(w).Encode(character)
}

func MakeANewCharacter(w http.ResponseWriter, r *http.Request) {
	var newCharacter models.Character
	json.NewDecoder(r.Body).Decode(&newCharacter)
	database.DB.Create(&newCharacter)
	json.NewEncoder(w).Encode(newCharacter)
}

func DeletACharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var character models.Character
	database.DB.Delete(&character, id)
	json.NewEncoder(w).Encode(character)

}

func EditACharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var character models.Character

	database.DB.First(&character, id)
	json.NewDecoder(r.Body).Decode(&character)
	database.DB.Save(&character)
	json.NewEncoder(w).Encode(character)
}
