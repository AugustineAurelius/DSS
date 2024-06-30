package hashbuffer

import (
	"crypto"
	"encoding/binary"
	"math/bits"
	"sync"
)

func init() {
	Default.reset()
}

var Default = &Buf512{
	function: crypto.SHA512,
	m:        sync.Mutex{},
}

type Buf512 struct {
	h        [8]uint64
	x        [128]byte
	nx       int
	len      uint64
	function crypto.Hash
	digest   [64]byte

	m sync.Mutex
}

func New() *Buf512 {
	b := &Buf512{
		function: crypto.SHA512,
		m:        sync.Mutex{},
	}
	b.reset()
	return b
}

func (b *Buf512) Hash(p []byte) [64]byte {
	b.m.Lock()
	defer b.m.Unlock()
	defer b.reset()

	b.write(p)

	l := b.len
	// Padding. Add a 1 bit and 0 bits until 112 bytes mod 128.
	var tmp [128 + 16]byte // padding + length buffer
	tmp[0] = 0x80

	var t uint64
	if l%128 < 112 {
		t = 112 - l%128
	} else {
		t = 128 + 112 - l%128
	}

	// Length in bits.
	l <<= 3
	padlen := tmp[:t+16]
	// Upper 64 bits are always zero, because len variable has type uint64,
	// and tmp is already zeroed at that index, so we can skip updating it.
	// binary.BigEndian.PutUint64(padlen[t+0:], 0)
	binary.BigEndian.PutUint64(padlen[t+8:], l)
	b.write(padlen)

	if b.nx != 0 {
		panic("d.nx != 0")
	}

	binary.BigEndian.PutUint64(b.digest[0:], b.h[0])
	binary.BigEndian.PutUint64(b.digest[8:], b.h[1])
	binary.BigEndian.PutUint64(b.digest[16:], b.h[2])
	binary.BigEndian.PutUint64(b.digest[24:], b.h[3])
	binary.BigEndian.PutUint64(b.digest[32:], b.h[4])
	binary.BigEndian.PutUint64(b.digest[40:], b.h[5])
	binary.BigEndian.PutUint64(b.digest[48:], b.h[6])
	binary.BigEndian.PutUint64(b.digest[56:], b.h[7])

	return b.digest

}

func (b *Buf512) reset() {
	b.h[0] = 0x6a09e667f3bcc908
	b.h[1] = 0xbb67ae8584caa73b
	b.h[2] = 0x3c6ef372fe94f82b
	b.h[3] = 0xa54ff53a5f1d36f1
	b.h[4] = 0x510e527fade682d1
	b.h[5] = 0x9b05688c2b3e6c1f
	b.h[6] = 0x1f83d9abfb41bd6b
	b.h[7] = 0x5be0cd19137e2179
	b.nx = 0
	b.len = 0
	b.digest = [64]byte{}
}

func (b *Buf512) write(p []byte) (nn int, err error) {
	nn = len(p)
	b.len += uint64(nn)
	if b.nx > 0 {
		n := copy(b.x[b.nx:], p)
		b.nx += n
		if b.nx == 128 {
			blockGeneric(b, b.x[:])
			b.nx = 0
		}
		p = p[n:]
	}
	if len(p) >= 128 {
		n := len(p) &^ (128 - 1)
		blockGeneric(b, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		b.nx = copy(b.x[:], p)
	}
	return

}

func blockGeneric(dig *Buf512, p []byte) {
	var w [80]uint64
	h0, h1, h2, h3, h4, h5, h6, h7 := dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7]
	for len(p) >= 128 {
		for i := 0; i < 16; i++ {
			j := i * 8
			w[i] = uint64(p[j])<<56 | uint64(p[j+1])<<48 | uint64(p[j+2])<<40 | uint64(p[j+3])<<32 |
				uint64(p[j+4])<<24 | uint64(p[j+5])<<16 | uint64(p[j+6])<<8 | uint64(p[j+7])
		}
		for i := 16; i < 80; i++ {
			v1 := w[i-2]
			t1 := bits.RotateLeft64(v1, -19) ^ bits.RotateLeft64(v1, -61) ^ (v1 >> 6)
			v2 := w[i-15]
			t2 := bits.RotateLeft64(v2, -1) ^ bits.RotateLeft64(v2, -8) ^ (v2 >> 7)

			w[i] = t1 + w[i-7] + t2 + w[i-16]
		}

		a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7

		for i := 0; i < 80; i++ {
			t1 := h + (bits.RotateLeft64(e, -14) ^ bits.RotateLeft64(e, -18) ^ bits.RotateLeft64(e, -41)) + ((e & f) ^ (^e & g)) + _K[i] + w[i]

			t2 := (bits.RotateLeft64(a, -28) ^ bits.RotateLeft64(a, -34) ^ bits.RotateLeft64(a, -39)) + ((a & b) ^ (a & c) ^ (b & c))

			h = g
			g = f
			f = e
			e = d + t1
			d = c
			c = b
			b = a
			a = t1 + t2
		}

		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
		h5 += f
		h6 += g
		h7 += h

		p = p[128:]
	}

	dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h0, h1, h2, h3, h4, h5, h6, h7
}

