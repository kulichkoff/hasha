package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"hasha/internal/hashing"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags] <file>|<text>\n", os.Args[0])
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	alg := flag.String("a", "sha256", "define hash algorithm")
	flag.Usage = usage
	flag.Parse()

	if err := hashing.SetAlgorithm(*alg); err != nil {
		fmt.Printf("%s is unknown algorith", *alg)
		return
	}

	if isInPipe() {
		sum, err := hashing.HashStream(os.Stdin)
		if err != nil {
			fmt.Printf("failed to read pipe: %v", err)
			fmt.Println("supported: sha256, md5")
			return
		}

		hashed := hex.EncodeToString(sum)
		fmt.Println(hashed)
		return
	}

	args := flag.Args()
	if len(args) != 1 || args[0] == "-h" {
		fmt.Println("Usage:\thasha [text]")
		return
	}

	if hashed, err := getFileHash(args[0]); err == nil {
		fmt.Println(hashed)
		return
	}

	hashed := getHash([]byte(args[0]))
	fmt.Println(hashed)
}

func isInPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func getHash(input []byte) string {
	data := append(input, '\n')
	hash := hashing.Hash(data)
	hashHex := hex.EncodeToString(hash[:])
	return hashHex
}

func getFileHash(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hashed, err := hashing.HashStream(f)
	if err != nil {
		return "", err
	}

	hashHex := hex.EncodeToString(hashed)
	return hashHex, nil
}
