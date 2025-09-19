package main

import (
	"fmt"

	"github.com/syeero7/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", cfg)

	if err := cfg.SetUser("Emal"); err != nil {
		fmt.Println(err)
		return
	}

	cfg2, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", cfg2)
}
