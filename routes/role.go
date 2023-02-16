package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func RoleRoutes(r *mux.Router) {
	roleRepository := repositories.RepositoryRole(mysql.DB)
	h := handlers.HandlerRole(roleRepository)

	r.HandleFunc("/roles", h.FindRoles).Methods("GET")
	r.HandleFunc("/role/{id}", h.GetRole).Methods("GET")
	r.HandleFunc("/role", h.CreateRole).Methods("POST")
	r.HandleFunc("/role/{id}", h.UpdateRole).Methods("PATCH")
	r.HandleFunc("/role/{id}", h.DeleteRole).Methods("DELETE")
}
