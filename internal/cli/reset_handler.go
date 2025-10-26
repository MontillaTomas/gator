package cli

import "context"

func resetHandler(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}
