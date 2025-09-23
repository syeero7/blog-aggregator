package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/syeero7/blog-aggregator/internal/database"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return errors.New("missing required arguments 'feed name' and 'feed url'")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedData := database.CreateFeedParams{
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feed, err := s.db.CreateFeed(context.Background(), feedData)
	if err != nil {
		return err
	}

	feedFollowData := database.CreateFeedFollowParams{
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if _, err := s.db.CreateFeedFollow(context.Background(), feedFollowData); err != nil {
		return err
	}

	fmt.Printf("feed added successfully: %+v\n", feed)
	return nil
}
