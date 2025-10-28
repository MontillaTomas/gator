package cli

import (
	"context"
	"fmt"

	"github.com/MontillaTomas/blog-aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	var feedURL string
	if len(cmd.args) > 0 {
		feedURL = cmd.args[0]
	} else {
		feedURL = "https://www.wagslane.dev/index.xml"
	}

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("%s\n\n", feed.Channel.Title)
	fmt.Printf("%s\n\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("%s\n\n", item.Title)
		fmt.Printf("%s\n\n", item.Description)
	}

	return nil
}
