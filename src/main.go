package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%v: %v remote-addr: %v host: %v", r.Method,r.RequestURI, r.RemoteAddr, r.Host)
        next.ServeHTTP(w, r)
    })
}

func requestHandler(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/random", ReturnRandomSong)
	router.HandleFunc("/health", HealthCheck)
	router.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main(){
	err := godotenv.Load()
  	if err != nil {
    log.Println("Couldn't find .env file. Using local environment variables instead.")
  	}

	requestHandler()
}
