package main

import (
	"context"
	"fmt"
)

func handleGetFeeds(s *state, _ command) error {
	feeds, err := s.db.GetALLFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Printf("url: %s\n", feed.Url)
		fmt.Printf("user: %s\n", feed.Username)
		fmt.Println()
	}

	return nil
}
