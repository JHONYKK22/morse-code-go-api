package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/JHONYKK22/morse-code-api/handlers/encoder"
	"github.com/gorilla/mux"
	env "github.com/joho/godotenv"
)

func main() {

	err := env.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	router := mux.NewRouter()

	router.Use(commonMiddleware)
	router.Use(handlerLoggerMiddleware)

	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)

	encodeSub := router.PathPrefix("/encode").Subrouter()
	encodeSub.HandleFunc("/{text}", encoder.EncodePath).Methods(http.MethodGet)
	encodeSub.HandleFunc("", encoder.EncodeBody).Methods(http.MethodPost)

	decodeSub := router.PathPrefix("/decode").Subrouter()
	decodeSub.HandleFunc("/{text}", encoder.DecodePath).Methods(http.MethodGet)
	decodeSub.HandleFunc("", encoder.DecodeBody).Methods(http.MethodPost)

	fmt.Println("Running...")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handlerLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("::", r.Method, "->", r.URL)

		next.ServeHTTP(w, r)

	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Welcome to Morse code Encoder/Decoder"))
	res := map[string]string {"msg": "Welcome to Morse code Encoder/Decoder"}
	json.NewEncoder(w).Encode(res)
}