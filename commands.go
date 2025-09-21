package main

import (
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("invalid command: '%s'", cmd.name)
	}

	if err := cmdFunc(s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
