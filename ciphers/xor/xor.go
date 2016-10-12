package xor

import (
	"github.com/cderwin/cryptopals/ciphers"
)

type XorAlgorithm struct {
	key    []byte
	keySet bool
}

func NewXorAlgorithm() *XorAlgorithm {
	return &XorAlgorithm{}
}

func (algo *XorAlgorithm) ParseKey(key []byte) error {
	algo.key = key
	return nil
}

func (algo *XorAlgorithm) Encrypt(plaintext []byte) ([]byte, error) {
	if !algo.keySet {
		return nil, ciphers.NoKeySetError
	}
	return xor(plaintext, algo.key), nil
}

func (algo *XorAlgorithm) Decrypt(ciphertext []byte) ([]byte, error) {
	if !algo.keySet {
		return nil, ciphers.NoKeySetError
	}
	return xor(ciphertext, algo.key), nil
}

func xor(text, key []byte) []byte {
	keyLength := len(key)
	result := make([]byte, len(text))
	for i := range text {
		result[i] = text[i] ^ key[i%keyLength]
	}

	return result
}
