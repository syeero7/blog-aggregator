package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/lib/pq"
	"github.com/syeero7/blog-aggregator/internal/database"
)

func handleAggregation(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("missing required argument 'time between requests'")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s.db); err != nil {
			return err
		}
	}
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	feed := RSSFeed{}
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	return &feed, nil
}

func unescapeFeed(feed *RSSFeed) *RSSFeed {
	f := *feed
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	escapedItems := make([]RSSItem, len(f.Channel.Item))

	for _, item := range f.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		escapedItems = append(escapedItems, item)
	}

	f.Channel.Item = escapedItems
	return &f
}

func scrapeFeeds(db *database.Queries) error {
	nextFeed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feedToMark := database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	if err := db.MarkFeedFetched(context.Background(), feedToMark); err != nil {
		return err
	}

	escFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	feed := unescapeFeed(escFeed)
	for _, post := range feed.Channel.Item {
		if post.Title == "" || post.PubDate == "" {
			continue
		}
		pubDate, err := parsePubDate(post.PubDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if pubDate.IsZero() {
			continue
		}

		postData := database.CreatePostParams{
			Title:       post.Title,
			Description: post.Description,
			Url:         post.Link,
			PublishedAt: pubDate,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			FeedID:      nextFeed.ID,
		}
		if err := db.CreatePost(context.Background(), postData); err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}

			fmt.Printf("failed to create post titled %s\n", post.Title)
			fmt.Println("error: ", err)
			continue
		}
	}

	return nil
}

func parsePubDate(pubDate string) (time.Time, error) {
	dateFormats := []string{
		time.RFC1123,
		time.RFC1123Z,
	}

	for _, format := range dateFormats {
		t, err := time.Parse(format, pubDate)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("failed to parse pubDate %s", pubDate)
}
