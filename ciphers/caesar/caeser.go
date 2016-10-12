package caesar

import (
	"errors"

	"github.com/cderwin/cryptopals/ciphers"
)

var (
	InvalidKeyLengthError = errors.New("Key must be exactly one character")
	InvalidKeyValueError  = errors.New("Key must be an alphabetic character")
)

type CaesarAlgorithm struct {
	key    byte
	keySet bool
}

func NewCaesarAlgorithm() *CaesarAlgorithm {
	return &CaesarAlgorithm{}
}

func (caesar *CaesarAlgorithm) ParseKey(key []byte) error {
	if len(key) != 1 {
		return InvalidKeyLengthError
	}

	shift := key[0]
	if shift >= byte('a') && shift <= byte('z') {
		caesar.key = shift - byte('a')
		caesar.keySet = true
		return nil
	}

	if shift >= byte('A') && shift <= byte('Z') {
		caesar.key = shift - byte('A')
		caesar.keySet = true
		return nil
	}

	return InvalidKeyValueError
}

func (caesar *CaesarAlgorithm) Encrypt(plaintext []byte) ([]byte, error) {
	if !caesar.keySet {
		return nil, ciphers.NoKeySetError
	}

	ciphertext := make([]byte, len(plaintext))
	for i, char := range plaintext {
		ciphertext[i] = shiftChar(char, caesar.key)
	}

	return ciphertext, nil
}

func (caesar *CaesarAlgorithm) Decrypt(ciphertext []byte) ([]byte, error) {
	if !caesar.keySet {
		return nil, ciphers.NoKeySetError
	}

	shift := (26 - caesar.key) % 26
	plaintext := make([]byte, len(ciphertext))
	for i, char := range ciphertext {
		plaintext[i] = shiftChar(char, shift)
	}

	return plaintext, nil
}

func shiftChar(char, shift byte) byte {
	var baseIndex byte
	if char >= byte('a') && char <= byte('z') {
		baseIndex = byte('a')
	} else if char >= byte('A') && char <= byte('Z') {
		baseIndex = byte('A')
	} else {
		return char
	}

	charIndex := char - baseIndex
	newCharIndex := (charIndex + shift) % 26
	return baseIndex + newCharIndex
}
