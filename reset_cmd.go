package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, _ command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("user table has been successfully reset")
	return nil
}
