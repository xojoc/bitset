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

//

/*
 Package bitset implements a BitSet data structure.

 A BitSet is a mapping between unsigned integers and boolean values.
 You can Set, Clear, Toggle single bits or Union, Intersect, Difference sets.

 BitSets are fixed-sized if you access an index outside of the bounds, panic will be called.

 Many methods return the BitSet they modify which enables chaining of operations.

 All methods are in place to avoid futile memory usage. If you want to keep
 the original BitSet simply Clone it.
*/
package bitset

// Bit tricks: http://graphics.stanford.edu/~seander/bithacks.html

// Bits per word
const bpw int = 64

type BitSet struct {
	// underlying vector
	v []uint64
	// n bits in the last word
	nb int
}

func maskLastBits(n int) uint64 {
	return ^(^uint64(0) << uint(n))
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

// Set sets the bit at index i.
// Returns s.
func (s *BitSet) Set(i int) *BitSet {
	s.v[i/bpw] |= 1 << uint(i%bpw)
	return s
}

// SetAll sets all the bits.
// Returns s.
func (s *BitSet) SetAll() *BitSet {
	for i := range s.v {
		s.v[i] = 1
	}
	return s
}

// Clear clears the bit at index i.
// Returns s.
func (s *BitSet) Clear(i int) *BitSet {
	s.v[i/bpw] &= ^(1 << uint(i%bpw))
	return s
}

// ClearAll clears all the bits.
// Returns s.
func (s *BitSet) ClearAll() *BitSet {
	for i := range s.v {
		s.v[i] = 0
	}
	return s
}

// Get gets the bit at index i.
func (s *BitSet) Get(i int) bool {
	// FIXME: should we panic if i%bpw > nb ?
	return (s.v[i/bpw] & (1 << uint(i%bpw))) != 0
}

// Toggle inverts the bit at index i.
// Returns s.
func (s *BitSet) Toggle(i int) *BitSet {
	s.v[i/bpw] ^= 1 << uint(i%bpw)
	return s
}

// Len returns the number of bits in s.
func (s *BitSet) Len() int {
	if len(s.v) == 0 {
		return 0
	}
	return (len(s.v)-1)*bpw + s.nb
}

// Count counts the number of set bits.
func (s *BitSet) Count() int {
	c := 0
	for _, e := range s.v {
		for e != 0 {
			c++
			e &= e - 1
		}
	}
	return c
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

// Clone makes a copy of s.
func (s *BitSet) Clone() *BitSet {
	b := New(s.Len())
	copy(b.v, s.v)
	b.nb = s.nb
	return b
}

// Equal returns true if a and b have the same Len and the same bits set, false otherwise.
func (a *BitSet) Equal(b *BitSet) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < len(a.v)-1; i++ {
		if a.v[i] != b.v[i] {
			return false
		}
	}
	m := maskLastBits(a.nb)
	if a.v[len(a.v)-1]&m != b.v[len(b.v)-1]&m {
		return false
	}
	return true
}

// All returns true if all bits are set, false otherwise.
func (s *BitSet) All() bool {
	for _, e := range s.v[:len(s.v)-1] {
		if e != ^uint64(0) {
			return false
		}
	}
	m := maskLastBits(s.nb)
	if s.v[len(s.v)-1]&m != m {
		return false
	}
	return true
}

// Any returns true if any bit is set, false otherwise.
func (s *BitSet) Any() bool {
	for _, e := range s.v[:len(s.v)-1] {
		if e != 0 {
			return true
		}
	}
	m := maskLastBits(s.nb)
	if s.v[len(s.v)-1]&m != 0 {
		return true
	}
	return false
}

// None returns true if no bit is set, false otherwise.
func (s *BitSet) None() bool {
	return !s.Any()
}

// Complement inverts all the bits of s.
// Returns s.
func (s *BitSet) Complement() *BitSet {
	for i := range s.v {
		s.v[i] = ^s.v[i]
	}
	return s
}

// Union stores in a the true bits from either a or b.
// If the length of b is greater than the length of a,
// a is extended to include all the extra bits from b.
// Returns a.
func (a *BitSet) Union(b *BitSet) *BitSet {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] | b.v[i]
	}

	if len(b.v) > len(a.v) {
		a.v = append(a.v, b.v[len(a.v):]...)
		a.nb = b.nb
	} else if len(b.v) == len(a.v) && b.nb > a.nb {
		a.nb = b.nb
	}
	return a
}

// Insersect stores in a the true bits common to both a and b.
// If the length of a is less than the length of b,
// a is truncated.
// Returns a.
func (a *BitSet) Intersect(b *BitSet) *BitSet {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] & b.v[i]
	}
	if len(a.v) > len(b.v) {
		a.v = a.v[:len(b.v)]
		a.nb = b.nb
	} else if len(a.v) == len(b.v) && b.nb < a.nb {
		a.nb = b.nb
	}
	return a
}

// Difference stores in a the true bits present in a and not in b.
// Returns a.
func (a *BitSet) Difference(b *BitSet) *BitSet {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] & ^b.v[i]
	}
	return a
}