var _K = []uint64{
	0x428a2f98d728ae22,
	0x7137449123ef65cd,
	0xb5c0fbcfec4d3b2f,
	0xe9b5dba58189dbbc,
	0x3956c25bf348b538,
	0x59f111f1b605d019,
	0x923f82a4af194f9b,
	0xab1c5ed5da6d8118,
	0xd807aa98a3030242,
	0x12835b0145706fbe,
	0x243185be4ee4b28c,
	0x550c7dc3d5ffb4e2,
	0x72be5d74f27b896f,
	0x80deb1fe3b1696b1,
	0x9bdc06a725c71235,
	0xc19bf174cf692694,
	0xe49b69c19ef14ad2,
	0xefbe4786384f25e3,
	0x0fc19dc68b8cd5b5,
	0x240ca1cc77ac9c65,
	0x2de92c6f592b0275,
	0x4a7484aa6ea6e483,
	0x5cb0a9dcbd41fbd4,
	0x76f988da831153b5,
	0x983e5152ee66dfab,
	0xa831c66d2db43210,
	0xb00327c898fb213f,
	0xbf597fc7beef0ee4,
	0xc6e00bf33da88fc2,
	0xd5a79147930aa725,
	0x06ca6351e003826f,
	0x142929670a0e6e70,
	0x27b70a8546d22ffc,
	0x2e1b21385c26c926,
	0x4d2c6dfc5ac42aed,
	0x53380d139d95b3df,
	0x650a73548baf63de,
	0x766a0abb3c77b2a8,
	0x81c2c92e47edaee6,
	0x92722c851482353b,
	0xa2bfe8a14cf10364,
	0xa81a664bbc423001,
	0xc24b8b70d0f89791,
	0xc76c51a30654be30,
	0xd192e819d6ef5218,
	0xd69906245565a910,
	0xf40e35855771202a,
	0x106aa07032bbd1b8,
	0x19a4c116b8d2d0c8,
	0x1e376c085141ab53,
	0x2748774cdf8eeb99,
	0x34b0bcb5e19b48a8,
	0x391c0cb3c5c95a63,
	0x4ed8aa4ae3418acb,
	0x5b9cca4f7763e373,
	0x682e6ff3d6b2b8a3,
	0x748f82ee5defb2fc,
	0x78a5636f43172f60,
	0x84c87814a1f0ab72,
	0x8cc702081a6439ec,
	0x90befffa23631e28,
	0xa4506cebde82bde9,
	0xbef9a3f7b2c67915,
	0xc67178f2e372532b,
	0xca273eceea26619c,
	0xd186b8c721c0c207,
	0xeada7dd6cde0eb1e,
	0xf57d4f7fee6ed178,
	0x06f067aa72176fba,
	0x0a637dc5a2c898a6,
	0x113f9804bef90dae,
	0x1b710b35131c471b,
	0x28db77f523047d84,
	0x32caab7b40c72493,
	0x3c9ebe0a15c9bebc,
	0x431d67c49c100d4c,
	0x4cc5d4becb3e42b6,
	0x597f299cfc657e2a,
	0x5fcb6fab3ad6faec,
	0x6c44198c4a475817,
}
