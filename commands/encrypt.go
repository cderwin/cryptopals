package commands

import (
	"errors"
	"os"

	"github.com/cderwin/cryptopals/ciphers"
	"github.com/cderwin/cryptopals/ciphers/caesar"
	"github.com/cderwin/cryptopals/ciphers/xor"
	"github.com/cderwin/cryptopals/utils"
	"github.com/spf13/cobra"
)

// Errors
var (
	InvalidKeyError = errors.New("Key not given (or empty string)")
	InvalidCipher   = errors.New("Cipher must either be `xor` or `caesar`")
)

// Flags
var (
	key    string
	cipher string
)

var EncryptCommand = &cobra.Command{
	Use:   "encrypt [flags]",
	Short: "Encrypt plaintext under various raw ciphers (i.e. no block modes)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if key == "" {
			return InvalidKeyError
		}

		// Parse cipher argument
		var algo ciphers.Algorithm
		switch cipher {
		case "xor":
			algo = xor.NewXorAlgorithm()
		case "caesar":
			algo = caesar.NewCaesarAlgorithm()
		default:
			return InvalidCipher
		}

		reader, err := utils.NewFileReader(args)
		if err != nil {
			return err
		}

		return ciphers.EncryptMain(algo, []byte(key), reader, os.Stdout)
	},
}

func init() {
	flags := EncryptCommand.Flags()
	flags.StringVarP(&key, "key", "k", "", "Key to encrypt with.  Validity depends on cipher.")
	flags.StringVarP(&cipher, "cipher", "c", "", "Cipher to encrypt with.  Must be `xor` or `caesar`.")
}
