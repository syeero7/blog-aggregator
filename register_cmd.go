package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/syeero7/blog-aggregator/internal/database"
)

func handleRegistration(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("missing required argument 'username'")
	}

	userData := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.arguments[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := s.db.CreateUser(context.Background(), userData)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("new user registration successful: %+v\n", user)
	return nil
}
