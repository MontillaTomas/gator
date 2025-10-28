package cli

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	err = s.cfg.SetUser("")
	if err != nil {
		return fmt.Errorf("failed to reset user: %w", err)
	}
	fmt.Printf("Users table has been reset\n")
	return nil
}
