package commands

import (
	"github.com/spf13/cobra"

	"github.com/mimatache/birthdaytrivia/internal/commands/about"
	"github.com/mimatache/birthdaytrivia/internal/commands/run"
)

// Root is the start of the application. All other commands should either be added to this command, or to a command that is in intself added to this command
func Root() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "birthdaytrivia",
		Short: "Hello Mariashi, would you like to play a game?",
	}

	rootCommand.AddCommand(
		about.Command(),
		run.Command(),
	)

	return rootCommand
}
