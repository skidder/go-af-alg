package sha256

import (
	"crypto/sha256"
)

const (
	BlockSize = sha256.BlockSize
	Size      = sha256.Size
)

type Hasher interface {
	Write(data []byte) (n int, err error)
	Sum() ([Size]byte, error)
	Close()
}

func SHA256(data []byte) (result [Size]byte, err error) {
	hash, err := New()
	if err != nil {
		return result, err
	}
	defer hash.Close()
	if _, err := hash.Write(data); err != nil {
		return result, err
	}
	return hash.Sum()
}
