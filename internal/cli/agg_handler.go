package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/MontillaTomas/blog-aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	var timeString string
	if len(cmd.args) > 0 {
		timeString = cmd.args[0]
	}
	timeBetweenRequests, err := time.ParseDuration(timeString)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("failed to scrape feeds: %w", err)
		}
	}

	return nil
}

func printFeed(feed *rss.RSSFeed) {
	fmt.Printf("%s\n\n", feed.Channel.Title)
	fmt.Printf("%s\n\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("%s\n\n", item.Title)
		fmt.Printf("%s\n\n", item.Description)
	}
}

func scrapeFeeds(s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeedToFetch.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}

	feed, err := rss.FetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	printFeed(feed)

	return nil
}
