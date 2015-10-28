package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://golang.org/")
	if err != nil {
		log.Fatal(err)
	}

	summer, ch := checksummer(resp.Body)

	_, err = io.Copy(os.Stdout, summer)
	if err != nil {
		log.Fatal(err)
	}

	b := <-ch
	for i := range b {
		fmt.Fprintf(os.Stderr, "%02x", b[i])
	}
	fmt.Fprint(os.Stderr, "\n")
}

type Summer struct {
	src io.Reader
	h   hash.Hash
	c   chan<- []byte
}

func (r *Summer) Read(p []byte) (int, error) {
	n, err := r.src.Read(p)
	// From https://golang.org/pkg/hash/#Hash
	// "It never returns an error"
	_, _ = r.h.Write(p[:n])
	b := r.h.Sum([]byte{})
	if err == io.EOF {
		go func() {
			r.c <- b
		}()
	}
	return n, err
}

func checksummer(src io.Reader) (io.Reader, <-chan []byte) {
	ch := make(chan []byte)
	s := Summer{src, sha256.New(), ch}
	return &s, ch
}