package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func handleAggregation(_ *state, _ command) error {
	tmpURL := "https://www.wagslane.dev/index.xml"
	escFeed, err := fetchFeed(context.Background(), tmpURL)
	if err != nil {
		return err
	}

	feed := unescapeFeed(escFeed)
	fmt.Printf("%+v\n", *feed)
	return nil
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
