package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

type HashMode string

const (
	SHA256 HashMode = "sha256"

	SHA384 HashMode = "sha384"
	SHA512 HashMode = "sha512"
)

// implement the flag.Value interface

func (h *HashMode) String() string {
	return string(*h)
}

func (h *HashMode) Set(value string) error {
	mode := HashMode(value)
	switch mode {
	case SHA256, SHA384, SHA512:
	default:
		return fmt.Errorf("unsupported hash mode: %s", value)
	}
	*h = mode
	return nil
}

func main() {
	var mode HashMode = SHA256
	flag.Var(&mode, "mode", "the hashing algorithm to use (supported: sha256, sha384, sha512)")

	usage := flag.Usage
	flag.Usage = func() {
		usage()
		fmt.Fprintf(flag.CommandLine.Output(), "Example: %s -mode=sha512 input_string1\n", os.Args[0]) // ignore error
	}

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "No input to hash provided\n") // ignore error
		os.Exit(1)
	}

	input := args[0]
	var res []byte
	switch mode {
	case SHA256:
		hash := sha256.Sum256([]byte(input))
		res = hash[:]
	case SHA384:
		hash := sha512.Sum384([]byte(input))
		res = hash[:]
	case SHA512:
		hash := sha512.Sum512([]byte(input))
		res = hash[:]
	}

	fmt.Printf("%x\n", res)
}
