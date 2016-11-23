// Written by https://xojoc.pw. Public Domain.

package bitset

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

var cfg = &quick.Config{MaxCount: 1000, Rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano()))}
var limit = 10 * 1000

func (BitSet) Generate(rand *rand.Rand, size int) reflect.Value {
	s := BitSet{}
	for i := 0; i < size; i++ {
		s.Set(rand.Intn(limit))
	}
	return reflect.ValueOf(s)
}

type index int

func (index) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(index(rand.Intn(limit)))
}

func check(t *testing.T, f interface{}) {
	if err := quick.Check(f, cfg); err != nil {
		t.Error(err)
	}
}

func TestAny(t *testing.T) {
	any := func(i index) bool {
		s := &BitSet{}
		if s.Any() {
			return false
		}
		s.Set(int(i))
		return s.Any()
	}
	check(t, any)
}
func TestAnyRange(t *testing.T) {
	any := func(i, l, r index) bool {
		s := &BitSet{}
		if s.AnyRange(int(i-l), int(i+r)+1) {
			return false
		}
		s.Set(int(i))
		return s.AnyRange(int(i-l), int(i+r)+1)
	}
	check(t, any)
}
func TestCardinality(t *testing.T) {
	card := func(is []index) bool {
		s := &BitSet{}
		if s.Cardinality() != 0 {
			return false
		}
		count := 0
		for _, i := range is {
			j := int(i)
			if s.Get(j) {
				continue
			}
			s.Set(j)
			count++
		}
		return s.Cardinality() == count
	}
	check(t, card)
}
func TestClear(t *testing.T) {
	c := func(i index) bool {
		s := &BitSet{}
		j := int(i)
		s.Set(j)
		s.Clear(j)
		return !s.Get(j)
	}
	check(t, c)
}
func TestClearRange(t *testing.T) {
	c := func(i, j index) bool {
		s := &BitSet{}
		k, t := int(i), int(j)
		s.SetRange(k, t)
		s.ClearRange(k, t)
		return s.None()
	}
	check(t, c)
}
func TestClone(t *testing.T) {
	c := func(s BitSet) bool {
		t := s.Clone()
		return s.Equal(t)
	}
	check(t, c)
}
func TestDifference(t *testing.T) {
	d := func(a BitSet, b BitSet) bool {
		a1 := a.Clone()
		b1 := b.Clone()
		a1.Difference(b1)
		if !a1.SubSet(&a) {
			return false
		}
		a1.Intersect(b1)
		if a1.Any() {
			return false
		}

		a1 = a.Clone()
		b1 = b.Clone()
		b1.Difference(a1)
		if !b1.SubSet(&b) {
			return false
		}
		b1.Intersect(a1)
		if b1.Any() {
			return false
		}

		return true
	}
	check(t, d)
}
func TestEqual(t *testing.T) {
	eq := func(is []index) bool {
		s := &BitSet{}
		t := &BitSet{}
		for _, i := range is {
			s.Set(int(i))
			t.Set(int(i))
		}
		return s.Equal(t)
	}
	check(t, eq)
}
func TestGet(t *testing.T) {
	g := func(i index) bool {
		s := &BitSet{}
		j := int(i)
		if s.Get(j) {
			return false
		}
		s.Set(j)
		return s.Get(j)
	}
	check(t, g)
}
func TestGetRange(t *testing.T) {
	g := func(i, j index) bool {
		s := &BitSet{}
		k, t := int(i), int(j)
		/*
			if s.GetRange(k, t) {
				return false
			}
		*/
		s.SetRange(k, t)
		return s.GetRange(k, t)
	}
	check(t, g)
}
func TestIntersect(t *testing.T) {
	u := func(a, b, c BitSet) bool {
		// A ∩ B = B ∩ A
		a1 := a.Clone()
		a1.Intersect(&b)
		b2 := b.Clone()
		b2.Intersect(&a)
		if !a1.Equal(b2) {
			return false
		}

		// A ∩ (B ∩ C) = (A ∩ B) ∩ C
		a3 := a.Clone()
		b3 := b.Clone()
		b3.Intersect(&c)
		a3.Intersect(b3)
		a4 := a.Clone()
		a4.Intersect(&b)
		a4.Intersect(&c)
		if !a3.Equal(a4) {
			return false
		}

		// A ∩ B ⊆ A
		a5 := a.Clone()
		a5.Intersect(&b)
		if !a5.SubSet(&a) {
			return false
		}

		// A ∩ A = A
		a6 := a.Clone()
		a6.Intersect(&a)
		if !a.Equal(a6) {
			return false
		}

		// A ∩ ∅ = ∅
		a7 := &BitSet{}
		a7.Intersect(&a)
		if !a7.Equal(&BitSet{}) {
			return false
		}

		// A ⊆ B iff A ∩ B = A
		a8 := &BitSet{}
		a8.Intersect(&b)
		if a.SubSet(&b) && !a8.Equal(&a) {
			return false
		}

		return true
	}
	check(t, u)
}
func TestLen(t *testing.T) {
	len := func(is []index) bool {
		s := &BitSet{}
		if s.Len() != 0 {
			return false
		}
		max := 0
		for _, i := range is {
			j := int(i)
			s.Set(j)
			if j+1 > max {
				max = j + 1
			}
			if s.Len() != max {
				return false
			}
		}
		return true
	}
	check(t, len)
}
func TestNext(t *testing.T) {
	n := func(s BitSet) bool {
		i := -1
		for {
			previ := i
			i = s.Next(previ)
			if i == -1 {
				break
			}
			for j := previ + 1; j < i; j++ {
				if s.Get(j) {
					return false
				}
			}
			if !s.Get(i) {
				return false
			}
		}
		return true
	}
	check(t, n)
}
func TestPrev(t *testing.T) {
	p := func(s BitSet) bool {
		i := s.Len()
		for {
			previ := i
			i = s.Prev(previ)
			if i == -1 {
				break
			}
			for j := previ - 1; j > i; j-- {
				if s.Get(j) {
					return false
				}
			}
			if !s.Get(i) {
				return false
			}
		}
		return true
	}
	check(t, p)
}

