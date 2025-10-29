package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func addfeedHandler(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("name and url arguments are required")
	}

	currentUserName := s.cfg.CurrentUserName
	if currentUserName == "" {
		return fmt.Errorf("no user is currently logged in")
	}
	currentUserId, err := s.db.GetUserByName(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUserId.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to add feed: %w", err)
	}

	paramsFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    currentUserId.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.db.CreateFeedFollow(context.Background(), paramsFollow)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	fmt.Printf("Feed name: %s\n", feed.Name)
	fmt.Printf("Feed URL: %s\n", feed.Url)
	fmt.Printf("Feed User ID: %s\n", feed.UserID)
	fmt.Printf("Feed Created At: %s\n", feed.CreatedAt)
	fmt.Printf("Feed Updated At: %s\n", feed.UpdatedAt)

	return nil
}
