package cli

import (
	"context"
	"fmt"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		currentUserName := s.cfg.CurrentUserName
		if currentUserName == "" {
			return fmt.Errorf("no user is currently logged in")
		}
		currUser, err := s.db.GetUserByName(context.Background(), currentUserName)
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}
		user := database.User{
			ID:        currUser.ID,
			Name:      currUser.Name,
			CreatedAt: currUser.CreatedAt,
			UpdatedAt: currUser.UpdatedAt,
		}

		return handler(s, c, user)
	}
}
