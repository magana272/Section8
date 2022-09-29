package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/magana272/Section8/pkg/models"
	"github.com/magana272/Section8/pkg/views"
	uuid "github.com/satori/go.uuid"
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

func Index(w http.ResponseWriter, r *http.Request) {
	tmpPath, err := views.GetTemplatePath()
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	if err != nil {
		fmt.Println("=========================================")
		fmt.Println("==============Error while getting tmep=============")

		fmt.Println("=========================================")
		fmt.Println("=========================================")

		fmt.Println(err)

	}
	t, err := template.ParseFiles(tmpPath+"/index.html", tmpPath+"/navbar.html")
	if err != nil {
		fmt.Println("error in parsing")
	}
	user, ok := models.GetUserWithCookie(cookie.Value)
	if ok != nil {
		t.Execute(w, nil)
	}
	detail := make(map[string]interface{})
	detail["user"] = &user
	t.Execute(w, detail)

}

func Assistents(w http.ResponseWriter, r *http.Request) {
	tmpPath, err := views.GetTemplatePath()
	if err != nil {
		fmt.Println("==============Error while getting temp=============")

		fmt.Println(err)

	}
	t, err := template.ParseFiles(tmpPath+"/assist.html", tmpPath+"/navbar.html", tmpPath+"/head.html")
	if err != nil {
		fmt.Println("==============Error while Parsing temp=============")

	}
	cookie, err := r.Cookie("session")
	detail := make(map[string]interface{})
	user, _ := models.GetUserWithCookie(cookie.Value)
	detail["user"] = &user
	fmt.Println(detail)
	t.Execute(w, detail)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	tmpPath, err := views.GetTemplatePath()
	if err != nil {
		fmt.Println("==============Error while getting tmep=============")
		fmt.Println(err)

	}
	println(tmpPath)
	t, err := template.ParseFiles(tmpPath+"/index.html", tmpPath+"/navbar.html")
	if err != nil {
		fmt.Println("==============couldn't parse /home.html=============")

	}
	t.Execute(w, nil)

}
func Siginup(w http.ResponseWriter, r *http.Request) {
	templatepath, err := views.GetTemplatePath()
	detail := make(map[string]interface{})
	if err != nil {
		fmt.Println("error finding template")

	}
	if r.Method == "POST" {
		// Check if session exist
		cookie, err := r.Cookie("session")
		if err != nil {
			id := uuid.NewV4()
			cookie = &http.Cookie{
				Name:     "session",
				Value:    id.String(),
				HttpOnly: true,
				Path:     "/",
			}
			http.SetCookie(w, cookie)
		}
		// Create User
		r.ParseForm()
		fN := r.FormValue("firstName")
		lN := r.FormValue("lastName")
		eM := r.FormValue("email")
		password := r.FormValue("password")

		Check, err := models.GetUserWithEmail(eM)
		if err != nil {
			fmt.Println(Check.Name)
			fmt.Println("exist")
			detail["Exist"] = true
			t, _ := template.ParseFiles(templatepath+"/signup.html", templatepath+"/navbar.html")
			err = t.Execute(w, detail)

		}
		newUser := &models.User{Name: fN + " " + lN,
			Email:    eM,
			Password: password,
		}
		newUser.CreateUser()
		// Create the session for the user
		newSession := models.Session{Cookie: cookie.Value}
		newSession.CreateSession(newUser)
		// http redirecting
		http.Redirect(w, r, "/", 301)

	} else {
		cookie, _ := r.Cookie("session")
		currUser, err := models.GetUserWithCookie(cookie.Value)
		t, err := template.ParseFiles(templatepath+"/signup.html", templatepath+"/navbar.html")
		if err != nil {
			fmt.Printf(" error while parsing%s", templatepath)
		}
		detail["user"] = &currUser
		err = t.Execute(w, detail)

	}
}
func Login(w http.ResponseWriter, r *http.Request) {

	templatepath, err := views.GetTemplatePath()
	cookie, _ := r.Cookie("session")
	currUserCookie, err := models.GetUserWithCookie(cookie.Value)
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")
		CheckUser, _ := models.GetUserWithEmail(email)
		if CheckUser.Password != password {
			fmt.Println(password)
			http.Redirect(w, r, "/IncorrectLogin", 301)
		} else {
			cookie, err := r.Cookie("session")
			if err != nil {
				id := uuid.NewV4()
				cookie = &http.Cookie{
					Name:     "session",
					Value:    id.String(),
					HttpOnly: true,
					Path:     "/"}
			}
			http.SetCookie(w, cookie)
			newSession := models.Session{Cookie: cookie.Value}
			newSession.CreateSession(CheckUser)
			http.Redirect(w, r, "/", 301)
		}
	}
	t, err := template.ParseFiles(templatepath+"/login.html", templatepath+"/navbar.html")
	if err != nil {
		fmt.Printf(" error while parsing%s", templatepath)
	}
	detail := make(map[string]interface{})
	detail["user"] = &currUserCookie
	err = t.Execute(w, detail)

}
func LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	id := uuid.NewV4()
	cookieClear := &http.Cookie{
		Name:     "session",
		Value:    id.String(),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookieClear)
	models.DeleteSession(cookie.Value)
	http.Redirect(w, r, "/", 301)
}

func IncorrectLogin(w http.ResponseWriter, r *http.Request) {
	templatepath, err := views.GetTemplatePath()
	t, err := template.ParseFiles(templatepath+"/IC.html", templatepath+"/navbar.html")
	if err != nil {
		fmt.Printf(" error while parsing%s", templatepath)
	}
	id := uuid.NewV4()
	cookieClear := &http.Cookie{
		Name:     "session",
		Value:    id.String(),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookieClear)
	cookie, _ := r.Cookie("session")
	currUser, err := models.GetUserWithCookie(cookie.Value)
	detail := make(map[string]interface{})
	detail["user"] = &currUser
	err = t.Execute(w, detail)

}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	//  get the profile
	tempPath, _ := views.GetTemplatePath()
	t, _ := template.ParseFiles(tempPath+"/Profile.html", tempPath+"/navbar.html")
	cookie, _ := r.Cookie("session")
	currUser, _ := models.GetUserWithCookie(cookie.Value)
	detail := make(map[string]interface{})
	detail["user"] = &currUser
	_ = t.Execute(w, detail)
}
