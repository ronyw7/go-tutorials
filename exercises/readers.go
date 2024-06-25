package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return len(b), nil
}

func rot13(b byte) byte {
	switch {
	case 'a' <= b && b <= 'z':
		return 'a' + (b - 'a' + 13) % 26
	case 'A' <= b && b <= 'Z':
		return 'A' + (b - 'A' + 13) % 26
	default:
		return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
