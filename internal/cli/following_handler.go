package cli

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	currentUserName := s.cfg.CurrentUserName
	if currentUserName == "" {
		return fmt.Errorf("no user is currently logged in")
	}

	user, err := s.db.GetUserByName(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get followed feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Println("Followed feeds:")
	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.FeedName)
	}

	return nil
}
