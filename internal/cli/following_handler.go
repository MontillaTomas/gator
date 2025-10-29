package cli

import (
	"context"
	"fmt"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
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
