package equalcmp

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/types"
)

//EqualCmp uses go-cmp to compare actual with expected.  Equal is strict about
//types when performing comparisons.
//It is an error for both actual and expected to be nil.  Use BeNil() instead.
//goland:noinspection GoUnusedExportedFunction
func EqualCmp(expected interface{}, options ...cmp.Option) types.GomegaMatcher {
	return &equalCmpMatcher{
		expected: expected,
		options:  options,
	}
}

type equalCmpMatcher struct {
	expected interface{}
	options  cmp.Options
}

func (matcher *equalCmpMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil && matcher.expected == nil {
		return false, fmt.Errorf("Refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized")
	}
	return cmp.Equal(actual, matcher.expected, matcher.options), nil
}

func (matcher *equalCmpMatcher) FailureMessage(actual interface{}) (message string) {
	diff := cmp.Diff(actual, matcher.expected, matcher.options)
	return "Mismatch(-actual +expected):\n" + diff
}

func (matcher *equalCmpMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	diff := cmp.Diff(actual, matcher.expected, matcher.options)
	return "Mismatch(-actual +expected):\n" + diff
}
