package main

import (
	"log"
	"context"
	"net/http"
	"fmt"
	"math/rand"
	"time"
	"encoding/json"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
)

type Response struct {
	Status string `json:"status"`
	Data spotify.FullTrack `json:"data"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
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

func ReturnRandomSong(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint reached: ReturnRandomSong")
	token := GetToken()

	ctx := context.Background()
	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	query := fmt.Sprintf("%v", RandomLetter())

	results, err := client.Search(ctx, query, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())


	resp := Response{Status: "200", Data: results.Tracks.Tracks[rand.Intn(len(results.Tracks.Tracks)+1)+1]}
	json.NewEncoder(w).Encode(resp)

}
