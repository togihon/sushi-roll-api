package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model sushi roll
type SushiRoll struct {
	ID         string `json: "id"`
	Name       string `json: "name"`
	Decription string `json: "description"`
	Ingredient string `json: "ingredient"`
}

// semua data disimpan didalam slice
var sushirolls []SushiRoll

// get all sushi
func getAllSushiRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set header response json
	json.NewEncoder(w).Encode(sushirolls)              //untuk merender slice sushirolls menjadi json
}

// get single sushi
func getSushiRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get parameter menggunakan mux
	for _, item := range sushirolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) //mengirim response
			return
		}
	}
}

// create sushi
func createSushiRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newSushiRoll SushiRoll //instance struct

	json.NewDecoder(r.Body).Decode(&newSushiRoll)       //membaca request body
	newSushiRoll.ID = strconv.Itoa(len(sushirolls) + 1) //id secara otomatis +1
	sushirolls = append(sushirolls, newSushiRoll)       //menambahkan instance ke slice
	json.NewEncoder(w).Encode(newSushiRoll)             //mengirim response
}

// update sushi
func updateSushiRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get parameter menggunakan mux

	for i, item := range sushirolls {
		if item.ID == params["id"] {
			//untuk ngasi tahu objek yang di appen | objek yang akan diappend
			sushirolls = append(sushirolls[:i], sushirolls[i+1:]...) //hapus data
			var newSushiRoll SushiRoll                               //instance struct
			//add new data yang diedit
			json.NewDecoder(r.Body).Decode(&newSushiRoll)
			newSushiRoll.ID = params["id"]
			sushirolls = append(sushirolls, newSushiRoll)
			json.NewEncoder(w).Encode(newSushiRoll)

			return
		}
	}
}

// delete sushi
func deleteSushiRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get parameter menggunakan mux

	for i, item := range sushirolls {
		if item.ID == params["id"] {
			sushirolls = append(sushirolls[:i], sushirolls[i+1:]...)
			break
		}
	}
}

func main() {

	//init router
	router := mux.NewRouter()

	//handle endpoints/routing
	router.HandleFunc("/sushi", getAllSushiRoll).Methods("GET")
	router.HandleFunc("/sushi/{id}", getSushiRoll).Methods("GET")
	router.HandleFunc("/sushi", createSushiRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateSushiRoll).Methods("PUT")
	router.HandleFunc("/sushi/{id}", deleteSushiRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
