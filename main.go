package main

import (
	"fmt"
	"os"

	"github.com/syeero7/blog-aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := state{config: &cfg}
	c := commands{commands: make(map[string]func(*state, command) error)}
	if len(os.Args) < 2 {
		fmt.Println("insufficient arguments provided")
		os.Exit(1)
	}

	c.register("login", handleLogin)

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	if err := c.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
