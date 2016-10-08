package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"io"
)

type EncodingWriter struct {
	encoding Encoding
	writer   io.Writer
}

func NewEncodingWriter(writer io.Writer, encoding Encoding) io.WriteCloser {
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
