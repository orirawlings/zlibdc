package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "zlibdc: Error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	r, err := zlib.NewReader(os.Stdin)
	check(err)
	defer r.Close()
	_, err = io.Copy(os.Stdout, r)
	check(err)
}
