package main

import (
	"context"
	"fmt"

	"github.com/iahta/blog_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	err := s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		Url:  url,
		Name: user.Name,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}
	fmt.Printf("%s has been unfollowed", url)
	return nil
}
