package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13Reader *rot13Reader) Read(b []byte) (size int, err error) {
	size, err = rot13Reader.r.Read(b)
	for i, letter := range b {
		b[i] = rot13(letter)
	}
	return
}

func rot13(b byte) byte {
	r := rune(b)
	switch {
	case r >= 'a' && r <= 'z':
		r = 'a' + (r-'a'+13)%26
	case r >= 'A' && r <= 'Z':
		r = 'A' + (r-'A'+13)%26
	}
	return byte(r)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
