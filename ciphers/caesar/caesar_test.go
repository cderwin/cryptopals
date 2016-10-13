package caesar

import (
	"testing"
)

type charShiftCase struct {
	char  byte
	shift byte
}

func TestShiftChar(t *testing.T) {
	cases := map[charShiftCase]byte{
		charShiftCase{byte('a'), byte(0)}: byte('a'),
		charShiftCase{byte('z'), byte(0)}: byte('z'),

		charShiftCase{byte('a'), byte(8)}: byte('i'),
		charShiftCase{byte('z'), byte(8)}: byte('h'),

		charShiftCase{byte('a'), byte(13)}: byte('n'),
		charShiftCase{byte('z'), byte(13)}: byte('m'),

		charShiftCase{byte('a'), byte(18)}: byte('s'),
		charShiftCase{byte('z'), byte(18)}: byte('r'),

		charShiftCase{byte('a'), byte(26)}: byte('a'),
		charShiftCase{byte('z'), byte(26)}: byte('z'),

		charShiftCase{byte('a'), byte(32)}: byte('g'),
		charShiftCase{byte('z'), byte(32)}: byte('f'),
	}

	for args, expected := range cases {
		result := shiftChar(args.char, args.shift)
		if result != expected {
			t.Errorf("Result (%q) not equal to expected (%q)", []byte{result}, []byte{expected})
		}
	}
}

func TestParseKey(t *testing.T) {
	cases := map[string]CaesarAlgorithm{
		"a": CaesarAlgorithm{key: byte(0), keySet: true},
		"t": CaesarAlgorithm{key: byte(19), keySet: true},
		"A": CaesarAlgorithm{key: byte(0), keySet: true},
		"T": CaesarAlgorithm{key: byte(19), keySet: true},
	}

	for key, expected := range cases {
		algo := NewCaesarAlgorithm()
		algo.ParseKey([]byte(key))
		if *algo != expected {
			t.Errorf("Result (`%+v`) not equal to expected (`%+v`)", algo, expected)
		}
	}
}

func TestParseKeyErrors(t *testing.T) {
	cases := map[string]error{
		"hello": InvalidKeyLengthError,
		"":      InvalidKeyLengthError,
		"7":     InvalidKeyValueError,
		"\x07":  InvalidKeyValueError,
	}

	for key, expected := range cases {
		algo := NewCaesarAlgorithm()
		result := algo.ParseKey([]byte(key))
		if result != expected {
			t.Errorf("Result error (%s) not equal to expected error (%s)", result, expected)
		}
	}
}
