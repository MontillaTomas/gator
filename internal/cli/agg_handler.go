package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
	"github.com/MontillaTomas/blog-aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	var timeString string
	if len(cmd.args) < 1 {
		return fmt.Errorf("time between requests argument is required")
	}
	timeString = cmd.args[0]
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

	//printFeed(feed)

	// Saving posts from feed

	for _, item := range feed.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return fmt.Errorf("failed to parse published date: %w", err)
		}
		params := database.CreatePostParams{
			FeedID:      nextFeedToFetch.ID,
			Title:       item.Title,
			Description: newNullString(item.Description),
			Url:         item.Link,
			PublishedAt: publishedAt,
		}
		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				continue
			}
			return fmt.Errorf("failed to create post: %w", err)
		}
	}

	return nil
}

func newNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
