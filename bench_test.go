// Written by http://xojoc.pw. Public Domain.

package bitset

import (
	"math/big"
	"testing"
)

func Benchmark(b *testing.B) {
}

func BenchmarkClone(b *testing.B) {
	s := &BitSet{}
	s.Set(1000 * 100 * 1000)
	for n := 0; n < b.N; n++ {
		n := s.Clone()
		n.Get(0)
	}
}
func BenchmarkSet(b *testing.B) {
	len := 1000
	s := &BitSet{}
	for n := 0; n < b.N; n++ {
		for i := 0; i < len; i++ {
			s.Set(i)
		}
	}
}
func BenchmarkSetBig(b *testing.B) {
	len := 1000
	s := big.NewInt(0)
	for n := 0; n < b.N; n++ {
		for i := 0; i < len; i++ {
			s.SetBit(s, i, 1)
		}
	}
}
func BenchmarkSetHigh(b *testing.B) {
	s := &BitSet{}
	i := 1
	for n := 0; n < b.N; n++ {
		s.Set(100 * 1000 * 1000 * i)
		i++
	}
}
func BenchmarkSetRange(b *testing.B) {
	s := &BitSet{}
	s.SetRange(0, 100*1000*1000)
}
func BenchmarkClearRange(b *testing.B) {
	s := &BitSet{}
	s.SetRange(0, 100*1000*1000)
	s.ClearRange(0, 100*1000*1000)
}
func BenchmarkToggleRange(b *testing.B) {
	s := &BitSet{}
	s.SetRange(50*1000*1000, 100*1000*1000)
	s.ToggleRange(0, 100*1000*1000)
}
func BenchmarkGetEmpty(b *testing.B) {
	s := &BitSet{}
	len := 1000 * 1000
	for n := 0; n < b.N; n++ {
		for i := 0; i < len; i++ {
			s.Get(i)
		}
	}
}
func BenchmarkGetEmptyBig(b *testing.B) {
	s := big.NewInt(0)
	len := 1000 * 1000
	s.SetBit(s, len-1, 1)
	for n := 0; n < b.N; n++ {
		for i := 0; i < len; i++ {
			s.Bit(i)
		}
	}
}

/*
func BenchmarkUnion(b *testing.B) {
	a := New(1000)
	c := New(1000)
	for n := 0; n < b.N; n++ {
		a.Union(c)
	}
}

func BenchmarkUnionBig(b *testing.B) {
	len := 1000
	a := big.NewInt(0)
	a.SetBit(a, len-1, 1)
	c := big.NewInt(0)
	c.SetBit(c, len-1, 1)
	for n := 0; n < b.N; n++ {
		a.Or(a, c)
	}
}

func BenchmarkString(b *testing.B) {
	a := New(100000)
	for n := 0; n < b.N; n++ {
		_ = a.String()
	}
}
*/
