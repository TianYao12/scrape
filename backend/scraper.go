package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/TianYao12/scraper/backend/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:")
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			// go scrapeFeed(db, wg, feed)
			go scrapeFireballFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("could not parse date %v with err %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: publishedAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key") {
					continue
				}
				log.Println("failed to create post: ", err)
			}
	}
	log.Printf("Feed %s collected, %v posts foiunds", feed.Name, len(rssFeed.Channel.Item))
}
func scrapeFireballFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
    defer wg.Done()

    log.Printf("Fetching fireball feed: %v", feed.Url)

    fireballFeed, err := urlToFireballFeed(feed.Url)
    if err != nil {
        log.Println("Error fetching fireball feed:", err)
        return
    }

    log.Printf("Parsed rows: %d", len(fireballFeed))

    for _, row := range fireballFeed {
        if row.TotalRadiatedEnergyJ == 0 {
            continue
        }

        _, err := db.CreateFireball(context.Background(), database.CreateFireballParams{
            ID:                uuid.New(),
            CreatedAt:         time.Now().UTC(),
            UpdatedAt:         time.Now().UTC(),
            TotalRadiatedEnergy: row.TotalRadiatedEnergyJ,
            FeedID:            feed.ID,
        })

        if err != nil {
            if strings.Contains(err.Error(), "duplicate key") {
                continue
            }
            log.Println("Failed to create fireball record: ", err)
        }
    }

    log.Printf("Fireball feed %s collected, %d rows found", feed.Name, len(fireballFeed))
}
