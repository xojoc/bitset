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
	s.Set(1).Set(bpw - 1)
	rs := "0100000000000000000000000000000000000000000000000000000000000001"
	if s.String() != rs {
		t.Errorf("string %q want %q", s, rs)
	}

	s = New(bpw + 1)
	s.Set(1).Set(bpw - 1).Set(bpw)
	rs = "01000000000000000000000000000000000000000000000000000000000000011"
	if s.String() != rs {
		t.Errorf("String %q want %q", s, rs)
	}

	s = New(1)
	if s.Set(0).Clear(0).Get(0) != false {
		t.Errorf("Get(0) %v want %v", s.Get(0), false)
	}
}

func TestBitSet_SetAll(t *testing.T) {
	s := New(bpw*2 + 3)
	s.SetAll()
	if s.All() != true {
		t.Errorf("Not all bits are set")
	}
}

func TestBitSet_ClearAll(t *testing.T) {
	s := New(bpw*2 + 3).Set(0).Set(bpw).Set(bpw + 1).Set(bpw*2 + 1)
	s.ClearAll()
	if s.Any() {
		t.Errorf("Not all bits are cleared")
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
	a := New(1).Toggle(0)
	if a.Get(0) != true {
		t.Errorf("Get %v want %v", a.Get(0), true)
	}
}

func TestBitSet_Clone(t *testing.T) {
	a := New(3).Set(1)
	b := a.Clone()
	if a.Len() != b.Len() {
		t.Errorf("Copy len %d want %d", b.Len(), a.Len())
	}
	if a.String() != b.String() {
		t.Errorf("Copy string %q want %q", b, a)
	}
}

func TestBitSet_All(t *testing.T) {
	a := New(3).Set(0).Set(1).Set(2)
	if a.All() != true {
		t.Errorf("Get %v want %v", a.All(), true)
	}

	a = New(3).Set(1).Set(2)
	if a.All() != false {
		t.Errorf("Get %v want %v", a.All(), false)
	}

	a = New(bpw + 1)
	for i := 0; i < bpw+1; i++ {
		a.Set(i)
	}
	a.Clear(0)
	if a.All() != false {
		t.Errorf("Get %v want %v", a.All(), false)
	}
}

func TestBitSet_Any(t *testing.T) {
	a := New(3).Set(1)
	if a.Any() != true {
		t.Errorf("Get %v want %v", a.Any(), true)
	}

	a = New(3)
	if a.Any() != false {
		t.Errorf("Get %v want %v", a.Any(), false)
	}

	a = New(bpw + 1)
	a.Set(0)
	if a.Any() != true {
		t.Errorf("Get %v want %v", a.Any(), true)
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

func TestBitSet_ToggleAll(t *testing.T) {
	a := New(3).Set(1)
	a.ToggleAll()
	if a.String() != "101" {
		t.Errorf("Complement %q want %q", a, "101")
	}
}

func TestBitSet_Count(t *testing.T) {
	a := New(3).Set(1)
	if a.Count() != 1 {
		t.Errorf("Count %d want %d", a.Count(), 1)
	}

	a = New(bpw*2 + 3).Set(0).Set(bpw).Set(bpw * 2).Set(bpw*2 + 2)
	if a.Count() != 4 {
		t.Errorf("Count %d want %d", a.Count(), 4)
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
