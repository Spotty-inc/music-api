package main

import (
	"log"
	"context"
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2"
)

type Song struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Artist string `json:"artist`
	Popularity int `json:"popularity"`
}

var Songs []Song

func getToken() *oauth2.Token {
	SPOTIFY_ID := os.Getenv("SPOTIFY_ID")
	SPOTIFY_SECRET := os.Getenv("SPOTIFY_SECRET")
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID: SPOTIFY_ID,
		ClientSecret: SPOTIFY_SECRET,
		TokenURL: spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	return token
}

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

func returnRandomSong(w http.ResponseWriter, r *http.Request){
	token := getToken()
	ctx := context.Background()
	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	results, err := client.Search(ctx, "One Dance", spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}
	if results.Tracks != nil {
		fmt.Println("Songs:")
		for _, item := range results.Tracks.Tracks {
			fmt.Println("   ", item.Name)
		}
	}

}

func returnSingleSong(w http.ResponseWriter, r *http.Request){
	// token := getToken()
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
	router.HandleFunc("/random", returnRandomSong)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main(){
	Songs = []Song{
		Song{Id: "1", Title: "Titanium", Artist: "Dave", Popularity: 60},
		Song{Id: "2", Title: "Low", Artist: "Sza", Popularity: 75},
	}
	err := godotenv.Load()
  	if err != nil {
    log.Fatal("Error loading .env file")
  	}
	requestHandler()
}
