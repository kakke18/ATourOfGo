package main

import "../../../../Go/misc/tour/src/golang.org/x/tour/reader"

type MyReader struct{}

func (read MyReader) Read(b []byte) (int, error){
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
