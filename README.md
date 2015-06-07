# BitSet data structure
*bitset* is a library written in [golang](http://golang.org) implementing a BitSet data structure.

# Usage
First install the library with
```
go get github.com/xojoc/bitset
```

As an example let's list all the composite numbers below 10:
```
package main

import (
	"fmt"
	"github.com/xojoc/bitset"
)

func main() {
        c := &bitset.BitSet{}
	// Set all prime numbers to true.
        c.Set(1)
        c.Set(2)
        c.Set(4)
        c.Set(6)
        c.ToggleRange(0,10)
	for i := 0; i < c.Len(); i++ {
		if c.Get(i) == true {
			fmt.Printf("%d is composite\n", i+1)
		}
	}
}      
```
Output:
```
1 is composite
4 is composite
6 is composite
8 is composite
9 is composite
10 is composite
```
see [godoc](http://godoc.org/github.com/xojoc/bitset) for the complete documentation.

# Why?

 * Great API
 * 100% test coverage
 * [Great documentation](http://godoc.org/github.com/xojoc/bitset)

Also see [why xojoc/bitset and not math/big](http://typed.pw/a/29).

# Who?
*bitset* was written by Alexandru cojocaru (http://xojoc.pw).

# License
*bitset* is released under the GPLv3 or later, see [COPYING](COPYING).