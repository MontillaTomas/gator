package cli

import (
	"context"
	"fmt"
)

func feedsHandler(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("* %s (%s) added by %s\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil
}
