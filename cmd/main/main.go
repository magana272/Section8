package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magana272/Section8/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLocationRouter(r)
	routes.RegisterOwnerRouter(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./my-app/build/")))
	fmt.Println("Now Lsitening and Serveing at LocalHost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
