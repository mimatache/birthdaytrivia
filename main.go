package main

import (
	"fmt"

	"github.com/mimatache/mariashi/internal/commands"
)

func main() {
	if err := commands.Root().Execute(); err != nil {
		panic(fmt.Errorf("error executing the root command: %w", err))
	}
}
