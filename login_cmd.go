package main

import (
	"context"
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("missing required argument 'username'")
	}

	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("setting the user '%s' was successful\n", user.Name)
	return nil
}
