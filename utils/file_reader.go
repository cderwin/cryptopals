package utils

import (
	"io"
	"os"
)

type FileReader struct {
	openFile  io.ReadCloser
	files     []string
	stdinOnly bool
}

func NewFileReader(args []string) (io.Reader, error) {
	if len(args) < 1 {
		reader := &FileReader{files: args, stdinOnly: true, openFile: os.Stdin}
		return reader, nil
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			return nil, err
		}
		reader := &FileReader{files: args[1:], stdinOnly: false, openFile: file}
		return reader, nil
	}
}

func (f *FileReader) Read(b []byte) (n int, err error) {
	n, err = f.openFile.Read(b)
	if err == io.EOF && len(f.files) > 0 && !f.stdinOnly {
		// Close file and open next
		f.openFile.Close()

		var filename string
		filename, f.files = f.files[0], f.files[1:]
		file, err := os.Open(filename)
		if err != nil {
			return 0, err
		}
		f.openFile = file

		n, err = file.Read(b)
	}

	return n, err
}

func (f *FileReader) Close() error {
	if f.stdinOnly {
		return os.Stdin.Close()
	}

	return f.openFile.Close()
}
