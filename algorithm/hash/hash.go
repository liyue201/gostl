package hash

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
)

// GenHashInts generate n hash values by the seed passed
func GenHashInts(seed []byte, n int) []uint64 {
	data := seed
	var hashInts []uint64
	for len(hashInts) < n {
		data := Hash512(data)
		temp := make([]uint64, len(data)/8)
		buf := bytes.NewBuffer(data)
		binary.Read(buf, binary.BigEndian, temp)
		hashInts = append(hashInts, temp...)
	}
	return hashInts[:n]
}

func Hash512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}
