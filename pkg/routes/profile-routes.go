package routes

import (
	"github.com/gorilla/mux"
	"github.com/magana272/Section8/pkg/controllers"
)

var RegisterProfileRouter = func(r *mux.Router) {
	// r.Handle("/", controllers.IsAuthorized(controllers.Profile)).Methods("GET")
	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	r.HandleFunc("/Signup", controllers.Siginup).Methods("GET", "POST")
	r.HandleFunc("/Assistents", controllers.Assistents).Methods("GET")
	r.HandleFunc("/LogOut", controllers.LogOut)
	r.HandleFunc("/IncorrectLogin", controllers.IncorrectLogin)
	r.HandleFunc("/Profile", controllers.ProfilePage)
}
