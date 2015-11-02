package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

var flagSHA256 = flag.String("sha256", "",
	"SHA-256 from the SHA-2 standard (as hexstring)")

var sha256RE = regexp.MustCompile("[[:xdigit:]]{64}")

func main() {
	flag.Parse()

	if !sha256RE.MatchString(*flagSHA256) {
		log.Fatalf("-sha256 \"%v\" should be a SHA-256 hexadecimal string", *flagSHA256)
	}

	resp, err := http.Get("http://golang.org/")
	if err != nil {
		log.Fatal(err)
	}

	h := sha256.New()
	// We know from https://golang.org/pkg/hash/#Hash
	// that tee-ing to the hash "never returns an error".
	body := io.TeeReader(resp.Body, h)

	_, err = io.Copy(os.Stdout, body)
	if err != nil {
		log.Fatal(err)
	}

	b := h.Sum([]byte{})
	for i := range b {
		fmt.Fprintf(os.Stderr, "%02x", b[i])
	}
	fmt.Fprint(os.Stderr, "\n")
}
