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

import "fmt"

func Example() {
	s := New(6)
	s.Set(2, true)
	s.Set(3, true)
	s.Set(5, true)
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.Get(i))
	}

	fmt.Println(s)

	// Output:
	// false
	// false
	// true
	// true
	// false
	// true
	// 001101
}

func ExampleBitSet_Union() {
	a := New(2)
	b := New(4)
	a.Set(0, true)
	b.Set(3, true)
	fmt.Println(a)
	fmt.Println(b)
	a.Union(b)
	fmt.Println(a)

	// Output:
	// 10
	// 0001
	// 1001
}

func ExampleBitSet_Intersect() {
	a := New(4)
	b := New(2)
	a.Set(0, true)
	a.Set(3, true)
	b.Set(0, true)
	b.Set(1, true)
	fmt.Println(a)
	fmt.Println(b)
	a.Intersect(b)
	fmt.Println(a)

	// Output:
	// 1001
	// 11
	// 10
}

func ExampleBitSet_Difference() {
	a := New(3)
	a.Set(0, true)
	a.Set(1, true)
	a.Set(2, true)
	b := New(2)
	b.Set(1, true)
	a.Difference(b)
	fmt.Println(a)

	// Output:
	// 101
}

func ExampleBitSet_Complement() {
	a := New(3)
	a.Set(1, true)
	a.Complement()
	fmt.Println(a)

	// Output:
	// 101
}

func ExampleBitSet_String() {
	a := New(3)
	a.Set(1, true)
	fmt.Println(a) // fmt automatically calls String

	// Output:
	// 010
}
