// Written by https://xojoc.pw. Public Domain.

package bitset

import "testing"

func TestBitSet_ShiftLeft(t *testing.T) {
	s := &BitSet{}
	s.ShiftLeft(5)
	s.Set(1)
	s.Set(3)
	s.Set(5)
	s.ShiftLeft(2)
	if s.String() != "0101" {
		t.Errorf("ShiftLeft %q want %q", s.String(), "0101")
	}
}
func TestBitSet_ShiftRight(t *testing.T) {
	s := &BitSet{}
	s.ShiftRight(5)
	s.Set(0)
	s.Set(2)
	s.ShiftRight(1)
	if s.String() != "0101" {
		t.Errorf("ShiftRight %q want %q", s.String(), "0101")
	}
}
