package routes

import (
	"github.com/gorilla/mux"
	"github.com/magana272/Section8/pkg/controllers"
)

var RegisterLocationRouter = func(r *mux.Router) {
	r.HandleFunc("/home", controllers.GetAllHome).Methods("GET")
	r.HandleFunc("/home", controllers.AddHome).Methods("POST")
	r.HandleFunc("/home/{id}", controllers.GetHomeByID).Methods("GET")
	r.HandleFunc("/home/{id}", controllers.UpdateHome).Methods("PUT")
	r.HandleFunc("/home/{id}", controllers.DeleteHome).Methods("DELETE")
	r.HandleFunc("/upload", controllers.UploadHome).Methods("POST")

}
