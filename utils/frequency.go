package utils

// IO Util Reader

import (
	"bytes"
	"fmt"
)

//
// Scoring frequency distributions
//

var englishFreq = map[byte]float64{
	byte('a'): 8.12,
	byte('b'): 1.49,
	byte('c'): 2.71,
	byte('d'): 4.32,
	byte('e'): 12.02,
	byte('f'): 2.30,
	byte('g'): 2.03,
	byte('h'): 5.92,
	byte('i'): 7.31,
	byte('j'): 0.10,
	byte('k'): 0.69,
	byte('l'): 3.98,
	byte('m'): 2.61,
	byte('n'): 6.95,
	byte('o'): 7.68,
	byte('p'): 1.82,
	byte('q'): 0.11,
	byte('r'): 6.02,
	byte('s'): 6.28,
	byte('t'): 9.10,
	byte('u'): 2.88,
	byte('v'): 1.11,
	byte('w'): 2.09,
	byte('x'): 0.17,
	byte('y'): 2.11,
	byte('z'): 0.07,
	byte(' '): 16.39,
}

type Frequency struct {
	table  map[byte]int
	length int
}

func NewFrequency(input []byte) *Frequency {
	input = bytes.ToLower(input)

	frequency := make(map[byte]int)
	for _, char := range input {
		frequency[char] += 1
	}

	return &Frequency{table: frequency, length: len(input)}
}

func (freq *Frequency) Score() float64 {
	// Chi squared
	var score float64
	for char, count := range freq.table {
		english, ok := englishFreq[char]
		if !ok {
			english = 0.0001
		}

		expected := english * float64(freq.length)
		score += (float64(count) - expected) * (float64(count) - expected) / expected
	}
	return score
}

func (freq *Frequency) Compare() {
	// Prints put character-by-character comparison with English frequencies
	fmt.Printf("Character frequency comparison.\n")
	fmt.Printf("Letter\tSample\tEnglish\n")

	for char, freq := range freq.table {
		fmt.Printf("%c     \t%-06f  \t%-06f\n", char, freq, englishFreq[char])
	}
}
