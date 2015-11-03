package main

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

var flagOut = flag.String("out", "./goget.out",
	"pathname of output file")
var flagSHA256 = flag.String("sha256", "",
	"expected SHA-256 (hex) of file")

func main() {
	flag.Parse()

	outDir, outBase := path.Split(*flagOut)
	out, err := ioutil.TempFile(outDir, outBase)
	if err != nil {
		log.Fatal(err)
	}

	source := flag.Arg(0)

	resp, err := http.Get(source)
	if err != nil {
		log.Fatal(err)
	}

	h := sha256.New()
	// We know from https://golang.org/pkg/hash/#Hash
	// that tee-ing to the hash "never returns an error".
	body := io.TeeReader(resp.Body, h)

	_, err = io.Copy(out, body)
	if err != nil {
		log.Fatal(err)
	}
	outPath := out.Name()
	out.Close()

	buf := bytes.NewBuffer(nil)
	bs := h.Sum([]byte{})
	for i := range bs {
		fmt.Fprintf(buf, "%02x", bs[i])
	}
	gotChecksum := buf.String()

	if gotChecksum != *flagSHA256 {
		_ = os.Remove(outPath)
		fmt.Fprint(os.Stderr,
			"Expected checksum ", *flagSHA256,
			"\nGot checksum ", gotChecksum, "\n")
		os.Exit(1)
	}
	err = os.Rename(outPath, *flagOut)
	if err != nil {
		log.Fatal(err)
	}
}
