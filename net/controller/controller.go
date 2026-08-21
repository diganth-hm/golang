package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diganth-hm/golang/buildapi/database"
)

func GetAllmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Getting all movies")
	allmovies := database.Getallmovies()
	json.NewEncoder(w).Encode(allmovies)
}
