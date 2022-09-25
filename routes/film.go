package routes

import (
	"be-dumbflix/handlers"
	"be-dumbflix/pkg/middleware"
	"be-dumbflix/pkg/mysql"
	"be-dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/films", h.FindFilms).Methods("GET")
	r.HandleFunc("/film/{id}", middleware.Auth(h.GetFilm)).Methods("GET")
	r.HandleFunc("/film/{id}", middleware.Auth(h.UpdateFilm)).Methods("PATCH")
	r.HandleFunc("/film/{id}", middleware.Auth(h.DeleteFilm)).Methods("DELETE")
}
