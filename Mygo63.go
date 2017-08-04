package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte)(n int,err error)  {
	rot.r.Read(b)
	length := len(b)

	for i := 0;i < length ; i++  {
		switch  {
		case b[i] >= 'a' && b[i] < 'n':
			fallthrough
		case b[i] >= 'A' && b[i] < 'N':
			b[i] = b[i] + 13
		case b[i] >= 'n' && b[i] <= 'z':
			fallthrough
		case b[i] >= 'N' && b[i] <= 'Z':
			b[i] = b[i] - 13
		}
	}

	return length,nil

}

func main() {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout,&r)

}
