package commands

import (
	"os"

	"github.com/cderwin/cryptopals/encoding"
	"github.com/cderwin/cryptopals/utils"
	"github.com/spf13/cobra"
)

var EncodingCommand = &cobra.Command{
	Use:   "encode [action] [flags]",
	Short: "Convert data between different encodings",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputEncoding, err := encoding.ParseEncoding(in)
		if err != nil {
			return err
		}

		outputEncoding, err := encoding.ParseEncoding(out)
		if err != nil {
			return err
		}

		fileReader, err := utils.NewFileReader(args)
		if err != nil {
			return err
		}

		return encoding.EncodeMain(inputEncoding, outputEncoding, fileReader, os.Stdout)
	},
}

// Flags

var (
	in  string
	out string
)

func init() {
	flags := EncodingCommand.Flags()
	flags.StringVar(&in, "in", "", "The encoding of the input data.  Options are bin (binary), hex, or b64 (base64).")
	flags.StringVar(&out, "out", "", "The desired encoding of the output data.  Options are bin (binary), hex, or b64 (base64).")
}

func addEncodingFlags(cmd *cobra.Command) (*string, *string) {
	flags := cmd.PersistentFlags()
	in := flags.String("inFormat", "bin", "Input format.  One of bin, hex, or base64.")
	out := flags.String("outFormat", "bin", "Output format.  One of bin, hex, or base64.")
	return in, out
}
