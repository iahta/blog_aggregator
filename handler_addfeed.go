package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iahta/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	//Add feed expects only 2 argument
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> %s <url>", cmd.Name, cmd.Args)
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

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(feed)
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:		%v\n", feed.ID)
	fmt.Printf(" * Name:	%v\n", feed.Name)
	fmt.Printf(" * URL:		%v\n", feed.Url)
	fmt.Printf(" * UserID:	%v\n", feed.UserID)
}

func handlerFeed(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't find any feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("%s\n", feed.Username)

	}
	return nil

}
