package main

import (
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("missing required argument 'username'")
	}

	username := cmd.arguments[0]
	if err := s.config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("setting the user '%s' was successful\n", username)
	return nil
}
