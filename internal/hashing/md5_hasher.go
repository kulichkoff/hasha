package hashing

import (
	"bufio"
	"crypto/md5"
	"io"
)

type MD5Hasher struct{}

func (h *MD5Hasher) Hash(input []byte) []byte {
	hashed := md5.Sum(input)
	return hashed[:]
}

func (h *MD5Hasher) HashStream(reader io.Reader) ([]byte, error) {
	rdr := bufio.NewReader(reader)
	hasher := md5.New()

	if _, err := io.Copy(hasher, rdr); err != nil {
		return nil, err
	}

	hashed := hasher.Sum(nil)
	return hashed, nil
}
