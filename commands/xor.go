package commands

import (
	"errors"
	"os"

	"github.com/cderwin/cryptopals/basic"
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
		return basic.XorHook(os.Stdin, []byte(key), os.Stdout)
	},
}

func init() {
	flags := XorCommand.Flags()
	flags.StringVarP(&key, "key", "k", "", "Xor encryotion key")
}
