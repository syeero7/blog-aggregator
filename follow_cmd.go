package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/syeero7/blog-aggregator/internal/database"
)

func handleFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return errors.New("missing required argument 'feed url'")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	feedFollowData := database.CreateFeedFollowParams{
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	followFeed, err := s.db.CreateFeedFollow(context.Background(), feedFollowData)
	if err != nil {
		return err
	}

	fmt.Printf("user '%s' is now following the '%s' feed", followFeed.Username, followFeed.FeedName)
	return nil
}

func handleGetFollowing(s *state, _ command) error {
	followingFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	fmt.Printf("feeds followed by %s", s.cfg.CurrentUserName)
	for _, feed := range followingFeeds {
		fmt.Printf("* feed: %s | creator: %s\n", feed.FeedName, feed.Creator)
	}

	return nil
}
