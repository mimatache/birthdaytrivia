package commands

import (
	"github.com/spf13/cobra"

	"github.com/mimatache/mariashi/internal/commands/about"
)

// Root is the start of the application. All other commands should either be added to this command, or to a command that is in intself added to this command
func Root() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "mariashi",
		Short: "Hello Mariashi, would you like to play a game?",
	}

	rootCommand.AddCommand(
		about.Version,
	)

	return rootCommand
}
