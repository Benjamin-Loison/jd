package jd

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
)

func hash(input []byte) [8]byte {
	h := fnv.New64a()
	h.Write(input)
	var a [8]byte
	binary.LittleEndian.PutUint64(a[:], h.Sum64())
	return a
}

type hashCodes [][8]byte

func (h hashCodes) Len() int {
	return len(h)
}

func (h hashCodes) Less(i, j int) bool {
	if bytes.Compare(h[i][:], h[j][:]) == -1 {
		return true
	} else {
		return false
	}
}

func (h hashCodes) Swap(i, j int) {
	h[j], h[i] = h[i], h[j]
}
