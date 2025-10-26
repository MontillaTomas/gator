package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/MontillaTomas/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("name argument is required")
	}

	name := cmd.args[0]

	u, err := s.db.GetUserByName(context.Background(), name)
	if err == nil {
		if u.Name == name {
			return fmt.Errorf("user %s already exists", name)
		}
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	s.cfg.SetUser(name)
	fmt.Printf("User %s registered and logged in\n", name)

	return nil
}
