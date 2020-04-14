/* mt - implementation of Mersenne Twister PRNG in GOLang
Ported based on https://github.com/pigulla/mersennetwister

Copyright (C) 2020 Appditto LLC
Copyright (C) 1997 - 2002, Makoto Matsumoto and Takuji Nishimura,

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
	notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright
	notice, this list of conditions and the following disclaimer in the
	documentation and/or other materials provided with the distribution.

3. The names of its contributors may not be used to endorse or promote
	products derived from this software without specific prior written
	permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package rand

const (
	n                 = 624
	m                 = 397
	matrix_a   uint32 = 0x9908b0df
	upper_mask uint32 = 0x80000000
	lower_mask uint32 = 0x7fffffff
)

type MT19937 struct {
	mt  []uint32
	mti int
}

// Init - initializes new instance of MT19937, unseeded
func Init() *MT19937 {
	return &MT19937{
		mt:  make([]uint32, n),
		mti: n + 1,
	}
}

// Seed - seed the PRNG with value
func (mt19937 *MT19937) Seed(seed uint32) {
	var s uint32
	mt := mt19937.mt[:]
	mti := mt19937.mti

	mt[0] = seed

	for mti = 1; mti < n; mti++ {
		s = mt[mti-1] ^ (mt[mti-1] >> 30)
		mt[mti] =
			(((((s & 0xffff0000) >> 16) * 1812433253) << 16) + (s&0x0000ffff)*1812433253) + uint32(mti)
	}
	mt19937.mti = mti
}

// Uint32 - generates a random number on [0,0xffffffff]-interval
func (mt19937 *MT19937) Uint32() uint32 {
	var y uint32
	var kk int

	mt := mt19937.mt[:]
	mti := mt19937.mti
	mag01 := [2]uint32{0, matrix_a}

	if mti >= n {
		if mti == n+1 {
			mt19937.Seed(5489) // default seed
		}

		for kk = 0; kk < n-m; kk++ {
			y = (mt[kk] & upper_mask) | (mt[kk+1] & lower_mask)
			mt[kk] = mt[kk+m] ^ (y >> 1) ^ mag01[y&1]
		}

		for ; kk < n-1; kk++ {
			y = (mt[kk] & upper_mask) | (mt[kk+1] & lower_mask)
			mt[kk] = mt[kk+(m-n)] ^ (y >> 1) ^ mag01[y&1]
		}

		y = (mt[n-1] & upper_mask) | (mt[0] & lower_mask)
		mt[n-1] = mt[m-1] ^ (y >> 1) ^ mag01[y&1]

		mti = 0
	}

	y = mt[mti]
	y ^= (y >> 11)
	y ^= ((y << 7) & 0x9D2C5680)
	y ^= ((y << 15) & 0xEFC60000)
	y ^= (y >> 18)
	mt19937.mti = mti + 1
	return y
}

// Int31 generates a random number on [0,0x7fffffff]-interval
func (mt19937 *MT19937) Int31() int32 {
	return int32(mt19937.Uint32() & 0x7fffffff)
}

// Int31n generates a random number in [0, n]

// For implementation details, see:
// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction
// https://lemire.me/blog/2016/06/30/fast-random-shuffling
// https://golang.org/src/math/rand/rand.go
func (mt19937 *MT19937) Int31n(n int32) int32 {
	v := mt19937.Uint32()
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = mt19937.Uint32()
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
		}
	}
	return int32(prod >> 32)
}
