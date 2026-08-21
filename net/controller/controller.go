package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diganth-hm/golang/buildapi/database"
	model "github.com/diganth-hm/golang/buildapi/models"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Getting all movies")
	allmovies := database.Getallmovies()
	json.NewEncoder(w).Encode(allmovies)
}

func Creatmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//by the below header we are controlling the typpes of methods i allow
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	Err(err)
	database.InsertOnemovie(movie)
	json.NewEncoder(w).Encode(movie)

}
