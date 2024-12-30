package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title		string		`xml:"title"`
		Link 		string 		`xml:"link"`
		Description string 		`xml:"description"`
		Language 	string 		`xml:"language"`
		Item 		[]RSSItem 	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title		string `xml:"title"`
	Link		string	`xml:"link"`
	Description	string	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
}
type FireballFeed struct {
    OuterRows []OuterRow `xml:"row"`
}

type OuterRow struct {
    Rows []FireballRow `xml:"row"`
}

type FireballRow struct {
    TotalRadiatedEnergyJ float64 `xml:"total_radiated_energy_j"`
}



func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssFeed, nil
}

func urlToFireballFeed(url string) ([]FireballRow, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch URL: %w", err)
    }
    defer resp.Body.Close()

    var feed FireballFeed
    if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
        return nil, fmt.Errorf("failed to decode XML: %w", err)
    }

    var rows []FireballRow
    for _, outerRow := range feed.OuterRows {
        rows = append(rows, outerRow.Rows...)
    }

    return rows, nil
}
