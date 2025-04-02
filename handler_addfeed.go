package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iahta/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	//Add feed expects only 2 argument
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> %s <url>", cmd.Name, cmd.Args)
	}
	userName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	feedName := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:		%v\n", feed.ID)
	fmt.Printf(" * Name:	%v\n", feed.Name)
	fmt.Printf(" * URL:		%v\n", feed.Url)
	fmt.Printf(" * UserID:	%v\n", feed.UserID)
}
