package main

import "fmt"

type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{
		make(chan byte, 1024),
	}
}
func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}
func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}
func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}
func main() {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("hello "))
		writer.Write([]byte("my name is raye"))
	}()
	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
}
