package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magana272/Section8/pkg/models"
	"github.com/magana272/Section8/pkg/utils"
	"github.com/magana272/Section8/pkg/views"
)

var Home models.Home

func GetAllHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Home := models.GetAllHome()
		// res, _ := json.Marshal(Home)
		// w.Header().Set("Content-Type", "pkglication")
		// w.WriteHeader(http.StatusOK)
		// w.Write(res)
		// Parsing the required html
		// file in same directory
		// Template
		tmpPath, err := views.GetTemplatePath()
		if err != nil {
			fmt.Println(err)

		}

		t, err := template.ParseFiles(tmpPath + "/home.html")
		// t, err := template.ParseFiles("../views/home.html")
		if err != nil {
			fmt.Println("=======ERROR======== ")

			fmt.Println("Could not parse .html")
			fmt.Println(err)

			fmt.Println("==================== ")

		} else {
			homeToreturn := Home[0]
			err = t.Execute(w, homeToreturn)
		}

	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)

	}

}
func GetHomeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Params := mux.Vars(r)
		homeid := Params["homeid"]
		id, err := strconv.ParseUint(homeid, 0, 0)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while Parsing")
		}
		Home, _ := models.GetHomeById(uint(id))
		res, _ := json.Marshal(Home)
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
	}

}
func AddHome(w http.ResponseWriter, r *http.Request) {
	var newhome = &models.Home{}
	v := mux.Vars(r)
	name := v["name"]
	per := models.GetPersonByName(name)
	utils.ParseBody(r, newhome)
	nh := newhome.CreateHome(per)
	res, _ := json.Marshal(nh)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteHome(w http.ResponseWriter, r *http.Request) {
	panic("not Implemnted")
}
func UpdateHome(w http.ResponseWriter, r *http.Request) {
	panic("not Implemnted")
}
