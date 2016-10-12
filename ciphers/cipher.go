package ciphers

import (
	"bufio"
	"errors"
	"io"
)

var NoKeySetError = errors.New("Key must be set before encryption or decryption")

func EncryptMain(algo Algorithm, key []byte, r io.Reader, w io.Writer) error {
	if err := algo.ParseKey(key); err != nil {
		return err
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		ciphertext, err := algo.Encrypt(bytes)
		if err != nil {
			return err
		}
		w.Write(ciphertext)
	}
	return nil
}

type Algorithm interface {
	ParseKey(key []byte) error
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}
