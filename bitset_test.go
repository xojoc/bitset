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

import "testing"

func TestBitSet_Set(t *testing.T) {
	s := New(bpw)
	s.Set(1, true)
	s.Set(bpw-1, true)
	rs := "0100000000000000000000000000000000000000000000000000000000000001"
	if s.String() != rs {
		t.Errorf("string %q want %q", s, rs)
	}

	s = New(bpw + 1)
	s.Set(1, true)
	s.Set(bpw-1, true)
	s.Set(bpw, true)
	rs = "01000000000000000000000000000000000000000000000000000000000000011"
	if s.String() != rs {
		t.Errorf("String %q want %q", s, rs)
	}

	s = New(1)
	s.Set(0, true)
	s.Set(0, false)
	if s.Get(0) != false {
		t.Errorf("Get(0) %v want %v", s.Get(0), false)
	}
}

func TestBitSet_Len(t *testing.T) {
	a := New(0)
	if a.Len() != 0 {
		t.Errorf("Len %d want %d", a.Len(), 0)
	}

	a = New(bpw)
	if a.Len() != bpw {
		t.Errorf("Len %d want %d", a.Len(), bpw)
	}

	a = New(bpw + 1)
	if a.Len() != bpw+1 {
		t.Errorf("Len %d want %d", a.Len(), bpw+1)
	}
}

func TestBitSet_Toggle(t *testing.T) {
	a := New(1)
	a.Toggle(0)
	if a.Get(0) != true {
		t.Errorf("Get %v want %v", a.Get(0), true)
	}
}

func TestBitSet_Copy(t *testing.T) {
	a := New(3)
	a.Set(1, true)
	b := a.Copy()
	if a.Len() != b.Len() {
		t.Errorf("Copy len %d want %d", b.Len(), a.Len())
	}
	if a.String() != b.String() {
		t.Errorf("Copy string %q want %q", b, a)
	}
}

func TestBitSet_Union(t *testing.T) {
	a := New(2)
	b := New(bpw + 1)
	a.Union(b)
	if a.Len() != b.Len() {
		t.Errorf("Len %d want %d", a.Len(), b.Len())
	}
	if a.nb != b.nb {
		t.Errorf("nb %d want %d", a.nb, b.nb)
	}
}

func TestBitSet_Insersect(t *testing.T) {
	a := New(bpw + 1)
	b := New(2)
	a.Intersect(b)
	if a.Len() != b.Len() {
		t.Errorf("Len %d want %d", a.Len(), b.Len())
	}
	if a.nb != b.nb {
		t.Errorf("nb %d want %d", a.nb, b.nb)
	}
}

func TestBitSet_Complement(t *testing.T) {
	a := New(3)
	a.Set(1, true)
	a.Complement()
	if a.String() != "101" {
		t.Errorf("Complement %q want %q", a, "101")
	}
}
