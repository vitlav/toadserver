package wire

import (
	"bytes"
	"crypto/sha256"
	"code.google.com/p/go.crypto/ripemd160"

	. "github.com/tendermint/tendermint/common"
)

func BinaryBytes(o interface{}) []byte {
	w, n, err := new(bytes.Buffer), new(int64), new(error)
	WriteBinary(o, w, n, err)
	if *err != nil {
		PanicSanity(*err)
	}
	return w.Bytes()
}

func JSONBytes(o interface{}) []byte {
	w, n, err := new(bytes.Buffer), new(int64), new(error)
	WriteJSON(o, w, n, err)
	if *err != nil {
		PanicSanity(*err)
	}
	return w.Bytes()
}

// NOTE: does not care about the type, only the binary representation.
func BinaryEqual(a, b interface{}) bool {
	aBytes := BinaryBytes(a)
	bBytes := BinaryBytes(b)
	return bytes.Equal(aBytes, bBytes)
}

// NOTE: does not care about the type, only the binary representation.
func BinaryCompare(a, b interface{}) int {
	aBytes := BinaryBytes(a)
	bBytes := BinaryBytes(b)
	return bytes.Compare(aBytes, bBytes)
}

// NOTE: only use this if you need 32 bytes.
func BinarySha256(o interface{}) []byte {
	hasher, n, err := sha256.New(), new(int64), new(error)
	WriteBinary(o, hasher, n, err)
	if *err != nil {
		PanicSanity(*err)
	}
	return hasher.Sum(nil)
}

// NOTE: The default hash function is Ripemd160.
func BinaryRipemd160(o interface{}) []byte {
	hasher, n, err := ripemd160.New(), new(int64), new(error)
	WriteBinary(o, hasher, n, err)
	if *err != nil {
		PanicSanity(*err)
	}
	return hasher.Sum(nil)
}
