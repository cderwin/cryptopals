package commands

import (
	"errors"
	"os"

	"github.com/cderwin/cryptopals/ciphers/xor"
	"github.com/cderwin/cryptopals/utils"
	"github.com/spf13/cobra"
)

// Errors
var (
	InvalidKeyError = errors.New("Key must be given and have length greater than zero")
)

// Flags
var (
	key string
)

var XorCommand = &cobra.Command{
	Use:   "xor [action] [flags]",
	Short: "Encrypt, decrypt, and detect vanilla xor ciphers",
	RunE: func(cmd *cobra.Command, args []string) error {
		if key == "" {
			return InvalidKeyError
		}

		reader, err := utils.NewFileReader(args)
		if err != nil {
			return err
		}

		return xor.XorMain(reader, []byte(key), os.Stdout)
	},
}

func init() {
	flags := XorCommand.Flags()
	flags.StringVarP(&key, "key", "k", "", "Xor encryotion key")
}
