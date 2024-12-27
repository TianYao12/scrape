package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is an empty string")
	}

	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	srv.ListenAndServe()

	fmt.Println("PORT: ", portString)
}
