package babe

import (
	"net/http"
	"encoding/json"
	"babelist.ph/api/config/database"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/ioutil"
)

type Babe struct {
	gorm.Model
	ID uint
	Name string
	Gender string
	Age string
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	var babes []Babe
	db.Find(&babes)
	json.NewEncoder(w).Encode(babes)
}

func Find(w http.ResponseWriter, r *http.Request) {
	// extract the id 
	vars := mux.Vars(r)
    id := vars["id"]

	db := database.Connect()
	
	var babe Babe
	db.Find(&babe, id)
	json.NewEncoder(w).Encode(babe)

}

func New(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()

	// get the body of our POST request
    reqBody, err := ioutil.ReadAll(r.Body)
	if (err != nil) {
		panic("error parsing request body")
	}
	
	var babe Babe
	//convert request body to json
    json.Unmarshal(reqBody, &babe)
	db.Create(&babe)

	json.NewEncoder(w).Encode(babe)

}

func Update(w http.ResponseWriter, r *http.Request) {

	// extract the id from the route
	vars := mux.Vars(r)
    id := vars["id"]

	db := database.Connect()

	// find by id
	var babe Babe
	db.Find(&babe, id)

	var updatedDetails Babe

	// get the body of our PUT request
    reqBody, err := ioutil.ReadAll(r.Body)
	if (err != nil) {
		panic("error parsing request body")
	}

	// convert the json body to struct
	err = json.Unmarshal(reqBody, &updatedDetails)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(&updatedDetails)

	db.Model(&babe).Updates(&updatedDetails)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	// extract the id 
	vars := mux.Vars(r)
    id := vars["id"]

	db := database.Connect()
	var babe Babe
	db.Delete(&babe, id)

}


