package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("feed URL is required")
	}
	feedURL := cmd.args[0]

	currentUserName := s.cfg.CurrentUserName
	if currentUserName == "" {
		return fmt.Errorf("no user is currently logged in")
	}
	user, err := s.db.GetUserByName(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	fmt.Printf("%s is now following feed %s\n", currentUserName, feed.Name)

	return nil
}
