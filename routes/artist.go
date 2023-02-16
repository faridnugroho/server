package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func ArtistRoutes(r *mux.Router) {
	artistRepository := repositories.RepositoryArtist(mysql.DB)
	h := handlers.HandlerArtist(artistRepository)

	r.HandleFunc("/artists", h.FindArtists).Methods("GET")
	r.HandleFunc("/artist/{id}", h.GetArtist).Methods("GET")
	r.HandleFunc("/artist", middleware.Auth(h.CreateArtist)).Methods("POST")
	r.HandleFunc("/artist/{id}", middleware.Auth(h.UpdateArtist)).Methods("PATCH")
	r.HandleFunc("/artist/{id}", middleware.Auth(h.DeleteArtist)).Methods("DELETE")
}
