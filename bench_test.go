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

package bitset

import (
	"math/big"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	s := New(1000)
	for n := 0; n < b.N; n++ {
		for i := 0; i < s.Len(); i++ {
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

func BenchmarkGet(b *testing.B) {
	s := New(1000)
	for n := 0; n < b.N; n++ {
		for i := 0; i < s.Len(); i++ {
			s.Get(i)
		}
	}
}

func BenchmarkGetBig(b *testing.B) {
	len := 1000
	s := big.NewInt(0)
	s.SetBit(s, len-1, 1)
	for n := 0; n < b.N; n++ {
		for i := 0; i < len; i++ {
			s.Bit(i)
		}
	}
}

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
