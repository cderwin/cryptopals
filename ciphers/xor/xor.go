package xor

import (
	"bufio"
	"io"
)

// Xor Encrypt Hook

func XorMain(r io.Reader, key []byte, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	keyLen := len(key)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		ciphertext := make([]byte, len(bytes))
		for i, b := range bytes {
			ciphertext[i] = b ^ key[i%keyLen]
		}
		w.Write(append(ciphertext, byte('\n')))
	}

	return nil
}

func xor(plaintext, key []byte) []byte {
	keyLength := len(key)
	ciphertext := make([]byte, len(plaintext))
	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ key[i%keyLength]
	}
	return ciphertext
}
