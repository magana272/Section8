package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

		t, err := template.ParseFiles(tmpPath+"/home.html", tmpPath+"/navbar.html")
		// t, err := template.ParseFiles("../views/home.html")
		if err != nil {
			fmt.Println("=======ERROR======== ")

			fmt.Println("Could not parse .html")
			fmt.Println(err)

			fmt.Println("==================== ")

		} else {
			// fmt.Println(Home)
			detail := make(map[string]interface{})
			cookie, _ := r.Cookie("session")
			currUser, _ := models.GetUserWithCookie(cookie.Value)
			detail["user"] = &currUser
			detail["homes"] = &Home
			err = t.Execute(w, detail)

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
func UploadHome(w http.ResponseWriter, r *http.Request) {
	// truncated for brevity
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > 1024*1024*100 {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Fprintf(w, "Upload successful")
}
