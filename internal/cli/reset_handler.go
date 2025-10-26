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
	fmt.Printf("Users table has been reset\n")
	return nil
}
