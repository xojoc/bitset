// Written by https://xojoc.pw. Public Domain.

// Package bitset implements a BitSet data structure.
//
// A BitSet is a mapping between unsigned integers and boolean values.
// You can Set, Clear, Toggle single bits or Union, Intersect, Difference sets.
//
// Indexes start at 0. Ranges have the first index included and the second
// one excluded (like go slices).
//
// BitSets are dynamicaly-sized they grow and shrink automatically.
//
// All methods modify their receiver in place to avoid futile memory usage.
// If you want to keep the original BitSet simply Clone it.
//
// Use Clone when you want to copy a BitSet. Plese note that this will
// *not* work:
//     var x BitSet
//     x.Add(1)
//     y := x  // wrong! use Clone
//     y.Add(2)
//
//
// If you wonder why you should use this package and not math/big see:
// https://typed.pw/a/29
package bitset // import "xojoc.pw/bitset"

// TODO: intersects next/prev zero
// TODO: fmt.Formatter

// Bit tricks: http://graphics.stanford.edu/~seander/bithacks.html

// Bits per word
const bpw int = 8 << (^uint(0)>>8&1 + ^uint(0)>>16&1 + ^uint(0)>>32&1)

// BitSet data structure.
type BitSet struct {
	// underlying vector
	v []uint
}

// All the functions below assume the bitsets in input
// have no trailing zero bytes. Functions that clear
// bits (Clear, Toggle, Intersect, Difference, SymmetricDifference)
// must call this function, which removes all the trailing zero bytes.
func (s *BitSet) autoShrink() {
	for i := len(s.v) - 1; i >= 0; i-- {
		if s.v[i] == 0 {
			s.v = s.v[:len(s.v)-1]
		} else {
			break
		}
	}
	s.v = s.v[:len(s.v):len(s.v)]
}

// Clone makes a copy of s.
func (s BitSet) Clone() *BitSet {
	t := &BitSet{}
	t.v = append(t.v, s.v...)
	return t
}

// String returns a string representation of s.
func (s BitSet) String() string {
	b := make([]byte, s.Len())
	for i := 0; i < s.Len(); i++ {
		if s.Get(i) {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	return string(b)
}

// Set sets the bit at index i.
func (s *BitSet) Set(i int) {
	if i < 0 {
		return
	}
	for i/bpw+1 > len(s.v) {
		s.v = append(s.v, 0)
	}
	s.v[i/bpw] |= 1 << uint(i%bpw)
}

// SetRange sets the bits between i (included) and j (excluded).
func (s *BitSet) SetRange(i, j int) {
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}
	for k := i; k < j; k++ {
		s.Set(k)
	}
}

// Clear clears the bit at index i.
func (s *BitSet) Clear(i int) {
	if i < 0 {
		return
	}
	if (i/bpw + 1) > len(s.v) {
		return
	}
	s.v[i/bpw] &= ^(1 << uint(i%bpw))
	s.autoShrink()
}

// ClearRange clears the bits between i (included) and j (excluded).
func (s *BitSet) ClearRange(i, j int) {
	for k := i; k < j; k++ {
		s.Clear(k)
	}
}

// Toggle inverts the bit at index i.
func (s *BitSet) Toggle(i int) {
	if i < 0 {
		return
	}
	if i/bpw+1 > len(s.v) {
		s.Set(i)
	} else {
		s.v[i/bpw] ^= 1 << uint(i%bpw)
		s.autoShrink()
	}
}

// ToggleRange inverts the bits between i (included) and j (excluded).
func (s *BitSet) ToggleRange(i, j int) {
	for k := i; k < j; k++ {
		s.Toggle(k)
	}
}

// Get returns true if the bit at index i is set, false otherwise.
// If i < 0, returns true.
func (s *BitSet) Get(i int) bool {
	if i < 0 {
		return true
	}
	if i/bpw+1 > len(s.v) {
		return false
	}
	return (s.v[i/bpw] & (1 << uint(i%bpw))) != 0
}

// GetRange returns true if the bits between i (included) and j (excluded) are set, false otherwise.
// If i < 0 and j < 0 return true.
func (s *BitSet) GetRange(i, j int) bool {
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}
	for k := i; k < j; k++ {
		if !s.Get(k) {
			return false
		}
	}
	return true
}

// Len returns the number of bits up to and including the highest bit set.
func (s *BitSet) Len() int {
	// NOTE: autoShrink is always called by functions that
	// set bits to zero, but just to be sure we call
	// it here anyway.
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
	for _, e := range s.v {
		if e != 0 {
			return true
		}
	}
	return false
}