/*
func Test(t *testing.T) {
	 := func(i index) bool {
	}
	check(t, )
}
*/

func TestUnion(t *testing.T) {
	u := func(a, b, c BitSet) bool {
		// A ∪ B = B ∪ A
		a1 := a.Clone()
		a1.Union(&b)
		b2 := b.Clone()
		b2.Union(&a)
		if !a1.Equal(b2) {
			return false
		}

		// A ∪ (B ∪ C) = (A ∪ B) ∪ C
		a3 := a.Clone()
		b3 := b.Clone()
		b3.Union(&c)
		a3.Union(b3)
		a4 := a.Clone()
		a4.Union(&b)
		a4.Union(&c)
		if !a3.Equal(a4) {
			return false
		}

		// A ⊆ (A ∪ B)
		a5 := a.Clone()
		a5.Union(&b)
		if !a.SubSet(a5) {
			return false
		}

		// A ∪ A = A
		a6 := a.Clone()
		a6.Union(&a)
		if !a.Equal(a6) {
			return false
		}

		// A ∪ ∅ = A
		a7 := &BitSet{}
		a7.Union(&a)
		if !a.Equal(a7) {
			return false
		}

		// A ⊆ B iff A ∪ B = B
		a8 := &BitSet{}
		a8.Union(&b)
		if a.SubSet(&b) && !a8.Equal(&b) {
			return false
		}

		return true
	}
	check(t, u)
}

func TestDeMorgan(t *testing.T) {
	m := func(a, b BitSet) bool {
		max := a.Len()
		if b.Len() > max {
			max = b.Len()
		}
		// (A ∪ B)′ = A′ ∩ B′
		a1 := a.Clone()
		a1.Union(&b)
		a1.ToggleRange(0, max)
		a2 := a.Clone()
		b2 := b.Clone()
		a2.ToggleRange(0, max)
		b2.ToggleRange(0, max)
		a2.Intersect(b2)
		if !a1.Equal(a2) {
			return false
		}

		// (A ∩ B)′ = A′ ∪ B′
		a3 := a.Clone()
		a3.Intersect(&b)
		a3.ToggleRange(0, max)
		a4 := a.Clone()
		b4 := b.Clone()
		a4.ToggleRange(0, max)
		b4.ToggleRange(0, max)
		a4.Union(b4)
		return a3.Equal(a4)
	}
	check(t, m)
}
