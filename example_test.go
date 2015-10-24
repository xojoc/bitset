// Written by http://xojoc.pw. Public Domain.

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
