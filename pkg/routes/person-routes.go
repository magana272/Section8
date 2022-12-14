package routes

import (
	"github.com/gorilla/mux"
	"github.com/magana272/Section8/pkg/controllers"
)

var RegisterOwnerRouter = func(r *mux.Router) {
	r.HandleFunc("/People", controllers.GetAllPeople).Methods("GET")
	r.HandleFunc("/AddPerson", controllers.AddPerson).Methods("POST")
	r.HandleFunc("/Person/{id}", controllers.DeletePersonWithId).Methods("DELETE")
	r.HandleFunc("/Person/{id}", controllers.UpdatePerson).Methods("PUT")
}
