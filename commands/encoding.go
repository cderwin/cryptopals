package commands

import (
	"github.com/cderwin/cryptopals/basic"
	"github.com/spf13/cobra"
)

var EncodingCommand = &cobra.Command{
	Use:   "encode [action] [flags]",
	Short: "Convert data between different encodings",
	RunE: func(cmd *cobra.Command, args []string) error {
		return basic.EncodeMain(args)
	},
}
