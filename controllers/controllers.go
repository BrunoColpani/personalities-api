package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personalities-api/database"
	"personalities-api/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func FindAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	fmt.Println("newPersonality", newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)

}
