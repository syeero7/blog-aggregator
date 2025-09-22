package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/syeero7/blog-aggregator/internal/config"
	"github.com/syeero7/blog-aggregator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	s := state{cfg: &cfg, db: dbQueries}
	c := commands{commands: make(map[string]func(*state, command) error)}
	if len(os.Args) < 2 {
		fmt.Println("insufficient arguments provided")
		os.Exit(1)
	}

	c.register("login", handleLogin)
	c.register("register", handleRegistration)
	c.register("reset", handleReset)
	c.register("users", handleListUsers)

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	if err := c.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
