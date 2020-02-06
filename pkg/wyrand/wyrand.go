package wyrand

import (
	"math"
	"math/rand"
)

var state uint64

func init() {
	state = rand.Uint64()
}

func Uint64() uint64 {
	state += 0x60bee2bee120fc15
	tmp := state * 0xa3b195354a39b70d
	m1 := (tmp >> 64) ^ tmp
	tmp = m1 * 0x1b03738712fad5c9
	return (tmp >> 64) ^ tmp
}

func Uint64N(n uint64) uint64 {
	return Uint64() % n
}

func Uint8() uint8 {
	return uint8(Uint64() % math.MaxUint8)
}

func Uint8N(n uint8) uint8 {
	return uint8(Uint64() % uint64(n))
}

func Uint16() uint16 {
	return uint16(Uint64() % math.MaxUint16)
}

func Uint16N(n uint16) uint16 {
	return uint16(Uint64() % uint64(n))
}

func Uint32() uint32 {
	return uint32(Uint64() % math.MaxUint32)
}

func Uint32N(n uint32) uint32 {
	return uint32(Uint64() % uint64(n))
}

func Int() int {
	return int(uint(Uint64()) >> 1)
}

func IntN(n int) int {
	if n <= 0 {
		return 0
	}
	return int(uint(Uint64()) >> 1) % n
}

func Bool() bool {
	return (Uint64() % 2) == 1
}

func Float64() float64 {
	return float64(Uint64() >> 11) / (1 << 53)
}

func Float64N(n float64) float64  {
	return Float64() * n
}