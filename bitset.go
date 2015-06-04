/* Copyright (C) 2015 by Alexandru Cojocaru */

/* This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>. */

// Package bitset implements a BitSet data structure.
package bitset

// Bits per word
const bpw int = 64

type BitSet struct {
	// underlying vector
	v []uint64
	// n bits in the last word
	nb int
}

// New returns an initialized BitSet. All bits are set to false.
func New(len int) *BitSet {
	if len <= 0 {
		return &BitSet{}
	}
	nw := (len-1)/bpw + 1
	nb := (len-1)%bpw + 1
	return &BitSet{v: make([]uint64, nw), nb: nb}
}

// Set sets the bit at index i to b.
func (s *BitSet) Set(i int, b bool) {
	if b == true {
		s.v[i/bpw] |= 1 << uint(i%bpw)
	} else {
		s.v[i/bpw] &= ^(1 << uint(i%bpw))
	}
}

// Get gets the bit at index i.
func (s *BitSet) Get(i int) bool {
	// FIXME: should we panic if i%bpw > nb ?
	b := s.v[i/bpw] & (1 << uint(i%bpw))
	if b == 0 {
		return false
	}
	return true
}

// Toggle inverts the bit at index i.
func (s *BitSet) Toggle(i int) {
	s.v[i/bpw] ^= 1 << uint(i%bpw)
}

// Len returns the number of bits in s.
func (s *BitSet) Len() int {
	if len(s.v) == 0 {
		return 0
	}
	return (len(s.v)-1)*bpw + s.nb
}

// String returns a string representation of s.
func (s *BitSet) String() string {
	str := ""
	for i := 0; i < s.Len(); i++ {
		if s.Get(i) == true {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

// Copy makes a copy of s.
func (s *BitSet) Copy() *BitSet {
	b := New(s.Len())
	copy(b.v, s.v)
	b.nb = s.nb
	return b
}

// Union stores in a the true bits from either a or b.
// If the length of b is greater than the length of a,
// a is extended to include all the extra bits from b.
func (a *BitSet) Union(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] | b.v[i]
	}

	if len(b.v) > len(a.v) {
		a.v = append(a.v, b.v[len(a.v):]...)
		a.nb = b.nb
	} else if len(b.v) == len(a.v) && b.nb > a.nb {
		a.nb = b.nb
	}
}

// Insersect stores in a the true bits common to both a and b.
// If the length of a is less than the length of b,
// a is truncated.
func (a *BitSet) Intersect(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] & b.v[i]
	}
	if len(a.v) > len(b.v) {
		a.v = a.v[:len(b.v)]
		a.nb = b.nb
	} else if len(a.v) == len(b.v) && b.nb < a.nb {
		a.nb = b.nb
	}
}

// Complement inverts all the bits of s.
func (s *BitSet) Complement() {
	for i := 0; i < len(s.v); i++ {
		s.v[i] = ^s.v[i]
	}
}

// Difference stores in a the true bits present in a and not in b.
func (a *BitSet) Difference(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] & ^b.v[i]
	}
}
