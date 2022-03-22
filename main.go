package main

import (
	"fmt"

	"github.com/mimatache/birthdaytrivia/internal/commands"
)

func main() {
	if err := commands.Root().Execute(); err != nil {
		panic(fmt.Errorf("error executing the root command: %w", err))
	}
}
