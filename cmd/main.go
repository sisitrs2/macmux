package main

import (
	"fmt"

	"github.com/sisitrs2/macmux/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf(err)
		return
	}

	fmt.Printf("Welcome to macmux :P")
	cfg.PrintConfig()
}
