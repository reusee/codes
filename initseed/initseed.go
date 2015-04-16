package initseed

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func init() {
	var seed int64
	binary.Read(crand.Reader, binary.LittleEndian, &seed)
	rand.Seed(seed)
}
