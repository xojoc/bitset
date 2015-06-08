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

func TestBitSet_Clone(t *testing.T) {
	a := &BitSet{}
	a.Set(1)
	b := a.Clone()
	if a.String() != b.String() {
		t.Errorf("Copy string %q want %q", b, a)
	}
}

func TestBitSet_String(t *testing.T) {
	s := &BitSet{}
	str := s.String()
	if str != "" {
		t.Errorf("String %q want %q", s.String(), "")
	}

	s.Set(1)
	str = s.String()
	if str != "01" {
		t.Errorf("String %q want %q", s.String(), "01")
	}
}
func TestBitSet_Set(t *testing.T) {
	s := &BitSet{}
	s.Set(1)
	s.Set(bpw - 1)
	if !s.Get(1) || !s.Get(bpw-1) {
		t.Errorf("Can't Get(1) || Get(bpw-1)")
	}

	s = &BitSet{}
	s.Set(1)
	s.Set(bpw*3 - 1)
	s.Set(bpw * 3)
	if !s.Get(1) || !s.Get(bpw*3-1) || !s.Get(bpw*3) {
		t.Errorf("Can't Get(1) || Get(bpw*3-1) || Get(bpw*3)")
	}
}

func TestBitSet_Clear(t *testing.T) {
	s := &BitSet{}
	s.Set(0)
	s.Clear(0)
	if s.Get(0) != false {
		t.Errorf("Get(0) %v want %v", s.Get(0), false)
	}
}

func TestBitSet_Toggle(t *testing.T) {
	a := &BitSet{}
	a.Toggle(0)
	if a.Get(0) != true {
		t.Errorf("Get %v want %v", a.Get(0), true)
	}

	a.Toggle(0)
	if a.Get(0) != false {
		t.Errorf("Get %v want %v", a.Get(0), false)
	}
}

func TestBitSet_Len(t *testing.T) {
	a := &BitSet{}
	if a.Len() != 0 {
		t.Errorf("Len %d want %d", a.Len(), 0)
	}

	a.Set(0)
	if a.Len() != 1 {
		t.Errorf("Len %d want %d", a.Len(), 1)
	}

	a.Set(999)
	if a.Len() != 1000 {
		t.Errorf("Len %d want %d", a.Len(), 1000)
	}

	a.Toggle(999)
	if a.Len() != 1 {
		t.Errorf("Len %d want %d", a.Len(), 1)
	}
}

func TestBitSet_Any(t *testing.T) {
	a := &BitSet{}
	if a.Any() != false {
		t.Errorf("Get %v want %v", a.Any(), false)
	}
	a.Set(bpw*2 + 1)
	if a.Any() != true {
		t.Errorf("Get %v want %v", a.Any(), true)
	}
}

func TestBitSet_Cardinality(t *testing.T) {
	a := &BitSet{}
	if a.Cardinality() != 0 {
		t.Errorf("Cardinality %d want %d", a.Cardinality(), 0)
	}
	a.Set(1)
	if a.Cardinality() != 1 {
		t.Errorf("Count %d want %d", a.Cardinality(), 1)
	}
	a.Set(bpw)
	a.Set(bpw * 2)
	a.Set(bpw*3 + 2)
	if a.Cardinality() != 4 {
		t.Errorf("Cardinality %d want %d", a.Cardinality(), 4)
	}
}

func TestBitSet_Next(t *testing.T) {
	a := BitSet{}
	i, b := a.Next(0)
	if i != -1 || b != false {
		t.Errorf("Next %v,%v want %v,%v", i, b, -1, false)
	}

	a.Set(bpw)
	i, b = a.Next(0)
	if i != bpw || b != true {
		t.Errorf("Next %v,%v want %v,%v", i, b, bpw, true)
	}

	a.Set(bpw * 2)
	i, b = a.Next(0)
	if i != bpw || b != true {
		t.Errorf("Next %v,%v want %v,%v", i, b, bpw, true)
	}

	i, b = a.Next(bpw)
	if i != bpw*2 || b != true {
		t.Errorf("Next %v,%v want %v,%v", i, b, bpw*2, true)
	}
}

func TestBitSet_Prev(t *testing.T) {
	a := BitSet{}
	i, b := a.Prev(bpw)
	if i != -1 || b != false {
		t.Errorf("Next %v,%v want %v,%v", i, b, -1, false)
	}

	a.Set(bpw)
	i, b = a.Prev(bpw * 2)
	if i != bpw || b != true {
		t.Errorf("Next %v,%v want %v,%v", i, b, bpw, true)
	}

	a.Set(bpw * 2)
	i, b = a.Prev(bpw * 2)
	if i != bpw || b != true {
		t.Errorf("Next %v,%v want %v,%v", i, b, bpw, true)
	}

	i, b = a.Prev(bpw * 3)
	if i != bpw*2 || b != true {
		t.Errorf("Next %v,%v want %v,%v", i, b, bpw*2, true)
	}
}

func TestBitSet_SuperSet(t *testing.T) {
	a := &BitSet{}
	b := &BitSet{}
	if a.SuperSet(b) != true {
		t.Errorf("SuperSet %v want %v", a.SuperSet(b), true)
	}
	a.Set(1)
	if a.SuperSet(b) != true {
		t.Errorf("SuperSet %v want %v", a.SuperSet(b), true)
	}
	b.Set(0)
	if a.SuperSet(b) != false {
		t.Errorf("SuperSet %v want %v", a.SuperSet(b), false)
	}
	a.Set(0)
	a.Set(bpw)
	b.Set(bpw)
	if a.SuperSet(b) != true {
		t.Errorf("SuperSet %v want %v", a.SuperSet(b), true)
	}
}

func TestBitSet_Union(t *testing.T) {
	a := &BitSet{}
	a.Set(0)
	b := &BitSet{}
	b.Set(2)
	c := &BitSet{}
	c.Set(0)
	c.Set(2)
	a.Union(b)
	if a.Equal(c) != true {
		t.Errorf("Equal %v want %v", a.Equal(c), true)
	}
}

/*
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

func TestBitSet_Equal(t *testing.T) {
	a := New(3).Set(1)
	b := New(2).Set(1)
	if a.Equal(b) != false {
		t.Errorf("Equal %v want %v", a.Equal(b), false)
	}

	a = New(bpw + 5).Set(bpw - 1).Set(bpw).Set(bpw + 1)
	b = New(bpw + 5).Set(bpw - 1).Set(bpw).Set(bpw + 1)
	if a.Equal(b) != true {
		t.Errorf("Equal %v want %v", a.Equal(b), true)
	}

	a = New(3).Set(0)
	b = New(3).Set(1)
	if a.Equal(b) != false {
		t.Errorf("Equal %v want %v", a.Equal(b), false)
	}

	a = New(bpw + 5).Set(bpw)
	b = New(bpw + 5).Set(bpw - 1)
	if a.Equal(b) != false {
		t.Errorf("Equal %v want %v", a.Equal(b), false)
	}
}
*/
