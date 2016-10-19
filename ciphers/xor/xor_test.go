package xor

import (
	"bytes"
	"testing"
)

func expectXor(t *testing.T, text, key, expected []byte) {
	result := xor(text, key)
	if !bytes.Equal(result, expected) {
		t.Fatalf("Xor of text %q with key %q return %q, not %q", text, key, result, expected)
	}
}

func TestXor(t *testing.T) {
	expectXor(t, []byte{0, 17, 24, 6, 0}, []byte{0}, []byte{0, 17, 24, 6, 0})
	expectXor(t, []byte{0, 17, 24, 6, 0}, []byte{255}, []byte{^byte(0), ^byte(17), ^byte(24), ^byte(6), ^byte(0)})
	expectXor(t, []byte{0, 17, 24, 6, 0}, []byte{0, 255}, []byte{0, ^byte(17), 24, ^byte(6), 0})
	expectXor(t, []byte{0, 17, 24, 6, 0}, []byte{255, 0}, []byte{^byte(0), 17, ^byte(24), 6, ^byte(0)})
	expectXor(t, []byte("Hello, world!"), []byte(""), []byte("Hello, world!"))
	expectXor(t, []byte("Hello, world!"), []byte("foobar"), []byte(""))
	expectXor(t, []byte("Hello, world!"), []byte("fooberfoobarfoobar"), []byte(""))
	expectXor(t, []byte(""), []byte("foobar"), []byte(""))
}
