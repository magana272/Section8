package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/magana272/Section8/pkg/views"
)

var mSK = []byte("key")

func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mSK, nil
			})
			if err != nil {
				fmt.Println("There was an error")
			}
			if token.Valid {
				fmt.Println("reached endpoint")
				endpoint(w, r)
			} else {
				fmt.Println("not authorized")

			}
		}
		fmt.Println("=========================================")
		fmt.Println("=========User was not Authorized=========")
		fmt.Println("=========Redirct to login page  =========")
		fmt.Println("=========================================")
		Login(w, r)
	})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	tmpPath, err := views.GetTemplatePath()
	if err != nil {
		fmt.Println("=========================================")
		fmt.Println("==============Error while getting tmep=============")

		fmt.Println("=========================================")
		fmt.Println("=========================================")

		fmt.Println(err)

	}
	println(tmpPath)
	t, err := template.ParseFiles(tmpPath + "/home.html")
	if err != nil {
		fmt.Println("=========================================")
		fmt.Println("==============couldn't parse /home.html=============")

		fmt.Println("=========================================")
		fmt.Println("=========================================")

	}
	t.Execute(w, nil)

}
func Siginup(w http.ResponseWriter, r *http.Request) {
	panic("err")

}
func Login(w http.ResponseWriter, r *http.Request) {
	templatepath, err := views.GetTemplatePath()
	if err != nil {
		fmt.Println("=========================================")
		fmt.Println("=========================================")
		fmt.Println("error finding template")
		fmt.Println("=========================================")
		fmt.Println("=========================================")
	}
	t, err := template.ParseFiles(templatepath + "/login.html")
	if err != nil {
		fmt.Println("=========================================")
		fmt.Println("=========================================")
		fmt.Printf(" error while parsing%s", templatepath)
		fmt.Println("=========================================")
		fmt.Println("=========================================")
	}
	err = t.Execute(w, nil)
}
