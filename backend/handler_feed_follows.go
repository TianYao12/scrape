package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/TianYao12/scraper/backend/internal/database"
	"github.com/google/uuid"
	"github.com/go-chi/chi"
)

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not parse feed follow id: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not delete feed follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}