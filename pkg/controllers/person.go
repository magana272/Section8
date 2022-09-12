package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/magana272/Section8/pkg/models"
	"github.com/magana272/Section8/pkg/utils"
)

var Person models.Person
var db *gorm.DB

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	people = models.GetAllPeople()
	res, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Issue while marshaling data")
	}
	w.Header().Set("Content-Type", "pgklication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func AddPerson(w http.ResponseWriter, r *http.Request) {
	var newperson = &models.Person{}
	utils.ParseBody(r, newperson)

	np := newperson.CreatePerson()
	for _, h := range np.Homes {
		h.CreateHome(np)
	}
	res, err := json.Marshal(&np)
	if err != nil {
		fmt.Println("Marshal error while adding person")
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
