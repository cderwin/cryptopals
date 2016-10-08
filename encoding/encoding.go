package encoding

import (
	"bufio"
	"errors"
	"io"
)

type Encoding int

const (
	Binary = iota
	Hex
	Base64
)

var (
	InvalidEncoding = errors.New("Invalid Encoding")
)

func EncodeMain(in, out Encoding, reader io.Reader, writer io.Writer) error {
	decoder := NewDecodingReader(reader, in)
	scanner := bufio.NewScanner(decoder)
	encoder := NewEncodingWriter(writer, out)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		_, err := encoder.Write(bytes)
		if err != nil {
			return err
		}
	}

	return encoder.Close()
}

func ParseEncoding(arg string) (Encoding, error) {
	switch arg {
	case "bin":
		return Binary, nil
	case "hex":
		return Hex, nil
	case "b64":
		return Base64, nil
	default:
		return 0, InvalidEncoding
	}
}
