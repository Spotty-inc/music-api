package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Song struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Artist string `json:"artist`
	Popularity int `json:"popularity"`
}

var Songs []Song

func healthCheck(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: healthCheck")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func returnSingleSong(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: returnSingleSong")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, song := range Songs {
		if song.Id == key {
			json.NewEncoder(w).Encode(song)
			return
		}
	}
	log.Printf("Error: ID %s does not exist", key)
}

func returnAllSongs(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: returnAllSongs")
	json.NewEncoder(w).Encode(Songs)
}

func requestHandler(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/songs", returnAllSongs)
	router.HandleFunc("/song/{id}", returnSingleSong)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main(){
	Songs = []Song{
		Song{Id: "1", Title: "Titanium", Artist: "Dave", Popularity: 60},
		Song{Id: "2", Title: "Low", Artist: "Sza", Popularity: 75},
		Song(Id: "3", Title: "Just Wanna Rock", Artist: "Lil Uzi Vert", Popularity: 80),
	}
	requestHandler()
}
