package basic

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
)

// Xor Encrypt Hook

func XorHook(r io.Reader, key []byte, w io.Writer) error {
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

// Xor Decryption

// A "candidate" for decryption is a (key, plaintext, score) tuple.  Lower scores are better.

type DecryptionCandidate struct {
	plaintext []byte
	key       []byte
	score     float64
}

func NewCandidate(ciphertext, key []byte) DecryptionCandidate {
	plaintext := xor(ciphertext, key)
	score := NewFrequency(plaintext).Score()
	return DecryptionCandidate{plaintext: plaintext, key: key, score: score}
}

// type `DecryptionCandidates` is a sortable slice of `[]DecryptionCandidate`

type DecryptionCandidates []DecryptionCandidate

func (candidates *DecryptionCandidates) Len() int {
	return len(*candidates)
}

func (candidates *DecryptionCandidates) Less(i, j int) bool {
	list := *candidates
	return list[i].score < list[j].score
}

func (candidates *DecryptionCandidates) Swap(i, j int) {
	list := *candidates
	list[i], list[j] = list[j], list[i]
}

var (
	InvalidKeyLen     = errors.New("Invalid key length (must be 1)")
	InvalidNumResults = errors.New("Invalid number of results (must be greater than 0)")
)

// Hook to break xor

func BreakXorHook(r io.Reader, w io.Writer, keyLen int, numResults int) error {
	if keyLen != 1 {
		return InvalidKeyLen
	}

	if numResults < 1 {
		return InvalidNumResults
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ciphertext := scanner.Bytes()

		candidates := make(DecryptionCandidates, 0, numResults)
		for key := byte(0); true; key++ {
			newCandidate := NewCandidate(ciphertext, []byte{key})

			if len(candidates) < numResults {
				candidates := append(candidates, newCandidate)
				sort.Sort(&candidates)
			} else if newCandidate.score < candidates[len(candidates)-1].score {
				candidates[len(candidates)-1] = newCandidate
				sort.Sort(&candidates)
			}

			// Break manually due to overflow
			if key == 0xff {
				break
			}
		}

		for _, candidate := range candidates {
			output := fmt.Sprintf("Key: %#x\nPlaintext: %q\nScore: %-02f\n\n", candidate.key, candidate.plaintext, candidate.score)
			w.Write([]byte(output))
		}
	}

	return nil
}
