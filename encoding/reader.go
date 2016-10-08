package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"io"
)

const bufSize = 1024

type DecodingReader struct {
	encoding   Encoding
	reader     io.Reader
	decodedBuf []byte
	eofFlag    bool
}

func NewDecodingReader(reader io.Reader, encoding Encoding) io.Reader {
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
