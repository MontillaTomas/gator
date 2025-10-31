package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
)

func browseHandler(s *state, cmd command, user database.User) error {
	var limitString string
	if len(cmd.args) < 1 {
		limitString = "2"
	} else {
		limitString = cmd.args[0]
	}
	limitInt, err := strconv.Atoi(limitString)
	if err != nil || limitInt <= 0 {
		return fmt.Errorf("invalid limit")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get followed feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limitInt),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to get posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts available from followed feeds.")
		return nil
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Published At: %s\n", post.PublishedAt)
		fmt.Println("--------------------------------------------------")
	}

	return nil
}
