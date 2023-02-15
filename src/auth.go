package main

import (
	"log"
	"context"
	"os"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2"
)

func GetToken() *oauth2.Token {
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
