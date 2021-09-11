# equal-cmp

A module that implements gomega's Equal with go-cmp.

The EqualCmp provided by this module makes it easier to understand the difference between the actual value and the expected value when they do not match than the Equal provided by gomega.

## Usage

```
import (
  . "github.com/onsi/gomega"
  . "github.com/kamikaze-zirou/equal-cmp"
)

...

Expect(actual).Should(EqualCmp(expected))
```
