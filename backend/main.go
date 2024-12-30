package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TianYao12/scraper/backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	feed, err := urlToFeed("https://wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(feed)

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is an empty string")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not in the env file broski")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	db := database.New(connection)
	apiConfig := apiConfig{DB: database.New(connection)}

	go startScraping(db, 10, time.Minute)
	
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/test", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Get("/users", apiConfig.middlewareAuth(apiConfig.handlerGetUser))
	
	v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerCreateFeed))
	v1Router.Get("/feeds", apiConfig.handlerGetFeeds)

	v1Router.Get("/posts", apiConfig.middlewareAuth(apiConfig.handlerGetPostsForUser))

	v1Router.Post("/feed_follows", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT: ", portString)
}
