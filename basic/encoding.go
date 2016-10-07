package basic

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"flag"
	"io"
	"os"

	"github.com/cderwin/cryptopals/utils"
)

const (
	Binary = iota
	Hex
	Base64
)

var (
	InvalidEncoding = errors.New("Invalid Encoding")
)

func EncodeMain(args []string) error {
	flagSet := flag.NewFlagSet("encoding", flag.ExitOnError)
	inEncoding := flagSet.String("in", "", "The encoding of the input data.  Options are bin (binary), hex, or b64 (base64).")
	outEncoding := flagSet.String("out", "", "The desired encoding of the output data.  Options are bin (binary), hex, or b64 (base64).")
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	inputEncoding, err := parseEncoding(*inEncoding)
	if err != nil {
		return err
	}
	outputEncoding, err := parseEncoding(*outEncoding)
	if err != nil {
		return err
	}

	filereader, err := utils.NewFileReader(flagSet.Args())
	if err != nil {
		return err
	}

	reader := NewDecodingReader(filereader, inputEncoding)
	scanner := bufio.NewScanner(reader)
	writer := NewEncodingWriter(os.Stdout, outputEncoding)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		_, err = writer.Write(bytes)
		if err != nil {
			return err
		}
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

func parseEncoding(arg string) (int, error) {
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

const bufSize = 1024

type DecodingReader struct {
	encoding   int
	reader     io.Reader
	decodedBuf []byte
	eofFlag    bool
}

func NewDecodingReader(reader io.Reader, encoding int) io.Reader {
	if encoding == Base64 {
		reader = base64.NewDecoder(base64.StdEncoding, reader)
	}
	return &DecodingReader{reader: reader, encoding: encoding}
}

func (r *DecodingReader) Read(buffer []byte) (int, error) {
	// return eof if flag is set
	if r.eofFlag {
		return 0, io.EOF
	}

	// read from decoded buffer and return in output buffer is full
	maxSize := len(buffer)
	count := 0
	if maxSize <= len(r.decodedBuf) {
		count += copy(buffer, r.decodedBuf[:maxSize])
		r.decodedBuf = r.decodedBuf[maxSize:]
		return count, nil
	}

	// fill as much as output buffer as possible
	count += copy(buffer, r.decodedBuf)
	r.decodedBuf = make([]byte, 0)

	// until output biffer is full, decode line into decode buffer and write to output buffer
	for count < maxSize {
		// Decode line into decoded buffer and handle error
		err := r.fillBuffer()
		if err != nil {
			if err == io.EOF {
				r.eofFlag = true
				return count, nil
			}

			return 0, err
		}

		// Return if output buffer can be filled
		if maxSize < count+len(r.decodedBuf) {
			count += copy(buffer[count:], r.decodedBuf[:maxSize-count])
			r.decodedBuf = r.decodedBuf[maxSize-count:]
			return count, nil
		}

		// Otherwise fill as much of output buffer as possible and continue
		count += copy(buffer[count:], r.decodedBuf)
		r.decodedBuf = make([]byte, 0)
	}

	return count, nil
}

func (r *DecodingReader) fillBuffer() error {
	switch r.encoding {
	case Binary, Base64:
		r.decodedBuf = make([]byte, bufSize)
		n, err := r.reader.Read(r.decodedBuf)
		r.decodedBuf = r.decodedBuf[:n]
		return err
	case Hex:
		encodedBuffer := make([]byte, 2*bufSize)
		n, err := r.reader.Read(encodedBuffer)
		encodedBuffer = encodedBuffer[:n]
		if err != nil {
			return err
		}

		r.decodedBuf = make([]byte, bufSize)
		n, err = hex.Decode(r.decodedBuf, encodedBuffer)
		r.decodedBuf = r.decodedBuf[:n]
		return err
	}

	return InvalidEncoding
}

type EncodingWriter struct {
	encoding int
	writer   io.Writer
}

func NewEncodingWriter(writer io.Writer, encoding int) io.WriteCloser {
	if encoding == Base64 {
		writer = base64.NewEncoder(base64.StdEncoding, writer)
	}
	return &EncodingWriter{encoding: encoding, writer: writer}
}

func (w *EncodingWriter) Write(buf []byte) (int, error) {
	switch w.encoding {
	case Binary, Base64:
		return w.writer.Write(buf)
	case Hex:
		encodedBytes := make([]byte, 2*len(buf))
		hex.Encode(encodedBytes, buf)
		n, err := w.writer.Write(encodedBytes)
		return n / 2, err
	}

	return 0, InvalidEncoding
}

func (w *EncodingWriter) Close() error {
	if w.encoding != Base64 {
		return nil
	}

	return w.writer.(io.Closer).Close()
}
