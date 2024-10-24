package hashing

import (
	"fmt"
	"io"
)

type Hasher interface {
	Hash(input []byte) []byte
	HashStream(reader io.Reader) ([]byte, error)
}

const (
	SHA256 string = "sha256"
	MD5    string = "md5"
)

var hasher Hasher = &SHA256Hasher{}

// sha256 is used by default. You can set another algorithm
// for hash. Supported: sha256, md5
func SetAlgorithm(alg string) error {
	switch alg {
	case SHA256:
		hasher = &SHA256Hasher{}
		return nil
	case MD5:
		hasher = &MD5Hasher{}
		return nil
	default:
		return fmt.Errorf("not known hash algorithm: %s", alg)
	}
}

func Hash(data []byte) []byte {
	return hasher.Hash(data)
}

func HashStream(reader io.Reader) ([]byte, error) {
	return hasher.HashStream(reader)
}
