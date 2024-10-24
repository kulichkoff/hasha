package hashing

import (
	"bufio"
	"crypto/sha256"
	"io"
)

type SHA256Hasher struct{}

func (h *SHA256Hasher) Hash(input []byte) []byte {
	hashed := sha256.Sum256(input)
	return hashed[:]
}

func (h *SHA256Hasher) HashStream(reader io.Reader) ([]byte, error) {
	rdr := bufio.NewReader(reader)
	hasher := sha256.New()

	if _, err := io.Copy(hasher, rdr); err != nil {
		return nil, err
	}

	hashed := hasher.Sum(nil)
	return hashed, nil
}
