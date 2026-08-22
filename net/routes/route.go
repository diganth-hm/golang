package routes

import (
	"github.com/diganth-hm/golang/buildapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllmovies).Methods("GET")
	router.HandleFunc("/movies/add", controller.Creatmovie).Methods("POST")
	router.HandleFunc("/movies/delone/{id}", controller.DeleteAmovie).Methods("DELETE")
	router.HandleFunc("/movies/delAll", controller.DeleteAll).Methods("DELETE")
	router.HandleFunc("/movie/marksaswatched/{id}", controller.MarkAswatched).Methods("PUT")
	return router
}