// AnyRange returns true if any bit between i (included) and j (excluded) is set, false otherwise.
// If i < 0 and j < 0 return true.
func (s *BitSet) AnyRange(i, j int) bool {
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}
	for k := i; k < j; k++ {
		if s.Get(k) {
			return true
		}
	}
	return false
}

// None returns true if no bit is set, false otherwise.
func (s *BitSet) None() bool {
	return !s.Any()
}

// NoneRange returns true if no bit between i (included) and j (excluded) is set, false otherwise.
// If i < 0 and j < 0 return true.
func (s *BitSet) NoneRange(i, j int) bool {
	return !s.AnyRange(i, j)
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
	c := 0
	for _, e := range s.v {
		c += countBits(e)
	}
	return c
}

// Next returns the index of the next bit set after i.
// If no bit was found returns -1.
func (s *BitSet) Next(i int) int {
	if i < 0 {
		i = -1
	}
	for j := i + 1; j < s.Len(); j++ {
		if s.Get(j) {
			return j
		}
	}
	return -1
}

// Prev returns the index of the previous bit set before i.
// If no bit was found returns -1.
func (s *BitSet) Prev(i int) int {
	for j := i - 1; j >= 0; j-- {
		if s.Get(j) {
			return j
		}
	}
	return -1
}

// Equal returns true if s and t have the same bits set, false otherwise.
func (s *BitSet) Equal(t *BitSet) bool {
	if len(s.v) != len(t.v) {
		return false
	}
	for i, u := range s.v {
		if u != t.v[i] {
			return false
		}
	}
	return true
}

// SuperSet returns true if s is a super set of t, false otherwise.
func (s *BitSet) SuperSet(t *BitSet) bool {
	if len(s.v) < len(t.v) {
		return false
	}
	for i := 0; i < len(t.v); i++ {
		if t.v[i] & ^s.v[i] != 0 {
			return false
		}
	}
	return true
}

// SubSet returns true if s is a sub set of t, false otherwise.
func (s *BitSet) SubSet(t *BitSet) bool {
	return t.SuperSet(s)
}

// ShiftLeft moves each bit n positions to the left.
func (s *BitSet) ShiftLeft(n int) {
	for i := n; i < s.Len(); i++ {
		if s.Get(i) {
			s.Set(i - n)
		} else {
			s.Clear(i - n)
		}
	}
	s.ClearRange(s.Len()-n, s.Len())
}

// ShiftRight moves each bit n positions to the right.
func (s *BitSet) ShiftRight(n int) {
	len := s.Len()
	for i := len - 1; i >= 0; i-- {
		if s.Get(i) {
			s.Set(i + n)
		} else {
			s.Clear(i + n)
		}
	}
	s.ClearRange(0, n)
}

// Union stores in a the true bits from either s or t.
func (s *BitSet) Union(t *BitSet) {
	for i := 0; i < len(s.v) && i < len(t.v); i++ {
		s.v[i] = s.v[i] | t.v[i]
	}
	if len(t.v) > len(s.v) {
		s.v = append(s.v, t.v[len(s.v):]...)
	}
}

// Intersect stores in s the true bits common to both s and t.
func (s *BitSet) Intersect(t *BitSet) {
	for i := 0; i < len(s.v) && i < len(t.v); i++ {
		s.v[i] = s.v[i] & t.v[i]
	}
	if len(s.v) > len(t.v) {
		// FIXME: probably we should clear a.v
		s.v = s.v[:len(t.v)]
	}
	s.autoShrink()
}

// Difference stores in s the true bits present in s and not in t.
func (s *BitSet) Difference(t *BitSet) {
	for i := 0; i < len(s.v) && i < len(t.v); i++ {
		s.v[i] = s.v[i] & ^t.v[i]
	}
	if len(s.v) <= len(t.v) {
		s.autoShrink()
	}
}

// SymmetricDifference stores in s the true bits which are either
// in s or in t, but not in both.
func (s *BitSet) SymmetricDifference(t *BitSet) {
	for i := 0; i < len(s.v) && i < len(t.v); i++ {
		s.v[i] = s.v[i] ^ t.v[i]
	}
	if len(s.v) == len(t.v) {
		s.autoShrink()
	} else if len(s.v) < len(t.v) {
		s.v = append(s.v, t.v[len(s.v):]...)
	}
}
