package cli

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username argument is required")
	}

	username := cmd.args[0]

	_, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("user %s does not exist", username)
		}
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("Logged in as %s\n", username)
	return nil
}
