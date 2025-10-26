package cli

import (
	"context"
	"fmt"
)

func usersHandler(s *state, cmd command) error {
	currentUserName := s.cfg.CurrentUserName
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}

	for _, user := range users {
		if currentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
