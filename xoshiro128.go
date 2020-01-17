package xoshiro128

type Xoshiro128 struct {
	s []uint32
}

func CreateXoshiro128(a, b, c, d uint32) Xoshiro128 {
	return Xoshiro128{
		s: []uint32{a, b, c, d},
	}
}

func rotl(x uint32, k uint32) uint32 {
	return (x << k) | (x >> (32 - k))
}

func (x *Xoshiro128) Next() uint32 {
	res := rotl(x.s[1]*5, 7) * 9
	t := x.s[1] << 9

	x.s[2] ^= x.s[0]
	x.s[3] ^= x.s[1]
	x.s[1] ^= x.s[2]
	x.s[0] ^= x.s[3]

	x.s[2] ^= t

	x.s[3] = rotl(x.s[3], 11)

	return res
}

func (x *Xoshiro128) Random() float64 {
	low := uint64(x.Next())
	high := uint64(x.Next())

	mul := (((high >> 11) << 32) | low)
	return float64(mul) * 1.110223e-16
}
