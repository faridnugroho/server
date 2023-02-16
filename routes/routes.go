package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	RoleRoutes(r)
	MusicRoutes(r)
	ArtistRoutes(r)
}
