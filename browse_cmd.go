package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/syeero7/blog-aggregator/internal/database"
)

func handlePostBrowsing(s *state, cmd command) error {
	postLimit := 2
	if len(cmd.arguments) > 0 {
		if limit, err := strconv.Atoi(cmd.arguments[0]); err == nil {
			postLimit = limit
		}
	}

	userPosts := database.GetPostsForUserParams{
		Name:  s.cfg.CurrentUserName,
		Limit: int32(postLimit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), userPosts)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("title: %s\n", post.Title)
		fmt.Printf("description: %s\n", post.Description)
		fmt.Printf("published at: %s\n", post.PublishedAt)
		fmt.Printf("url: %s\n", post.Url)
		fmt.Println()
	}

	return nil
}
