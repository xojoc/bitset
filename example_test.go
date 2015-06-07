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
	"fmt"
)

func Example() {
	// Create new BitSet
	s := &BitSet{}
	// Bitsets automatically grow
	s.Set(2)
	s.Set(3)
	fmt.Println(s.Get(0))
	fmt.Println(s.Get(2))
	// Out of range Get will return false
	fmt.Println(s.Get(1000))
	// Println automatically calls String method
	fmt.Println(s)

	t := &BitSet{}
	t.Set(2)
	t.Set(4)
	s.Intersect(t)
	fmt.Println(s)

	// Output:
	// false
	// true
	// false
	// 0011
	// 001
}

func ExampleBitSet_Union() {
	a := &BitSet{}
	a.Set(0)
	b := &BitSet{}
	b.Set(3)
	fmt.Println(a)
	fmt.Println(b)
	a.Union(b)
	fmt.Println(a)

	// Output:
	// 1
	// 0001
	// 1001
}

func ExampleBitSet_Intersect() {
	a := &BitSet{}
	a.Set(0)
	a.Set(3)
	b := &BitSet{}
	b.Set(0)
	b.Set(1)
	fmt.Println(a)
	fmt.Println(b)
	a.Intersect(b)
	fmt.Println(a)

	// Output:
	// 1001
	// 11
	// 1
}

func ExampleBitSet_Difference() {
	a := &BitSet{}
	a.Set(0)
	a.Set(1)
	a.Set(2)
	b := &BitSet{}
	b.Set(1)
	fmt.Println(a)
	fmt.Println(b)
	a.Difference(b)
	fmt.Println(a)

	// Output:
	// 111
	// 01
	// 101
}

func ExampleBitSet_SymmetricDifference() {
	a := &BitSet{}
	a.Set(0)
	a.Set(1)
	b := &BitSet{}
	b.Set(0)
	b.Set(2)
	fmt.Println(a)
	fmt.Println(b)
	a.SymmetricDifference(b)
	fmt.Println(a)

	// Output:
	// 11
	// 101
	// 011
}

func ExampleBitSet_String() {
	a := &BitSet{}
	a.Set(0)
	a.Set(2)
	fmt.Println(a) // fmt automatically calls String

	// Output:
	// 101
}
