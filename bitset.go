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

/*
 Package bitset implements a BitSet data structure.

 A BitSet is a mapping between unsigned integers and boolean values.
 You can Set, Clear, Toggle single bits or Union, Intersect, Difference sets.

 Indexes start at 0. Ranges have the first index included and the second
 one excluded (like go slices).

 BitSets are dynamicaly-sized they grow and shrink automatically.

 All methods modify their receiver in place to avoid futile memory usage.
 If you want to keep the original BitSet simply Clone it.

 Use Clone when you want to copy a BitSet. Plese note that this will
 *not* work:
     var x BitSet
     x.Add(1)
     y := x  // wrong! use Clone
     y.Add(2)


 If you wonder why you should use this package and not math/big see:
 http://typed.pw/a/29
*/
package bitset

// TODO: nextbit prevbit

// Bit tricks: http://graphics.stanford.edu/~seander/bithacks.html

// Bits per word
const bpw int = 8 << (^uint(0)>>8&1 + ^uint(0)>>16&1 + ^uint(0)>>32&1)

type BitSet struct {
	// underlying vector
	v []uint
}

// FIXME: see if unsafe.SetCap can help reducing memory usage.
// FIXME: better yet a[l:h:m]

func (s *BitSet) autoShrink() {
	for i := len(s.v) - 1; i >= 0; i-- {
		if s.v[i] == 0 {
			s.v = s.v[:len(s.v)-1]
		} else {
			break
		}
	}
}

// Clone makes a copy of s.
func (s *BitSet) Clone() *BitSet {
	s.autoShrink()
	b := &BitSet{}
	// TODO: benchmark append
	b.v = make([]uint, len(s.v))
	copy(b.v, s.v)
	return b
}

// String returns a string representation of s.
func (s *BitSet) String() string {
	s.autoShrink()
	b := make([]byte, s.Len())
	for i := 0; i < s.Len(); i++ {
		if s.Get(i) == true {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	return string(b)
}

// Set sets the bit at index i.
func (s *BitSet) Set(i int) {
	for (i/bpw + 1) > len(s.v) {
		// NOTE: len(s.v) may have been decreased by Intersect,
		// which means that there may be spurious bits after
		// len(s.v)
		s.v = append(s.v, 0)
	}
	s.v[i/bpw] |= 1 << uint(i%bpw)
}

// SetRange sets the bits between i (included) and j (excluded).
func (s *BitSet) SetRange(i, j int) {
	for k := i; k < j; k++ {
		s.Set(k)
	}
}

// Clear clears the bit at index i.
func (s *BitSet) Clear(i int) {
	if (i/bpw + 1) > len(s.v) {
		return
	}
	s.v[i/bpw] &= ^(1 << uint(i%bpw))
}

// ClearRange clears the bits between i (included) and j (excluded).
func (s *BitSet) ClearRange(i, j int) {
	for k := i; k < j; k++ {
		s.Clear(k)
	}
}

// Toggle inverts the bit at index i.
func (s *BitSet) Toggle(i int) {
	if i/bpw+1 > len(s.v) {
		s.Set(i)
	} else {
		s.v[i/bpw] ^= 1 << uint(i%bpw)
	}
}

// ToggleRange inverts the bits between i (included) and j (excluded).
func (s *BitSet) ToggleRange(i, j int) {
	for k := i; k < j; k++ {
		s.Toggle(k)
	}
}

// Get gets the bit at index i.
func (s *BitSet) Get(i int) bool {
	if i/bpw+1 > len(s.v) {
		return false
	}
	return (s.v[i/bpw] & (1 << uint(i%bpw))) != 0
}

// Len returns the number of bits up to and including the highest bit set.
func (s *BitSet) Len() int {
	s.autoShrink()
	if len(s.v) == 0 {
		return 0
	}
	e := s.v[len(s.v)-1]
	c := 0
	for e != 0 {
		e = e >> 1
		c++
	}
	return (len(s.v)-1)*bpw + c
}

// Any returns true if any bit is set, false otherwise.
func (s *BitSet) Any() bool {
	s.autoShrink()
	for _, e := range s.v {
		if e != 0 {
			return true
		}
	}
	return false
}

// None returns true if no bit is set, false otherwise.
func (s *BitSet) None() bool {
	return !s.Any()
}

func countBits(e uint) int {
	c := 0
	for e != 0 {
		c++
		e &= e - 1
	}
	return c
}

// Cardinality counts the number of set bits.
func (s *BitSet) Cardinality() int {
	s.autoShrink()
	c := 0
	for _, e := range s.v {
		c += countBits(e)
	}
	return c
}

// Equal returns true if a and b have the same bits set, false otherwise.
func (a *BitSet) Equal(b *BitSet) bool {
	a.autoShrink()
	b.autoShrink()
	if a.Len() != b.Len() {
		return false
	}
	// TODO: bench for e, i :=
	for i := 0; i < len(a.v); i++ {
		if a.v[i] != b.v[i] {
			return false
		}
	}
	return true
}

// SuperSet returns true if a is a super set of b, false otherwise.
func (a *BitSet) SuperSet(b *BitSet) bool {
	a.autoShrink()
	b.autoShrink()
	if b.Len() > a.Len() {
		return false
	}
	for i := 0; i < len(b.v); i++ {
		if b.v[i] & ^a.v[i] != 0 {
			return false
		}
	}
	return true
}

// Union stores in a the true bits from either a or b.
func (a *BitSet) Union(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] | b.v[i]
	}
	if len(b.v) > len(a.v) {
		a.v = append(a.v, b.v[len(a.v):]...)
	}
}

// Insersect stores in a the true bits common to both a and b.
func (a *BitSet) Intersect(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] & b.v[i]
	}
	if len(a.v) > len(b.v) {
		a.v = a.v[:len(b.v)]
	}
}

// Difference stores in a the true bits present in a and not in b.
func (a *BitSet) Difference(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] & ^b.v[i]
	}
}

// SymmetricDifference stores in a the true bits which are either
// in a or in b, but not in both.
func (a *BitSet) SymmetricDifference(b *BitSet) {
	for i := 0; i < len(a.v) && i < len(b.v); i++ {
		a.v[i] = a.v[i] ^ b.v[i]
	}
	if len(b.v) > len(a.v) {
		a.v = append(a.v, b.v[len(a.v):]...)
	}
}
