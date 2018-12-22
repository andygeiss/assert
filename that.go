package assert

import (
	"fmt"
	"testing"
)

// That checks if state matches a given value by using a Matcher.
func That(t *testing.T, state interface{}, m Matcher) {
	if !m.Matches(state) {
		msg := fmt.Sprintf("[%v] does not match [%v]!", state, m.Value())
		t.Fatal(msg)
	}
}
