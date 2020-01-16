# BitSet data structure
*bitset* is a library written in [golang](https://golang.org) implementing a BitSet data structure.

# Usage
First install the library with
```
go get xojoc.pw/bitset
```

...then run...

````
go test
``````

if something fails open an issue.

As an example let's list all the composite numbers below 10:
```
package main

import (
	"fmt"
	"xojoc.pw/bitset"
)

func main() {
        c := &bitset.BitSet{}
        // Set all prime numbers to true.
        c.Set(2)
        c.Set(3)
        c.Set(5)
        c.Set(7)
        c.ToggleRange(1,10+1)
        for i := 1; i < c.Len(); i++ {
                if c.Get(i) {
                        fmt.Printf("%d is composite\n", i)
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
see [godoc](https://godoc.org/xojoc.pw/bitset) for the complete documentation.

# Why?

 * Great and stable API
 * Lots of tests, [fuzzing](https://github.com/google/gofuzz) and [quicktesting](https://golang.org/pkg/testing/quick/)!
 * [Good documentation](https://godoc.org/xojoc.pw/bitset) with lots of examples.

Also see [why use xojoc.pw/bitset and not math/big](https://typed.pw/a/29).

# Who?
*bitset* was written by [Alexandru Cojocaru](https://xojoc.pw).

# [Donate!](https://xojoc.pw/donate)

# License
*bitset* is [Free Software](https://www.gnu.org/philosophy/free-sw.html) and in the Public Domain. No warranty.
