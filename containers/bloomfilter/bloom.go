package bloom

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"github.com/liyue201/gostl/containers/bitmap"
	"math"
)

const Salt = "g9hmj2fhgr"

type BloomFilter struct {
	m uint64
	k uint
	b *bitmap.Bitmap
}

func New(m uint64, k uint) *BloomFilter {
	return &BloomFilter{
		m: m,
		k: k,
		b: bitmap.New(m),
	}
}

func EstimateParameters(n uint, p float64) (m uint64, k uint) {
	m = uint64(math.Ceil(-1 * float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)))
	k = uint(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return
}

func NewWithEstimates(n uint, fp float64) *BloomFilter {
	m, k := EstimateParameters(n, fp)
	return New(m, k)
}

func (this *BloomFilter) Add(val string) {
	hashs := this.GenHashs(val)
	for i := uint(0); i < this.k; i++ {
		this.b.Set(hashs[i] % this.m)
	}
}

func (this *BloomFilter) Contains(val string) bool {
	hashs := this.GenHashs(val)
	for i := uint(0); i < this.k; i++ {
		if !this.b.IsSet(hashs[i] % this.m) {
			return false
		}
	}
	return true
}

func (this *BloomFilter) GenHashs(key string) []uint64 {
	var hashValues []uint64
	data := []byte(Salt + key)
	for len(hashValues) < int(this.k) {
		data = hash512([]byte(data))
		temp := make([]uint64, len(data)/8)
		buf := bytes.NewBuffer(data)
		binary.Read(buf, binary.BigEndian, temp)
		hashValues = append(hashValues, temp...)
	}
	return hashValues
}

func hash512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}
