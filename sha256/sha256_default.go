// +build !linux 386

package sha256

import (
	"crypto/sha256"
	"hash"
)

type hashWrapper struct {
	h hash.Hash
}

func New() (Hasher, error) {
	return hashWrapper{h: sha256.New()}, nil
}

func (h hashWrapper) Write(data []byte) (n int, err error) {
	return h.h.Write(data)
}

func (h hashWrapper) Sum() (rv [Size]byte, err error) {
	copy(rv[:], h.h.Sum(nil))
	return rv, nil
}

func (h hashWrapper) Close() {}
