# equal-cmp

A module that implements gomega's Equal with go-cmp.

The EqualCmp provided by this module makes it easier to understand the difference between the actual value and the expected value when they do not match than the Equal provided by gomega.

## Example

Given the following test code using ginkgo and gomega.

```go
package example_test

import (
	. "github.com/kamikazezirou/equal-cmp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Object struct {
	ID   int
	Name string
}

var _ = Describe("Example", func() {
	It("Example", func() {
		actual := Object{ID: 1, Name: "actual"}
		expected := Object{ID: 1, Name: "expected"}
		Expect(actual).Should(EqualCmp(expected))
	})
})
```

The error message is as follows.

```
• Failure [0.000 seconds]
Example
...
  Example [It]
  ...

  Mismatch(-actual +expected):
    equalcmp_test.Object{
    	ID:   1,
  - 	Name: "actual",
  + 	Name: "expected",
    }
```
