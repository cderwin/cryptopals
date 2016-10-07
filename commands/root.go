package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   os.Args[0],
	Short: "A probably (very) buggy crypto toolset",
	Long: `A crypto toolset designed for CTFs and specifically to solve the Matasano challenges.
The toolsetbis likely very buggy, and should not be used for anything that matters.
			
By Cameron Derwin, source at github.com/cderwin/cryptopals`,
}

func AddCommands() {
	RootCommand.AddCommand(XorCommand)
	RootCommand.AddCommand(EncodingCommand)
}

func Execute() {
	AddCommands()

	if err := RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
